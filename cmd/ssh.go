package cmd

import (
	"clit-git/action"

	"github.com/urfave/cli/v2"
)

func GetCmdSsh() *cli.Command {
	return &cli.Command{
		Name:    "ssh",
		Aliases: []string{"s"},
		Subcommands: []*cli.Command{
			{
				Name:    "build",
				Aliases: []string{"b"},
				Usage:   "build the organization's ssh credentials",
				Action:  action.ActionCmdSshBuild,
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add the ssh credentials to the ssh list",
				Action:  action.ActionCmdSshAdd,
			},
		},
	}
}
