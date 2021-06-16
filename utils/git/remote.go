package git

import (
	git2go "gopkg.in/libgit2/git2go.v25"
	"log"
)

func GitGetRemote(repoPath string, remoteName string) (string, error) {
	log.Printf("Getting remote (%s) for repo (%s)", remoteName, repoPath)
	repo, oerr := git2go.OpenRepository(repoPath)
	if oerr != nil {
		return "", oerr
	}

	remote, rerr := repo.Remotes.Lookup(remoteName)
	if rerr != nil {
		return "", rerr
	}

	log.Printf("Found remote URL: %s", remote.PushUrl())
	return remote.PushUrl(), nil
}
