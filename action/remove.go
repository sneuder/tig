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
	org, orgId := helper.GetCurrentOrg()

	if orgId == -1 {
		println("current organization not found")
	}

	helper.RemoveOrgFolder(org.Org)

	if orgName == org.Org {
		helper.UpdateGlobalGitConfig(helper.GlobalGitConfig{
			Name:  "",
			Email: "",
			Org:   "",
		})
	}

	return nil
}
