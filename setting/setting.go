package setting

import (
	"clit-git/helper"
	"os/exec"
)

func SetSettings() {
	setFolders()
}

func setFolders() {
	exec.Command("mkdir", "-p", helper.GetFolderPathCLI()).Run()
	exec.Command("mkdir", "-p", helper.GetFolderPathSHH()).Run()
}
