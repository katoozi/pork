package pork

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/katoozi/go-devops-tools/nap"

	"github.com/spf13/cobra"
)

type ReadmeResponse struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

// DocsCmd will read the repo READEME.md file
var DocsCmd = &cobra.Command{
	Use:   "docs",
	Short: "read the documentation for a repo",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatal("You must provide repo argument")
		}
		if err := GetRepoReadme(args[0]); err != nil {
			log.Fatalln("Failed to read docs: ", err)
		}
	},
}

// GetRepoReadme will return the README.md file of the repo
func GetRepoReadme(repo string) error {
	values := strings.Split(repo, "/")
	if len(values) != 2 {
		return fmt.Errorf("repository name must be in this format: owner/project")
	}
	return GitHubAPI().Call("docs", map[string]string{
		"owner":   values[0],
		"project": values[1],
	})
}

// ReadmeSuccess will execute when request was success
func ReadmeSuccess(resp *http.Response, _ interface{}) error {
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	respContent := &ReadmeResponse{}
	json.Unmarshal(content, respContent)
	buff, err := base64.StdEncoding.DecodeString(respContent.Content)
	if err != nil {
		return err
	}
	fmt.Println(string(buff))
	return nil
}

// ReadmeDefaultRouter is the default router for Read Readme file
func ReadmeDefaultRouter(resp *http.Response, _ interface{}) error {
	return fmt.Errorf("status code: %d", resp.StatusCode)
}

// GetReadmeRestResource create nap.RestResource for Read Readme.md file from repository
func GetReadmeRestResource() *nap.RestResource {
	Readrouter := nap.NewRouter()
	Readrouter.DefaultRouter = ReadmeDefaultRouter
	Readrouter.RegisterFunc(200, ReadmeSuccess)
	docs := nap.NewResource("/repos/{{.owner}}/{{.project}}/readme", "GET", Readrouter)
	return docs
}
