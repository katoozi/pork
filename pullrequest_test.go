package pork

import (
	"testing"

	"github.com/katoozi/go-devops-tools/nap"
)

func TestPullRequest(t *testing.T) {
	GitHubAPI().SetAuth(nap.NewAuthToken("7d748421a0e6e59c9a77a5080638b3950cb90ef3"))
	destRepo = "katoozi/django-simple-poll-app:master"
	sourceRepo = "django-simple-poll-app:dev"
	pullRequestTitle = "test pull request"
	pullRequestMessage = "here it is"
	if err := CreatePullRequest(); err != nil {
		t.Fail()
	}
}
