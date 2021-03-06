package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/kovetskiy/lab/internal/lab"
)

var issueCloseCmd = &cobra.Command{
	Use:     "close [remote] <id>",
	Aliases: []string{"delete"},
	Short:   "Close issue by id",
	Long:    ``,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rn, id, err := parseArgs(args)
		if err != nil {
			log.Fatal(err)
		}

		p, err := lab.FindProject(rn)
		if err != nil {
			log.Fatal(err)
		}

		err = lab.IssueClose(p.ID, int(id))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Issue #%d closed\n", id)
	},
}

func init() {
	issueCloseCmd.MarkZshCompPositionalArgumentCustom(1, "__lab_completion_remote")
	issueCloseCmd.MarkZshCompPositionalArgumentCustom(2, "__lab_completion_issue $words[2]")
	issueCmd.AddCommand(issueCloseCmd)
}
