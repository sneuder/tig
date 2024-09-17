package cmd

import (
	"clit-git/action"

	"github.com/urfave/cli/v2"
)

func GetCmdCheckOut() *cli.Command {
	return &cli.Command{
		Name:    "checkout",
		Aliases: []string{"c"},
		Usage:   "remove organization",
		Action:  action.ActionCmdCheckOut,
	}
}
