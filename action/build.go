package action

import (
	"clit-git/constant"
	"clit-git/helper"
	"clit-git/schema"
	"encoding/json"
	"log"

	"github.com/urfave/cli/v2"
)

func ActionCmdBuild(cCtx *cli.Context) error {
	var organizations = []schema.Organization{}

	exists := helper.CheckJSONFile(constant.ORG_FILE_NAME, ".")

	if !exists {
		log.Fatal("missing organization.json file")
	}

	organizationsData, _ := helper.ReadJSONFile(constant.ORG_FILE_NAME, ".")
	if err := json.Unmarshal(organizationsData, &organizations); err != nil {
		log.Fatal(err)
	}

	helper.RemoveJSONFile(constant.ORG_FILE_NAME, helper.GetFolderPathCLI())

	for _, org := range organizations {
		buildOrganization(org)
	}

	return nil
}
