package pork

import "testing"

func TestGetRepoReademe(t *testing.T) {
	repoReademe := GetRepoReadme("myrepo")
	if repoReademe != "myrepo" {
		t.Fail()
	}
}
