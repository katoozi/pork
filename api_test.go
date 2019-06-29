package pork

import "testing"

func TestGitHubApi(t *testing.T) {
	if GitHubAPI() == nil {
		t.Fail()
	}
}
