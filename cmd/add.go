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
			Name:    constant.Org,
			Aliases: []string{"o"},
			Usage:   "organization name",
		},
		&cli.StringFlag{
			Name:    constant.Email,
			Aliases: []string{"e"},
			Usage:   "email",
		},
		&cli.StringFlag{
			Name:     constant.Name,
			Aliases:  []string{"n"},
			Usage:    "user name",
			Required: true,
		},
		&cli.StringFlag{
			Name:     constant.Platform,
			Aliases:  []string{"p"},
			Usage:    "repository platform",
			Required: true,
		},
	}
}
