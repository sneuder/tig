package action

import (
	"clit-git/helper"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
)

func ActionCmdList(cCtx *cli.Context) error {
	orgs := helper.GetOrgs()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Organization", "Name", "Email", "Platform"})

	for _, org := range orgs {
		table.Append([]string{org.Org, org.Name, org.Email, org.Platform})
	}

	table.Render()
	return nil
}
