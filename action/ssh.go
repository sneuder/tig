package action

import (
	"clit-git/helper"
	"log"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func ActionCmdSshBuild(cCtx *cli.Context) error {
	orgName := cCtx.Args().Get(0)

	if orgName == "" {
		log.Fatal("provide an organization name")
	}

	org, _ := helper.FindOrgByName(orgName)

	if org == nil {
		log.Fatal("organization not found")
	}

	helper.CreateSSHKey(*org)
	println("organization ssh credentials re-built")

	return nil
}

func ActionCmdSshAdd(cCtx *cli.Context) error {
	orgName := cCtx.Args().Get(0)

	if orgName == "" {
		log.Fatal("provide an organization name")
	}

	org, _ := helper.FindOrgByName(orgName)

	if org == nil {
		log.Fatal("organization not found")
	}

	fullPathFolderOrgIdRsa := helper.GetFolderPathOrgIdRsa(org.Org)
	err := exec.Command("ssh-add", fullPathFolderOrgIdRsa).Run()

	if err != nil {
		log.Fatal("could not add ssh public key for " + org.Org + " organization")
	}

	println("ssh public key added")
	return nil
}
