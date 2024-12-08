package cmd

import (
	"clit-git/action"

	"github.com/urfave/cli/v2"
)

func GetCmdBuild() *cli.Command {
	return &cli.Command{
		Name:    "build",
		Aliases: []string{"b"},
		Usage:   "build multiple orgs from organization.json file in current path",
		Action:  action.ActionCmdBuild,
	}
}
