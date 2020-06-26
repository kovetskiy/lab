package cmd

import (
	"fmt"
	"log"

	"github.com/kovetskiy/lab/internal/git"
	"github.com/kovetskiy/lab/internal/lab"
	"github.com/spf13/cobra"
)

var projectUnprotectCmd = &cobra.Command{
	Use:     "unprotect [project] branch",
	Aliases: []string{"r"},
	Short:   "Unprotect project branch",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		var path string
		var branch string

		switch len(args) {
		case 0:
			log.Fatal("a new name of the project must be specified")
		case 1:
			branch = args[0]

			if !git.InsideGitRepo() {
				log.Fatal(
					"no git repo found, specify two arguments: " +
						"lab project unprotect <project> <branch>",
				)
			}

			ok, err := git.IsRemote("origin")
			if err != nil {
				log.Fatal(err)
			}

			if !ok {
				log.Fatal("no \"origin\" project found, specify project name")
			}

			pathWithNamespace, err := git.PathWithNameSpace("origin")
			if err != nil {
				log.Fatal(err)
			}

			path = pathWithNamespace
		case 2:
			path = args[0]
			branch = args[1]
		default:
			log.Fatal("too many arguments passed")
		}

		project, err := lab.FindProject(path)
		if err != nil {
			log.Fatal(err)
		}

		err = lab.BranchUnprotect(project.ID, branch)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Branch %q has been unprotected on %s\n", branch, path)
	},
}

func init() {
	projectCmd.AddCommand(projectUnprotectCmd)
}
