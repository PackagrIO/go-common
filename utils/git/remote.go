package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"log"
)

func GitGetRemote(repoPath string, remoteName string) (string, error) {
	log.Printf("Getting remote (%s) for repo (%s)", remoteName, repoPath)
	repo, oerr := git.PlainOpen(repoPath)
	if oerr != nil {
		return "", oerr
	}

	remote, rerr := repo.Remote(remoteName)
	if rerr != nil {
		return "", rerr
	}

	pushUrl := remote.Config().URLs[0]
	if len(pushUrl) > 0 {
		return pushUrl, nil
	}

	return remote.String(), nil
}

func GitSetRemote(repoPath string, remoteName string, remoteUrl string) (string, error) {
	log.Printf("Setting repo (%s) remote (%s) to url (%s)", repoPath, remoteName, remoteUrl)
	repo, oerr := git.PlainOpen(repoPath)
	if oerr != nil {
		return "", oerr
	}

	remote, err := repo.CreateRemote(&config.RemoteConfig{
		Name: remoteName,
		URLs: []string{remoteUrl},
	})
	if err != nil {
		return "", nil
	}
	return remote.String(), nil
}
