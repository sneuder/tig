package cmd

import (
	"clit-git/action"
	"clit-git/constant"

	"github.com/urfave/cli/v2"
)

func GetCmdAdd() *cli.Command {
	return &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Flags:   getFlagsCmdAdd(),
		Usage:   "add an organization",
		Action:  action.ActionCmdAdd,
	}
}

func getFlagsCmdAdd() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    constant.CmdAddFlags.Org,
			Aliases: []string{"o"},
			Usage:   "organization name",
		},
		&cli.StringFlag{
			Name:    constant.CmdAddFlags.Email,
			Aliases: []string{"e"},
			Usage:   "user name",
		},
		&cli.StringFlag{
			Name:     constant.CmdAddFlags.Name,
			Aliases:  []string{"n"},
			Usage:    "user name",
			Required: true,
		},
		&cli.StringFlag{
			Name:     constant.CmdAddFlags.Platform,
			Aliases:  []string{"p"},
			Usage:    "repository platform",
			Required: true,
		},
	}
}
