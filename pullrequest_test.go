package pork

import "testing"

func TestPullRequest(t *testing.T) {
	if err := PullRequest("myrepo"); err != nil {
		t.Fail()
	}
}
