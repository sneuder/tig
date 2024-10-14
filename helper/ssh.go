package helper

import (
	"clit-git/schema"
	"log"
	"os/exec"
	"path/filepath"
)

func CreateSSHKey(organization schema.Organization) {
	sshFileName := "id_rsa"
	fullPathFolderOrg := GetFolderPathOrg(organization.Org)

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
