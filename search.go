package pork

import (
	"fmt"

	"github.com/spf13/cobra"
)

type SearchResults struct {
	HomeURL    string `json:"html_url"`
	FullName   string `json:"full_name"`
	Name       string `json:"name"`
	ForksCount int    `json:"forks_count"`
}

type SearchReponse struct {
	Results           []*SearchResults `json:"items"`
	TotalCount        int              `json:"total_count"`
	IncompleteResults bool             `json:"incomplete_results"`
}

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
