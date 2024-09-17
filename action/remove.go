package action

import (
	"clit-git/helper"
	"log"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
)

func ActionCmdRemove(cCtx *cli.Context) error {
	orgName := strings.TrimSpace(cCtx.Args().Get(0))

	if orgName == "" {
		log.Fatal("please provide an organization name")
	}

	removeOrgByName(orgName)

	saveOrganizationInConfig()
	currentOrgName := getCurrentOrg()
	helper.RemoveOrgFolder(currentOrgName)

	if orgName == currentOrgName {
		updateGlobalGitConfig(GlobalGitConfig{
			Name:  "",
			Email: "",
			Org:   "",
		})
	}

	return nil
}

func removeOrgByName(orgName string) {
	orgs := helper.GetOrgs()
	org, orgIndex := helper.FindOrgByName(orgName)

	if org == nil {
		log.Fatal("organization not found")
	}

	orgs = append(orgs[:orgIndex], orgs[orgIndex+1:]...)
	helper.CreateOrgFile(orgs)
}

func getCurrentOrg() string {
	cmd := exec.Command("bash", "-c", "git config --global --list | grep 'user.org' | cut -d '=' -f 2")
	out, _ := cmd.Output()
	return strings.TrimSpace(string(out))
}
