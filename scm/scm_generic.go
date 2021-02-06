package scm

import (
	"github.com/packagrio/go-common/config"
	"github.com/packagrio/go-common/pipeline"
	"net/http"
)

type scmGeneric struct {
	PipelineData *pipeline.Data
}

// configure method will generate an authenticated client that can be used to comunicate with Github
// MUST set @git_parent_path
// MUST set @client field
func (s *scmGeneric) Init(pipelineData *pipeline.Data, myConfig config.BaseInterface, httpClient *http.Client) error {
	s.PipelineData = pipelineData

	//if !s.Config.IsSet("scm_bitbucket_username") {
	//	return errors.ScmAuthenticationFailed("Missing bitbucket username")
	//}
	//if !s.Config.IsSet("scm_bitbucket_password") && !s.Config.IsSet("scm_bitbucket_access_token") {
	//	return errors.ScmAuthenticationFailed("Bitbucket app password or access token is required")
	//}
	//if s.Config.IsSet("scm_git_parent_path") {
	//	s.PipelineData.GitParentPath = s.Config.GetString("scm_git_parent_path")
	//	os.MkdirAll(s.PipelineData.GitParentPath, os.ModePerm)
	//} else {
	//	dirPath, _ := ioutil.TempDir("", "")
	//	s.PipelineData.GitParentPath = dirPath
	//}
	//
	//if s.Config.IsSet("scm_bitbucket_password") {
	//	s.Client = bitbucket.NewBasicAuth(s.Config.GetString("scm_bitbucket_username"), s.Config.GetString("scm_bitbucket_password"))
	//} else {
	//	s.Client = bitbucket.NewOAuthbearerToken(s.Config.GetString("scm_bitbucket_access_token"))
	//}
	//if client != nil {
	//	//primarily used for testing.
	//	s.Client.HttpClient = client
	//}

	return nil
}

// Generic reps
func (s *scmGeneric) RetrievePayload() (*Payload, error) {
	return nil, nil
}

func (s *scmGeneric) Publish() error {
	return nil
}

func (s *scmGeneric) PublishAssets(releaseData interface{}) error {
	return nil
}

func (s *scmGeneric) Cleanup() error {
	return nil
}

func (s *scmGeneric) SetEnvironmentalVariable(name string, value string) error {
	return nil
}

// To prepend a string to PATH
func (s *scmGeneric) AddPath(path string) error {
	return nil
}

// To set an output for the step
func (s *scmGeneric) SetOutput(name string, value string) error {
	return nil
}

// To mask a value in the logs
func (s *scmGeneric) MaskSecret(secret string) error {
	return nil
}
