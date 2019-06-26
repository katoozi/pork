package pork

import (
	"log"

	"github.com/spf13/cobra"
)

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
	return nil
}
