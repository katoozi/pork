package pork

import (
	"github.com/spf13/cobra"
)

// PullRequestCmd is a command that create pull request
var PullRequestCmd = &cobra.Command{
	Use:   "pullrequest",
	Short: "Create Pull Request",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// PullRequest is function that send pullrequest request
func PullRequest(repo string) error {
	return nil
}
