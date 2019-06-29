package pork

import "testing"

func TestGetRepoReademe(t *testing.T) {
	if err := GetRepoReadme("katoozi/go-basic-topics"); err != nil {
		t.Fail()
	}
}
