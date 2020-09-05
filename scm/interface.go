package scm

import (
	"github.com/packagrio/go-common/pipeline"
)

// Create mock using:
// mockgen -source=scm/interface.go -destination=scm/mock/mock_scm.go
type Interface interface {

	// init method will generate an authenticated client that can be used to comunicate with Scm
	// MUST set pipelineData.GitParentPath
	Init(pipelineData *pipeline.Data) error

	// To set an environment variable for future out of process steps
	SetEnvironmentalVariable(name string, value string) error

	// To prepend a string to PATH
	AddPath(path string) error

	// To set an output for the step
	SetOutput(name string, value string) error

	// To mask a value in the logs
	MaskSecret(secret string) error
}
