package main

import (
	"os"
	usecases "tig/application/use-cases"
	"tig/infrastructure/consts"
	"tig/infrastructure/persistence/repositories"
	infrastructureServices "tig/infrastructure/services"
	"tig/representation/cli/cmds"
	"tig/representation/cli/settings"
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

	var addUserCmd schemas.CmdSchema = cmds.NewAddUserCmd(addNewUserUseCase)
	var removeUserCmd schemas.CmdSchema = cmds.NewRemoveUserCmd(removeUserUseCase)
	var listUserCmd schemas.CmdSchema = cmds.NewListUsersCmd(getRepositoryHostsUseCase, newGetUsersUseCase, getUserGitConfigUseCase, printService)
	var checkoutUserCmd schemas.CmdSchema = cmds.NewCheckoutUserCmd(checkoutUserUseCase)

	commands := settings.BuildCLI(settings.CmdCollections{
		User: []schemas.CmdSchema{addUserCmd, removeUserCmd, listUserCmd, checkoutUserCmd},
	})

	app := &urfave.App{
		Commands: commands,
		Version:  os.Getenv("VERSION"),
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
