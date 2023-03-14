package git

import (
	"fmt"
	goUtils "github.com/analogj/go-util/utils"
	"github.com/go-git/go-git/v5"
	"github.com/packagrio/go-common/errors"
	gitUrl "github.com/whilp/git-urls"
	"os"
	"path"
	"path/filepath"
)

// Clone a git repo into a local directory.
// Credentials need to be specified by embedding in gitRemote url.
// TODO: this pattern may not work on Bitbucket/GitLab
func GitClone(parentPath string, repositoryName string, gitRemote string) (string, error) {
	absPath, _ := filepath.Abs(path.Join(parentPath, repositoryName))

	if !goUtils.FileExists(absPath) {
		os.MkdirAll(absPath, os.ModePerm)
	} else {
		return "", errors.ScmFilesystemError(fmt.Sprintf("The local repository path already exists, this should never happen. %s", absPath))
	}

	gitRemoteUrl, err := gitUrl.Parse(gitRemote)
	if err != nil {
		return "", errors.ScmUnspecifiedError(fmt.Sprintf("Unable to parse git remote url. %s", err))
	}

	_, err = git.PlainClone(absPath, false, &git.CloneOptions{
		URL:             gitRemoteUrl.String(),
		InsecureSkipTLS: false,
		//Auth:            auth, //unncessary, Basic auth is correctly extracted from the gitRemote url.
		Progress: os.Stdout,
	})

	return absPath, err
}
