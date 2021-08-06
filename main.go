package main

import (
	"io/ioutil"
	"os"

	git2go "github.com/libgit2/git2go/v31"
	log "github.com/sirupsen/logrus"
)

func main() {
	dir, err := ioutil.TempDir("", "git2go-")
	if err != nil {
		log.Panic(err)
	}
	// defer removeDir(dir)
	log.Info(dir)

	repo, err := git2go.Clone(
		"https://github.com/kyma-project/kyma",
		dir,
		&git2go.CloneOptions{
			CheckoutOpts: &git2go.CheckoutOptions{},
		},
	)
	if err != nil {
		log.Panic(err)
	}
	log.Info("cloned")

	// NewBranchIterator
	log.Info("branches")
	branchIter, err := repo.NewBranchIterator(git2go.BranchAll)
	if err != nil {
		log.Panic(err)
	}

	branchIter.ForEach(func(b *git2go.Branch, bt git2go.BranchType) error {
		log.Info(b.Reference.Name())
		return nil
	})

	// List all tags
	log.Info("tags")
	refIter, err := repo.Tags.List()
	if err != nil {
		log.Panic(err)
	}

	for _, name := range refIter {
		log.Info(name)
	}

	// resolve tag/branch/commit
	log.Info(repo.Head())
	// Oid from tag/branch name ( refs/remotes/origin/release-0.4 or 0.4.0)
	dwim, err := repo.References.Dwim("refs/remotes/origin/release-0.4")
	if err != nil {
		log.Panic(err)
	}

	// Oid from branch
	// branch, err := repo.LookupBranch("release-0.4", git2go.BranchAll)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// Oid from commit
	// oid, err := git2go.NewOid("4d1c45f0b82f4e0dfb0d484be430f0aaac47a214")
	// if err != nil {
	// 	log.Panic(err)
	// }

	commit, err := repo.LookupCommit(dwim.Target())
	if err != nil {
		log.Panic(err)
	}

	log.Info(commit.Id())
	log.Info(commit.ShortId())
	log.Infof("%+v", dwim.Target())

	err = repo.ResetToCommit(commit, git2go.ResetHard, &git2go.CheckoutOptions{})
	if err != nil {
		log.Panic(err)
	}

	// tree, err := repo.LookupTree(commit.TreeId())
	// if err != nil {
	// 	log.Panic(err)
	// }

	// err = repo.CheckoutTree(tree, &git2go.CheckoutOptions{
	// 	Strategy: git2go.CheckoutForce,
	// })
	// if err != nil {
	// 	log.Panic(err)
	// }

	// no need to do things like this in our case
	// err = repo.SetHeadDetached(dwim.Target())
	// if err != nil {
	// 	log.Panic(err)
	// }

	log.Info(repo.Head())
}

func removeDir(path string) {
	os.RemoveAll(path)
	log.Infof("Dir %s removed", path)
}
