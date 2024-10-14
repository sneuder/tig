package helper

import (
	"clit-git/constant"
	"clit-git/schema"
	"encoding/json"
	"log"
	"os/exec"
	"strings"

	"github.com/sneuder/filesystem"
	"github.com/urfave/cli/v2"
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

func BuildOrganizationData(cCtx *cli.Context) schema.Organization {
	platform, platFormExists := GetPlatform(cCtx.String(constant.CmdAddFlags.Platform))

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

func SaveOrganization(organization schema.Organization) {
	organizations := GetOrgs()
	if foundOrg, _ := FindOrgByName(organization.Org); foundOrg != nil {
		log.Fatal("this organization already exists")
	}

	organizations = append(organizations, organization)
	CreateOrgFile(organizations)
}

func SaveOrganizationInConfig() {
	organizations := GetOrgs()
	var directives []filesystem.FileDirective

	for _, organization := range organizations {
		fullPathIdRsa := GetFolderPathOrgIdRsa(organization.Org)
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
		Path:       GetFolderPathSHH(),
		Directives: directives,
	}

	if _, err := filesystem.BuildFile(fileInfo, true); err != nil {
		log.Fatal(err)
	}
}

func RemoveOrgByName(orgName string) {
	orgs := GetOrgs()
	org, orgIndex := FindOrgByName(orgName)

	if org == nil {
		log.Fatal("organization not found")
	}

	orgs = append(orgs[:orgIndex], orgs[orgIndex+1:]...)
	CreateOrgFile(orgs)
}

func GetCurrentOrg() string {
	cmd := exec.Command("bash", "-c", "git config --global --list | grep 'user.org' | cut -d '=' -f 2")
	out, _ := cmd.Output()
	return strings.TrimSpace(string(out))
}
