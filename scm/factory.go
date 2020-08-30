package scm

import (
	"fmt"
	"github.com/packagrio/go-common/errors"
	"github.com/packagrio/go-common/pipeline"
)

func Create(scmType string, pipelineData *pipeline.Data) (Interface, error) {

	var scm Interface
	switch scmType {
	case "bitbucket":
		scm = new(scmBitbucket)
	case "github":
		scm = new(scmGithub)
	default:
		return nil, errors.ScmUnspecifiedError(fmt.Sprintf("Unknown Scm Type: %s", scmType))
	}

	if err := scm.Init(pipelineData); err != nil {
		return nil, err
	}
	return scm, nil
}
