package main

import (
	"fmt"
	"strings"

	"github.com/batatababa/cli"
	"github.com/libgit2/git2go"
)

func action_mapShow(c cli.Command) (err error) {
	fmt.Println("todo: map show")
	return err
}

func action_mapRemove(c cli.Command) (err error) {
	fmt.Println("todo: map remove")
	return err
}

func action_branchShow(c cli.Command) (err error) {
	var remotes []string

	f := func(b *git.Branch, t git.BranchType) error {
		name, err := b.Name()
		if err != nil {
			return err
		}

		if t == git.BranchRemote {
			remotes = append(remotes, name)
		}
		return nil
	}

	repoDir, err := getActiveRepo()
	if err != nil {
		return err
	}

	repo, err := git.OpenRepository(repoDir)
	if err != nil {
		fmt.Println(err)
		return err
	}

	iter, err := repo.NewBranchIterator(git.BranchAll)
	if err != nil {
		return err
	}

	err = iter.ForEach(f)
	if err != nil && !git.IsErrorCode(err, git.ErrIterOver) {
		return err
	}

	fmt.Printf(" Branches: \n")
	for _, r := range remotes {
		if !strings.Contains(r, "HEAD") {
			fmt.Printf("  %s\n", r)
		}
	}
	return err
}

func action_branchRename(c cli.Command) (err error) {
	fmt.Println("todo: branch rename")
	return err
}

func action_branchRemove(c cli.Command) (err error) {
	fmt.Println("todo: branch remove")
	return err
}
