package pork

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// DocsCmd will read the repo READEME.md file
var DocsCmd = &cobra.Command{
	Use:   "docs",
	Short: "read the documentation for a repo",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatal("You must provide repo argument")
		}
		repoReadme := GetRepoReadme(args[0])
		fmt.Println(repoReadme)
	},
}

// GetRepoReadme will return the README.md file of the repo
func GetRepoReadme(repo string) string {
	return repo
}
