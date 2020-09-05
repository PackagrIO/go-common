package scm

import (
	"github.com/packagrio/go-common/pipeline"
)

type scmBitbucket struct {
	PipelineData *pipeline.Data
}

// configure method will generate an authenticated client that can be used to comunicate with Github
// MUST set @git_parent_path
// MUST set @client field
func (b *scmBitbucket) Init(pipelineData *pipeline.Data) error {
	b.PipelineData = pipelineData

	//if !b.Config.IsSet("scm_bitbucket_username") {
	//	return errors.ScmAuthenticationFailed("Missing bitbucket username")
	//}
	//if !b.Config.IsSet("scm_bitbucket_password") && !b.Config.IsSet("scm_bitbucket_access_token") {
	//	return errors.ScmAuthenticationFailed("Bitbucket app password or access token is required")
	//}
	//if b.Config.IsSet("scm_git_parent_path") {
	//	b.PipelineData.GitParentPath = b.Config.GetString("scm_git_parent_path")
	//	os.MkdirAll(b.PipelineData.GitParentPath, os.ModePerm)
	//} else {
	//	dirPath, _ := ioutil.TempDir("", "")
	//	b.PipelineData.GitParentPath = dirPath
	//}
	//
	//if b.Config.IsSet("scm_bitbucket_password") {
	//	b.Client = bitbucket.NewBasicAuth(b.Config.GetString("scm_bitbucket_username"), b.Config.GetString("scm_bitbucket_password"))
	//} else {
	//	b.Client = bitbucket.NewOAuthbearerToken(b.Config.GetString("scm_bitbucket_access_token"))
	//}
	//if client != nil {
	//	//primarily used for testing.
	//	b.Client.HttpClient = client
	//}

	return nil
}

func (b *scmBitbucket) SetEnvironmentalVariable(name string, value string) error {
	return nil
}

// To prepend a string to PATH
func (b *scmBitbucket) AddPath(path string) error {
	return nil
}

// To set an output for the step
func (b *scmBitbucket) SetOutput(name string, value string) error {
	return nil
}

// To mask a value in the logs
func (b *scmBitbucket) MaskSecret(secret string) error {
	return nil
}
