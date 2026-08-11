package cmds

import (
	usecases "tig/application/use-cases"
	"tig/representation/services"

	"github.com/urfave/cli/v2"
)

type ListDataCmd struct {
	getRepositoryHostsUseCase *usecases.GetRepositoryHostsUseCase
	getUsersUseCase           *usecases.GetUsersUseCase
	getUserGitConfigUseCase   *usecases.GetUserGitConfigUseCase
	printService              *services.PrintService
}

func NewListUsersCmd(
	getRepositoryHostsUseCase *usecases.GetRepositoryHostsUseCase,
	getUsersUseCase *usecases.GetUsersUseCase,
	getUserGitConfigUseCase *usecases.GetUserGitConfigUseCase,
	printService *services.PrintService,
) *ListDataCmd {
	return &ListDataCmd{
		getRepositoryHostsUseCase: getRepositoryHostsUseCase,
		getUsersUseCase:           getUsersUseCase,
		printService:              printService,
		getUserGitConfigUseCase:   getUserGitConfigUseCase,
	}
}

func (ld *ListDataCmd) BuildSetting() *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"ls"},
		Flags:   ld.BuildFlags(),
		Usage:   "list users",
		Action:  ld.BuildAction,
	}
}

func (ld *ListDataCmd) BuildFlags() []cli.Flag {
	return []cli.Flag{}
}

func (ld *ListDataCmd) BuildAction(cCtx *cli.Context) error {
	users := ld.getUsersUseCase.Execute()

	table := [][]string{
		{"User Name", "Email", "Alias", "Repository"},
	}

	for _, user := range users {
		name := user.Name()

		if isGlobal := user.InGlobal(); isGlobal {
			name = "*" + name
		}

		table = append(table, []string{name, user.Email(), user.Alias(), user.RepositoryHost()})
	}

	ld.printService.RenderTable(table)

	return nil
}
