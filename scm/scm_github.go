package scm

import (
	"fmt"
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

// see https://github.com/actions/toolkit/blob/main/docs/commands.md

func (g *scmGithub) SetEnvironmentalVariable(name string, value string) error {
	fmt.Printf("\n::set-env name=%s::%s\n", name, value)
	return nil
}

// To prepend a string to PATH
func (g *scmGithub) AddPath(path string) error {
	fmt.Printf("\n::add-path::%s\n", path)
	return nil
}

// To set an output for the step
func (g *scmGithub) SetOutput(name string, value string) error {
	fmt.Printf("\n::set-output name=%s::%s\n", name, value)
	return nil
}

// To mask a value in the logs
func (g *scmGithub) MaskSecret(secret string) error {
	fmt.Printf("\n::add-mask::%s\n", secret)
	return nil
}
