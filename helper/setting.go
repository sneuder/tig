package helper

import (
	"clit-git/constant"
	"log"
	"os"
	"path/filepath"
)

func GetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return homeDir
}

func GetFolderPathCLI() string {
	homeDir := GetHomeDir()
	return homeDir + "/.cli-git" // TODO: change to tig-org
}

func GetFolderPathSHH() string {
	return filepath.Join(GetHomeDir(), ".ssh")
}

func GetFolderPathOrg(org string) string {
	return filepath.Join(GetFolderPathSHH(), org)
}

func GetFolderPathOrgIdRsa(org string) string {
	return filepath.Join(GetFolderPathOrg(org), "id_rsa")
}

func GetFolderPathOrgIdRsaPub(org string) string {
	return filepath.Join(GetFolderPathOrg(org), "id_rsa.pub")
}

func GetOrgFilepath() string {
	return filepath.Join(GetFolderPathCLI(), AddExtension(constant.ORG_FILE_NAME))
}
