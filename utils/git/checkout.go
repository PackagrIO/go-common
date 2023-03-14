package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"log"
)

//checkout remote branch, create local branch if it doesn't exist
func GitCheckout(repoPath string, branchName string) error {
	repo, oerr := git.PlainOpen(repoPath)
	if oerr != nil {
		return oerr
	}

	//Getting the reference for the remote branch
	var remoteBranch *plumbing.Reference
	refIter, rerr := repo.References()
	if rerr != nil {
		return rerr
	}
	refIter.ForEach(func(r *plumbing.Reference) error {
		//log.Printf("ref(remote: %v): %s vs %s", r.Name().IsRemote(), r.Name().Short(), branchName)
		if r.Name().IsRemote() && r.Name().Short() == fmt.Sprintf("origin/%s", branchName) {
			remoteBranch = r
		}
		return nil
	})
	if remoteBranch == nil {
		log.Print("Failed to find remote branch: " + branchName)
		return fmt.Errorf("Failed to find remote branch: " + branchName)
	}

	workTree, err := repo.Worktree()
	if err != nil {
		return err
	}
	localBranchRef := plumbing.NewBranchReferenceName(branchName)
	err = workTree.Checkout(&git.CheckoutOptions{
		Hash:   remoteBranch.Hash(), //plumbing.NewHash(),
		Branch: localBranchRef,
		Force:  false,
		Keep:   false,
		Create: true,
	})

	if err != nil {
		log.Print("Failed to checkout branch: " + branchName)
		return err
	}
	return nil
}
