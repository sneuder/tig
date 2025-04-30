package action

import (
	"clit-git/helper"
	"clit-git/schema"
	"log"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func ActionCmdSshBuild(cCtx *cli.Context) error {
	orgName := cCtx.Args().Get(0)
	var org *schema.Organization

	if orgName == "" {
		org, _ = helper.GetCurrentOrg()
	} else {
		org, _ = helper.FindOrgByName(orgName)
	}

	if org == nil {
		log.Fatal("organization not found")
	}

	helper.CreateSSHKey(*org)
	println("organization ssh credentials re-built for " + org.Org + " organization")

	return nil
}

func ActionCmdSshAdd(cCtx *cli.Context) error {
	orgName := cCtx.Args().Get(0)

	var org *schema.Organization

	if orgName == "" {
		org, _ = helper.GetCurrentOrg()
	} else {
		org, _ = helper.FindOrgByName(orgName)
	}

	if org == nil {
		log.Fatal("organization not found")
	}

	fullPathFolderOrgIdRsa := helper.GetFolderPathOrgIdRsa(org.Org)
	err := exec.Command("ssh-add", fullPathFolderOrgIdRsa).Run()

	if err != nil {
		log.Fatal("could not add ssh public key for " + org.Org + " organization")
	}

	println("ssh public key added for " + org.Org + " organization")
	return nil
}
