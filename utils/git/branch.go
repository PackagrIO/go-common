package git

import (
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
)

func GitGetBranch(repoPath string) (string, error) {

	repo, oerr := git.PlainOpen(repoPath)
	if oerr != nil {
		return "", oerr
	}
	head, herr := repo.Head()
	if herr != nil {
		return "", herr
	}
	return head.Name().Short(), nil
}

//create  branch (does not checkout)
func GitCreateBranchFromHead(repoPath string, localBranchName string) (string, error) {
	repo, oerr := git.PlainOpen(repoPath)
	if oerr != nil {
		return "", oerr
	}

	err := repo.CreateBranch(&config.Branch{
		Name:  localBranchName,
		Merge: plumbing.NewBranchReferenceName(localBranchName),
	})
	if err != nil {
		return "", err
	}
	return localBranchName, nil
}
