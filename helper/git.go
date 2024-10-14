package helper

import "os/exec"

type GlobalGitConfig struct {
	Email string
	Name  string
	Org   string
}

func UpdateGlobalGitConfig(globalConfig GlobalGitConfig) {
	exec.Command("git", "config", "--global", "user.name", globalConfig.Name).Run()
	exec.Command("git", "config", "--global", "user.email", globalConfig.Email).Run()
	exec.Command("git", "config", "--global", "user.org", globalConfig.Org).Run()
}
