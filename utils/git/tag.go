package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/packagrio/go-common/pipeline"
	"log"
)

func GitTag(repoPath string, version string, message string, signature *object.Signature) (string, error) {
	repo, oerr := git.PlainOpen(repoPath)
	if oerr != nil {
		return "", oerr
	}
	commitHead, herr := repo.Head()
	if herr != nil {
		return "", herr
	}

	commitObject, lerr := repo.CommitObject(commitHead.Hash())
	if lerr != nil {
		return "", lerr
	}
	commit := commitObject.Hash

	//tagId, terr := repo.Tags.CreateLightweight(version, commit, false)
	tagId, terr := repo.CreateTag(version, commit, &git.CreateTagOptions{
		Tagger:  signature,
		Message: fmt.Sprintf("(%s) %s", version, message),
	})
	if terr != nil {
		return "", terr
	}

	tagObj, terr := repo.TagObject(tagId.Hash())
	if terr != nil {
		return "", terr
	}
	return tagObj.Target.String(), terr
}

func GitGetTagDetails(repoPath string, tagName string) (*pipeline.GitTagDetails, error) {
	repo, oerr := git.PlainOpen(repoPath)
	if oerr != nil {
		return nil, oerr
	}

	tagRef, terr := repo.Tag(tagName)
	if terr != nil {
		return nil, terr
	}

	tag, lerr := repo.TagObject(tagRef.Hash()) //assume its an annotated tag.

	var currentTag *pipeline.GitTagDetails
	if lerr != nil {
		//this is a lightweight tag, not an annotated tag.
		commitRef, rerr := repo.CommitObject(tagRef.Hash())
		if rerr != nil {
			return nil, rerr
		}

		author := commitRef.Author

		log.Printf("Light-weight tag (%s) Commit ID: %s, DATE: %s", tagName, commitRef.Hash.String(), author.When.String())

		currentTag = &pipeline.GitTagDetails{
			TagShortName: tagName,
			CommitSha:    commitRef.Hash.String(),
			CommitDate:   author.When,
		}

	} else {

		log.Printf("Annotated tag (%s) Tag ID: %s, Commit ID: %s, DATE: %s", tagName, tag.Hash.String(), tag.Target.String(), tag.Tagger.When.String())

		currentTag = &pipeline.GitTagDetails{
			TagShortName: tagName,
			CommitSha:    tag.Target.String(),
			CommitDate:   tag.Tagger.When,
		}
	}
	return currentTag, nil
}

//TODO: enable this
// Get the nearest tag on branch.
// tag must be nearest, ie. sorted by their distance from the HEAD of the branch, not the date or tagname.
// basically `git describe --tags --abbrev=0`
//https://github.com/go-git/go-git/pull/584
//func GitFindNearestTagName(repoPath string) (string, error) {
//	repo, oerr := git2go.OpenRepository(repoPath)
//	if oerr != nil {
//		return "", oerr
//	}
//
//	//get the previous commit
//	ref, lerr := repo.References.Lookup("HEAD")
//	if lerr != nil {
//		return "", lerr
//	}
//	resRef, err := ref.Resolve()
//	if err != nil {
//		return "", err
//	}
//	headCommit, cerr := repo.LookupCommit(resRef.Target())
//	if cerr != nil {
//		return "", cerr
//	}
//
//	parentComit := headCommit.Parent(0)
//	defer parentComit.Free()
//
//	parentCommit, err := parentComit.AsCommit()
//	if err != nil {
//		return "", err
//	}
//
//	descOptions, derr := git2go.DefaultDescribeOptions()
//	if derr != nil {
//		return "", derr
//	}
//	descOptions.Strategy = git2go.DescribeTags
//	//descOptions.Pattern = "HEAD^"
//
//	formatOptions, ferr := git2go.DefaultDescribeFormatOptions()
//	if ferr != nil {
//		return "", ferr
//	}
//	formatOptions.AbbreviatedSize = 0
//
//	descr, derr := parentCommit.Describe(&descOptions)
//	if derr != nil {
//		return "", derr
//	}
//
//	nearestTag, ferr := descr.Format(&formatOptions)
//	if ferr != nil {
//		return "", ferr
//	}
//
//	return nearestTag, nil
//}
