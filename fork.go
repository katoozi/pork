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

// ForkResponse is github response schema
type ForkResponse struct {
	CloneURL string `json:"clone_url"`
	FullName string `json:"full_name"`
}

// ForkCmd  is a command that will fork given repo on github
var ForkCmd = &cobra.Command{
	Use:   "fork",
	Short: "fork a repo on Github",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatalln("You must provide repo")
		}
		if err := ForkRepo(args[0]); err != nil {
			log.Fatalln("error when forking repo: ", err)
		}
	},
}

// ForkRepo will fork repo
func ForkRepo(repo string) error {
	values := strings.Split(repo, "/")
	if len(values) != 2 {
		return fmt.Errorf("Github repo must be in this format => owner/project")
	}
	return GitHubAPI().Call("fork", map[string]string{
		"owner": values[0],
		"repo":  values[1],
	}, nil)
}

// GetForkResource will set the rest api options for fork command
func GetForkResource() *nap.RestResource {
	forkRouter := nap.NewRouter()
	forkRouter.RegisterFunc(202, ForkSuccess)
	forkRouter.RegisterFunc(401, func(_ *http.Response, _ interface{}) error {
		return fmt.Errorf("You must set Authentication token")
	})
	fork := nap.NewResource("/repos/{{.owner}}/{{.repo}}/forks", "POST", forkRouter)
	return fork
}

// ForkSuccess will handle the succes request of fork command
func ForkSuccess(resp *http.Response, _ interface{}) error {
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	respContent := ForkResponse{}
	json.Unmarshal(content, &respContent)
	fmt.Printf("Forked to repository: %s\n", respContent.FullName)
	return nil
}
