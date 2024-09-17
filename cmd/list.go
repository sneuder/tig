package cmd

import (
	"clit-git/action"

	"github.com/urfave/cli/v2"
)

func GetCmdList() *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"ls"},
		Usage:   "list organizations",
		Action:  action.ActionCmdList,
	}
}
