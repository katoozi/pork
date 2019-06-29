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

type PullRequestpayload struct {
	Title        string `json:"title"`
	Message      string `json:"message"`
	SourceBranch string `json:"head"`
	DestBranch   string `json:"base"`
	Modify       bool   `json:"maintainer_can_modify"`
}

type PullRequestResponse struct {
	URL string `json:"html"`
}

// PullRequestCmd is a command that create pull request
var PullRequestCmd = &cobra.Command{
	Use:   "pullrequest",
	Short: "Create Pull Request",
	Run: func(cmd *cobra.Command, args []string) {
		if err := CreatePullRequest(); err != nil {
			log.Fatal("Failed to create pull request.")
		}
	},
}

var (
	sourceRepo         string
	destRepo           string
	pullRequestTitle   string
	pullRequestMessage string
)

func init() {
	PullRequestCmd.PersistentFlags().StringVarP(&sourceRepo, "source", "s", "", "Source Repo")
	PullRequestCmd.PersistentFlags().StringVarP(&destRepo, "destination", "d", "", "Destination Repo")
	PullRequestCmd.PersistentFlags().StringVarP(&pullRequestTitle, "title", "t", "Basic Pull Request", "Pull Request Title")
	PullRequestCmd.PersistentFlags().StringVarP(&pullRequestMessage, "message", "m", "Basic Pull Request", "Pull Request Message")
}

// CreatePullRequest is function that send pullrequest request
func CreatePullRequest() error {
	sourceValues := strings.Split(sourceRepo, ":")
	if !(len(sourceValues) == 1 || len(sourceValues) == 2) {
		return fmt.Errorf("Source repository must in the format [owner:]branch")
	}
	destBranchValues := strings.Split(destRepo, ":")
	if len(destBranchValues) != 2 {
		return fmt.Errorf("Destionation repo must be in this format. owner/project")
	}
	destValues := strings.Split(destBranchValues[0], "/")
	if len(destValues) != 2 {
		return fmt.Errorf("Destionation repo must be in the format owner/project")
	}
	payload := &PullRequestpayload{
		Title:        pullRequestTitle,
		Message:      pullRequestMessage,
		SourceBranch: sourceRepo,
		DestBranch:   destBranchValues[1],
		Modify:       true,
	}
	return GitHubAPI().Call("pullrequest", map[string]string{
		"owner":   destValues[0],
		"project": destValues[1],
	}, payload)
}

func PullRequestSuccess(resp *http.Response, _ interface{}) error {
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	pullRequetsResponse := &PullRequestResponse{}
	if err := json.Unmarshal(content, pullRequetsResponse); err != nil {
		return err
	}
	fmt.Printf("Create Pull Request: %s", pullRequetsResponse.URL)
	return nil
}

func PullRequestDefaultRouter(resp *http.Response, _ interface{}) error {
	return fmt.Errorf("status code: %d", resp.StatusCode)
}

func GetPullReuqestRestResourse() *nap.RestResource {
	router := nap.NewRouter()
	router.DefaultRouter = PullRequestDefaultRouter
	router.RegisterFunc(201, PullRequestSuccess)
	return nap.NewResource("/repos/{{.owner}}/{{.project}}/pulls", "POST", router)
}
