package cmd

import (
	"clit-git/action"

	"github.com/urfave/cli/v2"
)

func GetCmdSsh() *cli.Command {
	return &cli.Command{
		Name:    "ssh",
		Aliases: []string{"s"},
		Usage:   "re-build organziation's private and public ssh keys",
		Action:  action.ActionCmdSsh,
	}
}
