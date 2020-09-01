package scm

import (
	"github.com/packagrio/go-common/pipeline"
)

type scmGithub struct {
	PipelineData *pipeline.Data
}

func (g *scmGithub) Init(pipelineData *pipeline.Data) error {
	g.PipelineData = pipelineData

	//if _, present := os.LookupEnv("GITHUB_TOKEN"); !present {
	//	return errors.ScmAuthenticationFailed("Missing github access token")
	//}
	//if g.Config.IsSet("scm_git_parent_path") {
	//	g.PipelineData.GitParentPath = g.Config.GetString("scm_git_parent_path")
	//	os.MkdirAll(g.PipelineData.GitParentPath, os.ModePerm)
	//} else {
	//	dirPath, _ := ioutil.TempDir("", "")
	//	g.PipelineData.GitParentPath = dirPath
	//}

	//if client != nil {
	//	//primarily used for testing.
	//	g.Client = github.NewClient(client)
	//} else {
	//	ctx := context.Background()
	//	ts := oauth2.StaticTokenSource(
	//		&oauth2.Token{AccessToken: g.Config.GetString("scm_github_access_token")},
	//	)
	//	tc := oauth2.NewClient(ctx, ts)
	//
	//	//TODO: autopaginate turned on.
	//	g.Client = github.NewClient(tc)
	//}
	//
	//if g.Config.IsSet("scm_github_api_endpoint") {
	//
	//	apiUrl, aerr := url.Parse(g.Config.GetString("scm_github_api_endpoint"))
	//	if aerr != nil {
	//		return aerr
	//	}
	//	g.Client.BaseURL = apiUrl
	//}

	return nil
}
