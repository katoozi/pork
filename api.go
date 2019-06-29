package pork

import (
	"github.com/katoozi/go-devops-tools/nap"
	"github.com/spf13/viper"
)

var api *nap.API

// GitHubAPI is singlton function that return github api type
func GitHubAPI() *nap.API {
	if api == nil {
		api = nap.NewAPI("https://api.github.com")
		token := viper.GetString("token")
		api.SetAuth(nap.NewAuthToken(token))
		api.AddResource("fork", GetForkResource())
		api.AddResource("search", GetSearchResource())
		api.AddResource("docs", GetReadmeRestResource())
		api.AddResource("pullrequest", GetPullReuqestRestResourse())
	}
	return api
}
