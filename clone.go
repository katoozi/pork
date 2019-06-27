package pork

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// CloneCmd is a command that clone Github repo into destination
var CloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "clone repo from Github",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatalln("You must provide repo")
		}
		if err := CloneRepo(args[0], ref, create); err != nil {
			log.Fatalln("error when clone repo: ", err)
		}
	},
}

// CloneRepo will clone github repo into destination
func CloneRepo(repo string, ref string, shouldCreate bool) error {
	repository, err := NewGHRepo(repo)
	if err != nil {
		return err
	}
	if err := repository.Clone(viper.GetString("location")); err != nil {
		return err
	}
	fmt.Printf("Cloned Repo to:%s\n", repository.RepoDir)
	if err := repository.Checkout(ref, shouldCreate); err != nil {
		return err
	}
	return nil
}

var (
	ref    string
	create bool
)

func init() {
	CloneCmd.PersistentFlags().StringVar(&ref, "ref", "master", "specific refrence to check out")
	CloneCmd.PersistentFlags().BoolVar(&create, "create", false, "create refrence if it does not exist.")
}
