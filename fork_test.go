package pork

import "testing"

func TestForkRepo(t *testing.T) {
	if err := ForkRepo("katoozi/pork"); err == nil {
		t.Fail()
	}
}
