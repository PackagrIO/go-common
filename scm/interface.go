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
}
