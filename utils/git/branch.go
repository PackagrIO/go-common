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

	//// Lookup head commit
	//commitHead, herr := repo.Head()
	//if herr != nil {
	//	return "", herr
	//}
	//
	//commit, lerr := repo.LookupCommit(commitHead.Target())
	//if lerr != nil {
	//	return "", lerr
	//}
	//newLocalBranch, err := repo.CreateBranch(localBranchName, commit, false)

	//commitHead, herr := repo.Head()
	//if herr != nil {
	//	return "", herr
	//}

	err := repo.CreateBranch(&config.Branch{
		Name:  localBranchName,
		Merge: plumbing.NewBranchReferenceName(localBranchName),
	})
	if err != nil {
		return "", err
	}
	return localBranchName, nil

	//
	//workTree, err := repo.Worktree()
	//if err != nil {
	//	return "", err
	//}
	//localBranchRef := plumbing.NewBranchReferenceName(localBranchName)
	//err = workTree.Checkout(&git.CheckoutOptions{
	//	Hash:   commitHead.Hash(), //plumbing.NewHash(),
	//	Branch: localBranchRef,
	//	Force:  true,
	//	Keep:   false,
	//	Create: true,
	//})
	//
	//if err != nil {
	//	log.Print("Failed to create local branch: " + localBranchName)
	//	return "", err
	//}
	//return localBranchRef.Short(), nil
}
