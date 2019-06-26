package pork

import (
	"fmt"

	"github.com/spf13/cobra"
)

// SearchCmd is a subcommand of pork
var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "search for github repositories by keyword",
	Run: func(cmd *cobra.Command, args []string) {
		reposList := SearchByKeyword(args)
		for _, repo := range reposList {
			fmt.Println(repo)
		}
	},
}

// SearchByKeyword will take slice of keywords and return the github repos
func SearchByKeyword(keywords []string) []string {
	return []string{"myrepo"}
}
