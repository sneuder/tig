package action

import (
	"clit-git/helper"
	"log"
	"strings"

	"github.com/urfave/cli/v2"
)

func ActionCmdCheckOut(cCtx *cli.Context) error {
	orgName := strings.TrimSpace(cCtx.Args().Get(0))
	if orgName == "" {
		log.Fatal("please provide an organization name")
	}

	org, _ := helper.FindOrgByName(orgName)
	if org == nil {
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
