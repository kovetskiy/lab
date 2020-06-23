package cmd

import (
	"fmt"
	"log"

	"github.com/kovetskiy/lab/internal/git"
	"github.com/kovetskiy/lab/internal/lab"
	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
)

var projectRenameCmd = &cobra.Command{
	Use:     "rename [from] to",
	Aliases: []string{"r"},
	Short:   "Rename project",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		var from string
		var to string

		switch len(args) {
		case 0:
			log.Fatal("a new name of the project must be specified")
		case 1:
			if !git.InsideGitRepo() {
				log.Fatal(
					"no git repo found, specify two arguments: " +
						"lab project rename <from> <to>",
				)
			}

			ok, err := git.IsRemote("origin")
			if err != nil {
				log.Fatal(err)
			}

			if !ok {
				log.Fatal("no \"origin\" remote found, specify project name")
			}

			pathWithNamespace, err := git.PathWithNameSpace("origin")
			if err != nil {
				log.Fatal(err)
			}

			from = pathWithNamespace
			to = args[0]
		case 2:
			from = args[0]
			to = args[1]
		default:
			log.Fatal("too many arguments passed")
		}

		project, err := lab.FindProject(from)
		if err != nil {
			log.Fatal(err)
		}

		newProject, err := lab.ProjectEdit(project.ID, &gitlab.EditProjectOptions{
			Name: gitlab.String(to),
			Path: gitlab.String(to),
		})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("project %q has been renamed to %q\n", from, to)

		if git.InsideGitRepo() {
			ok, err := git.IsRemote("origin")
			if err != nil {
				log.Println(err)
				return
			}

			if !ok {
				return
			}

			pathWithNamespace, err := git.PathWithNameSpace("origin")
			if err != nil {
				log.Println(err)
				return
			}

			if pathWithNamespace == project.PathWithNamespace {
				err = git.RemoteSet("origin", newProject.SSHURLToRepo, ".")
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	},
}

func init() {
	projectCmd.AddCommand(projectRenameCmd)
}
