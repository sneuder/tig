package helper

import (
	"clit-git/constant"
	"clit-git/schema"
	"encoding/json"
	"log"
	"os/exec"
)

const (
	ORG_FILE_NAME = "organization"
)

func GetOrgs() []schema.Organization {
	jsonFilePath := GetFolderPathCLI()
	jsonFileExists := CheckJSONFile(ORG_FILE_NAME, jsonFilePath)

	var organizations = []schema.Organization{}
	if jsonFileExists {
		organizationsData, _ := ReadJSONFile(ORG_FILE_NAME, jsonFilePath)
		if err := json.Unmarshal(organizationsData, &organizations); err != nil {
			log.Fatal(err)
		}
	}

	return organizations
}

func CreateOrgFile(organizations []schema.Organization) {
	jsonFilePath := GetFolderPathCLI()
	CreateJSONFile(organizations, ORG_FILE_NAME, jsonFilePath)
}

func RemoveOrgFolder(orgName string) {
	orgFolderPath := GetFolderPathOrg(orgName)
	exec.Command("rm", "-r", orgFolderPath).Run()
}

func GetPlatform(platformBase string) (string, bool) {
	platform, platFormExists := constant.PlatFormOptions[platformBase]
	return platform, platFormExists
}

func FindOrgByName(orgName string) (*schema.Organization, int) {
	organizations := GetOrgs()

	for i, org := range organizations {
		if org.Org == orgName {
			return &org, i
		}
	}

	return nil, -1
}
