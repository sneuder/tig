package cmd

import (
	"clit-git/action"

	"github.com/urfave/cli/v2"
)

func GetCmdExport() *cli.Command {
	return &cli.Command{
		Name:    "export",
		Aliases: []string{"e"},
		Usage:   "export org settings",
		Action:  action.ActionCmdExport,
	}
}
