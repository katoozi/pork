package pork

import "testing"

func TestCloneRepo(t *testing.T) {
	if err := CloneRepo("katoozi/pork", "develop", false); err == nil {
		t.Fail()
	}
}
