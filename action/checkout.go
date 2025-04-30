package action

import (
	"clit-git/helper"
	"clit-git/schema"
	"log"
	"strings"

	"github.com/urfave/cli/v2"
)

func ActionCmdCheckOut(cCtx *cli.Context) error {
	orgName := strings.TrimSpace(cCtx.Args().Get(0))
	var org *schema.Organization
	var orgId int

	if orgName == "" {
		org, orgId = helper.GetCurrentOrg()
	} else {
		org, orgId = helper.FindOrgByName(orgName)
	}

	if org == nil || orgId == -1 {
		log.Fatal("organization not found")
	}

	helper.UpdateGlobalGitConfig(helper.GlobalGitConfig{
		Name:  org.Name,
		Email: org.Email,
		Org:   org.Org,
	})

	println("current organization " + org.Org)
	return nil
}
