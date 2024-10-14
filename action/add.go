package action

import (
	"clit-git/helper"

	"github.com/urfave/cli/v2"
)

const (
	ORG_FILE_NAME = "organization"
)

func ActionCmdAdd(cCtx *cli.Context) error {
	organization := helper.BuildOrganizationData(cCtx)
	helper.SaveOrganization(organization)
	helper.SaveOrganizationInConfig()

	helper.CreateSSHKey(organization)
	helper.UpdateGlobalGitConfig(helper.GlobalGitConfig{
		Name:  organization.Name,
		Email: organization.Email,
		Org:   organization.Org,
	})

	return nil
}
