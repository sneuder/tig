package action

import (
	"clit-git/helper"
	"log"
	"strings"

	"github.com/urfave/cli/v2"
)

func ActionCmdRemove(cCtx *cli.Context) error {
	orgName := strings.TrimSpace(cCtx.Args().Get(0))

	if orgName == "" {
		log.Fatal("please provide an organization name")
	}

	helper.RemoveOrgByName(orgName)

	helper.SaveOrganizationInConfig()
	currentOrgName := helper.GetCurrentOrg()
	helper.RemoveOrgFolder(currentOrgName)

	if orgName == currentOrgName {
		helper.UpdateGlobalGitConfig(helper.GlobalGitConfig{
			Name:  "",
			Email: "",
			Org:   "",
		})
	}

	return nil
}
