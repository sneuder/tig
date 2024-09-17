package action

import (
	"clit-git/constant"
	"clit-git/helper"
	"clit-git/schema"
	"log"
	"os/exec"
	"path/filepath"

	"github.com/sneuder/filesystem"
	"github.com/urfave/cli/v2"
)

const (
	ORG_FILE_NAME = "organization"
)

func ActionCmdAdd(cCtx *cli.Context) error {
	organization := buildOrganizationData(cCtx)
	saveOrganization(organization)
	saveOrganizationInConfig()

	createSSHKey(organization)
	updateGlobalGitConfig(GlobalGitConfig{
		Name:  organization.Name,
		Email: organization.Email,
		Org:   organization.Org,
	})

	return nil
}

func buildOrganizationData(cCtx *cli.Context) schema.Organization {
	platform, platFormExists := helper.GetPlatform(cCtx.String(constant.CmdAddFlags.Platform))

	if !platFormExists {
		log.Fatal("the platform does not exists")
	}

	return schema.Organization{
		Org:      cCtx.String(constant.CmdAddFlags.Org),
		Name:     cCtx.String(constant.CmdAddFlags.Name),
		Email:    cCtx.String(constant.CmdAddFlags.Email),
		Platform: platform,
	}
}

func saveOrganization(organization schema.Organization) {
	organizations := helper.GetOrgs()
	if foundOrg, _ := helper.FindOrgByName(organization.Org); foundOrg != nil {
		log.Fatal("this organization already exists")
	}

	organizations = append(organizations, organization)
	helper.CreateOrgFile(organizations)
}

func saveOrganizationInConfig() {
	organizations := helper.GetOrgs()
	var directives []filesystem.FileDirective

	for _, organization := range organizations {
		fullPathIdRsa := helper.GetFolderPathOrgIdRsa(organization.Org)
		directive := []filesystem.FileDirective{
			{Content: "Host " + organization.Org, Indent: 0, NewLine: true},
			{Content: "HostName " + organization.Platform, Indent: 2, NewLine: false},
			{Content: "User git", Indent: 2, NewLine: false},
			{Content: "IdentityFile " + fullPathIdRsa, Indent: 2, NewLine: false},
		}

		directives = append(directives, directive...)

	}

	fileInfo := filesystem.FileInfo{
		Name:       "config",
		Path:       helper.GetFolderPathSHH(),
		Directives: directives,
	}

	if _, err := filesystem.BuildFile(fileInfo, true); err != nil {
		log.Fatal(err)
	}
}

func createSSHKey(organization schema.Organization) {
	sshFileName := "id_rsa"
	fullPathFolderOrg := helper.GetFolderPathOrg(organization.Org)

	exec.Command("rm", "-r", fullPathFolderOrg).Run()
	exec.Command("mkdir", "-p", fullPathFolderOrg).Run()

	fullPathNameSHH := filepath.Join(fullPathFolderOrg, sshFileName)
	cmd := exec.Command("ssh-keygen", "-t", "rsa", "-b", "4096", "-C", organization.Email, "-f", fullPathNameSHH, "-N", "")

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	exec.Command("eval", "$(ssh-agent -s)").Run()
	exec.Command("ssh-add", fullPathNameSHH).Run()
}
