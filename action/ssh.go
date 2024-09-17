package action

import (
	"clit-git/helper"
	"log"

	"github.com/sneuder/filesystem"
	"github.com/urfave/cli/v2"
)

func ActionCmdSsh(cCtx *cli.Context) error {
	orgName := cCtx.Args().Get(0)

	if orgName == "" {
		log.Fatal("provide an organization name")
	}

	org, _ := helper.FindOrgByName(orgName)

	if org == nil {
		log.Fatal("organization not found")
	}

	IdRsaPub, _ := filesystem.Read(helper.GetFolderPathOrgIdRsaPub(orgName))
	IdRsa, _ := filesystem.Read(helper.GetFolderPathOrgIdRsa(orgName))

	println("Public key")
	println(IdRsaPub)

	println("\nPrivate key")
	println(IdRsa)

	return nil
}
