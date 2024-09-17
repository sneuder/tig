package action

import (
	"clit-git/helper"
	"log"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
)

func ActionCmdCheckOut(cCtx *cli.Context) error {
	orgName := strings.TrimSpace(cCtx.Args().Get(0))
	if orgName == "" {
		log.Fatal("please provide an organization name")
	}

	org, _ := helper.FindOrgByName(orgName)
	if org == nil {
		log.Fatal("organization not found")
	}

	updateGlobalGitConfig(GlobalGitConfig{
		Name:  org.Name,
		Email: org.Email,
		Org:   org.Org,
	})

	println("current organization " + org.Name)
	return nil
}

type GlobalGitConfig struct {
	Email string
	Name  string
	Org   string
}

func updateGlobalGitConfig(globalConfig GlobalGitConfig) {
	exec.Command("git", "config", "--global", "user.name", globalConfig.Name).Run()
	exec.Command("git", "config", "--global", "user.email", globalConfig.Email).Run()
	exec.Command("git", "config", "--global", "user.org", globalConfig.Org).Run()
}
