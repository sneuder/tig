package cmd

import (
	"clit-git/action"

	"github.com/urfave/cli/v2"
)

func GetCmdRemove() *cli.Command {
	return &cli.Command{
		Name:    "remove",
		Aliases: []string{"r"},
		Usage:   "remove an organization by name",
		Action:  action.ActionCmdRemove,
	}
}
