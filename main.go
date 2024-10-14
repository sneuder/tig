package main

import (
	"clit-git/cmd"
	"clit-git/setting"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	setting.SetSettings()

	app := &cli.App{
		Commands: getCommands(),
		Version:  os.Getenv("VERSION"),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func getCommands() []*cli.Command {
	return []*cli.Command{
		cmd.GetCmdAdd(),
		cmd.GetCmdCheckOut(),
		cmd.GetCmdRemove(),
		cmd.GetCmdList(),
		cmd.GetCmdSsh(),
	}
}
