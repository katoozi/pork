package pork

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/katoozi/go-devops-tools/nap"

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

func (s *SearchReponse) ToString() string {
	response := ""
	for _, v := range s.Results {
		response += fmt.Sprintf("%s\n", v.FullName)
	}
	return response
}

// SearchCmd is a subcommand of pork
var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "search for github repositories by keyword",
	Run: func(cmd *cobra.Command, args []string) {
		if err := SearchByKeyword(args); err != nil {
			log.Fatalf("Search Failed: %s", err)
		}
	},
}

// SearchByKeyword will take slice of keywords and return the github repos
func SearchByKeyword(keywords []string) error {
	return GitHubAPI().Call("search", map[string]string{
		"query": strings.Join(keywords, "+"),
	})
}

func SearchSuccess(resp *http.Response, _ interface{}) error {
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	searchResponse := SearchReponse{}
	json.Unmarshal(content, &searchResponse)
	fmt.Print(searchResponse.ToString())
	return nil
}

func SearchDefaultRouter(resp *http.Response, _ interface{}) error {
	return fmt.Errorf("status code %d", resp.StatusCode)
}

func GetSearchResource() *nap.RestResource {
	searchRouter := nap.NewRouter()
	searchRouter.DefaultRouter = SearchDefaultRouter
	searchRouter.RegisterFunc(200, SearchSuccess)
	search := nap.NewResource("/search/repositories?q={{.query}}", "GET", searchRouter)
	return search
}
