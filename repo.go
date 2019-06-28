package pork

import (
	"fmt"
	"path/filepath"
	"strings"

	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"

	git "gopkg.in/src-d/go-git.v4"
)

// GHRepo is a struct that hold github repo info
type GHRepo struct {
	RepoDir string
	owner   string
	project string
	repo    *git.Repository
}

// NewGHRepo is GHRepo factory function
func NewGHRepo(repo string) (*GHRepo, error) {
	values := strings.Split(repo, "/")
	if len(values) != 2 {
		return nil, fmt.Errorf("repo must in this format owner/project")
	}
	return &GHRepo{
		owner:   values[0],
		project: values[1],
	}, nil
}

// RepoURL will generate github repo clone url
func (g *GHRepo) RepoURL() string {
	return fmt.Sprintf("https://github.com/%s/%s.git", g.owner, g.project)
}

// Clone will clone the github repo
func (g *GHRepo) Clone(dest string) error {
	fullPath := filepath.Join(dest, fmt.Sprintf("%s-%s", g.owner, g.project))
	repo, err := git.PlainClone(fullPath, false, &git.CloneOptions{
		URL: g.RepoURL(),
	})
	if err != nil {
		return err
	}
	g.repo = repo
	g.RepoDir = fullPath
	return nil
}

// Checkout can create branch or just check to existence branch
func (g *GHRepo) Checkout(ref string, create bool) error {
	opts := &git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(fmt.Sprintf("refs/heads/%s", ref)),
		Create: create,
	}
	if create {
		head, err := g.repo.Head()
		if err != nil {
			return err
		}
		opts.Hash = head.Hash()
	}
	workTree, err := g.repo.Worktree()
	if err != nil {
		return err
	}
	return workTree.Checkout(opts)
}

// AddUpstream will create upstream for github repo
func (g *GHRepo) AddUpstream(githubRepo *GHRepo) error {
	_, err := g.repo.CreateRemote(&config.RemoteConfig{
		Name: "upstream",
		URLs: []string{githubRepo.RepoURL()},
	})
	return err
}
