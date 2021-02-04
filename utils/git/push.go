package git

import (
	"fmt"
	git2go "gopkg.in/libgit2/git2go.v25"
)

func GitPush(repoPath string, localBranch string, remoteBranch string, tagName string) error {
	//- https://gist.github.com/danielfbm/37b0ca88b745503557b2b3f16865d8c3
	//- https://stackoverflow.com/questions/37026399/git2go-after-createcommit-all-files-appear-like-being-added-for-deletion
	repo, oerr := git2go.OpenRepository(repoPath)
	if oerr != nil {
		return oerr
	}

	// Push
	remote, lerr := repo.Remotes.Lookup("origin")
	if lerr != nil {
		return lerr
	}
	//remote.ConnectPush(gitRemoteCallbacks(), &git.ProxyOptions{}, []string{})

	//err = remote.Push([]string{"refs/heads/master"}, nil, signature, message)
	return remote.Push([]string{
		fmt.Sprintf("refs/heads/%s:refs/heads/%s", localBranch, remoteBranch),
		fmt.Sprintf("refs/tags/%s:refs/tags/%s", tagName, tagName),
	}, new(git2go.PushOptions))
}
