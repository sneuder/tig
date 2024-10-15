package action

import (
	"clit-git/constant"
	"clit-git/helper"
	"clit-git/schema"
	"log"

	"github.com/urfave/cli/v2"
)

const (
	ORG_FILE_NAME = "organization"
)

func ActionCmdAdd(cCtx *cli.Context) error {
	organization := buildOrganizationData(cCtx)
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

//

func buildOrganizationData(cCtx *cli.Context) schema.Organization {
	platform, platFormExists := helper.GetPlatform(cCtx.String(constant.Platform))

	if !platFormExists {
		log.Fatal("the platform does not exists")
	}

	return schema.Organization{
		Org:      cCtx.String(constant.Org),
		Name:     cCtx.String(constant.Name),
		Email:    cCtx.String(constant.Email),
		Platform: platform,
	}
}
