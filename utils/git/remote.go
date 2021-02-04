package git

import git2go "gopkg.in/libgit2/git2go.v25"

func GitGetRemote(repoPath string, remoteName string) (string, error) {
	repo, oerr := git2go.OpenRepository(repoPath)
	if oerr != nil {
		return "", oerr
	}

	remote, rerr := repo.Remotes.Lookup(remoteName)
	if rerr != nil {
		return "", rerr
	}

	return remote.PushUrl(), nil
}
