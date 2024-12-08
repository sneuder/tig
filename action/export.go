package action

import (
	"clit-git/helper"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func ActionCmdExport(cCtx *cli.Context) error {
	orgFilePath := helper.GetOrgFilepath()
	exec.Command("cp", orgFilePath, ".").Run()
	return nil
}
