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

	//https://github.com/go-git/go-git/blob/master/_examples/branch/main.go
	headRef, herr := repo.Head()
	if oerr != nil {
		return "", herr
	}

	// Create a new plumbing.HashReference object with the name of the branch
	// and the hash from the HEAD. The reference name should be a full reference
	// name and not an abbreviated one, as is used on the git cli.
	branchRef := plumbing.NewHashReference(plumbing.NewBranchReferenceName(localBranchName), headRef.Hash())

	// The created reference is saved in the storage.
	err := repo.Storer.SetReference(branchRef)
	if err != nil {
		return "", err
	}

	//this does not actually create a branch reference, just the repo config
	berr := repo.CreateBranch(&config.Branch{
		Name:  localBranchName,
		Merge: plumbing.NewBranchReferenceName(localBranchName),
	})
	if berr != nil {
		return "", berr
	}

	return localBranchName, nil
}
