package pork

import (
	"testing"

	"github.com/katoozi/go-devops-tools/nap"
)

func TestForkRepo(t *testing.T) {
	GitHubAPI().SetAuth(nap.NewAuthToken("7d748421a0e6e59c9a77a5080638b3950cb90ef3"))
	if err := ForkRepo("katoozi/pork"); err == nil {
		t.Fail()
	}
}
