package main

import (
	"os"
	usecases "tig/application/use-cases"
	"tig/infrastructure/consts"
	"tig/infrastructure/persistence/repositories"
	infrastructureServices "tig/infrastructure/services"
	cli "tig/representation/cli"
	"tig/representation/schemas"
	representationServices "tig/representation/services"

	urfave "github.com/urfave/cli/v2"
)

var jsonService *infrastructureServices.JsonService = infrastructureServices.NewJsonService()
var fsService *infrastructureServices.FsService = infrastructureServices.NewFsService()

func main() {
	preExecution()

	gitService := infrastructureServices.NewGitService(fsService, jsonService)
	organizationRepository := repositories.NewOrganizationRepository(jsonService, fsService)

	getRepositoryHostsUseCase := usecases.NewGetRepositoryHostsUseCase()
	newGetUsersUseCase := usecases.NewGetUsersUseCase(organizationRepository, gitService)
	addNewUserUseCase := usecases.NewAddNewUserUseCase(organizationRepository, gitService)
	removeUserUseCase := usecases.NewRemoveUserUseCase(organizationRepository, gitService)
	checkoutUserUseCase := usecases.NewCheckoutUserUseCase(organizationRepository, gitService)
	getUserGitConfigUseCase := usecases.NewGetUserGitConfigUseCase(organizationRepository, gitService)

	printService := representationServices.NewPrintService()

	var addUserCli schemas.CLISchema = cli.NewAddUserCli(addNewUserUseCase)
	var removeUserCli schemas.CLISchema = cli.NewRemoveUserCli(removeUserUseCase)
	var listDataCli schemas.CLISchema = cli.NewListDataCli(getRepositoryHostsUseCase, newGetUsersUseCase, getUserGitConfigUseCase, printService)
	var checkoutUserCli schemas.CLISchema = cli.NewCheckoutUserCli(checkoutUserUseCase)

	app := &urfave.App{
		Commands: []*urfave.Command{
			addUserCli.BuildSetting(),
			removeUserCli.BuildSetting(),
			listDataCli.BuildSetting(),
			checkoutUserCli.BuildSetting(),
		},
		Version: os.Getenv("VERSION"),
	}

	app.Run(os.Args)
}

func preExecution() {
	systemFolder := fsService.GetFolderPath([]string{consts.FOLDER})
	fsService.CreateDirectory(systemFolder)

	jsonPath := fsService.GetFolderPath([]string{consts.FOLDER, jsonService.AddExtension(consts.FILE_NAME)})

	if existedFile := fsService.ExistFile(jsonPath); existedFile {
		return
	}

	jsonService.CreateFile("", systemFolder, consts.FILE_NAME)
}
