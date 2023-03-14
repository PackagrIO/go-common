package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"strings"
)

// push local branch and tag to remote repository
// remoteUrl should include embededd username and password, but can also use ssh credentials.
func GitPush(repoPath string, localBranch string, remoteUrl string, remoteBranch string, tagName string) error {
	//- https://gist.github.com/danielfbm/37b0ca88b745503557b2b3f16865d8c3
	//- https://stackoverflow.com/questions/37026399/git2go-after-createcommit-all-files-appear-like-being-added-for-deletion
	repo, oerr := git.PlainOpen(repoPath)
	if oerr != nil {
		return oerr
	}

	// Push
	remote, rerr := repo.CreateRemoteAnonymous(&config.RemoteConfig{
		Name: "anonymous",
		URLs: []string{remoteUrl},
	})
	if rerr != nil {
		return rerr
	}

	//strip the fully qualified branch ref if present.
	localBranch = strings.TrimPrefix(localBranch, "refs/heads/")
	remoteBranch = strings.TrimPrefix(remoteBranch, "refs/heads/")
	pushSpecs := []config.RefSpec{
		config.RefSpec(fmt.Sprintf("refs/heads/%s:refs/heads/%s", localBranch, remoteBranch)),
		config.RefSpec(fmt.Sprintf("refs/tags/%s:refs/tags/%s", tagName, tagName)),
	}
	return remote.Push(&git.PushOptions{
		RefSpecs:        pushSpecs,
		InsecureSkipTLS: true,
	})
}
