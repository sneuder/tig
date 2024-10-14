package cmd

import (
	"clit-git/action"

	"github.com/urfave/cli/v2"
)

func GetCmdList() *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"ls"},
		Flags:   getFlagsCmdList(),
		Usage:   "list organizations",
		Action:  action.ActionCmdList,
	}
}

func getFlagsCmdList() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    "platform",
			Aliases: []string{"p"},
			Usage:   "shows up the available platforms",
			Value:   true,
		},
	}
}
