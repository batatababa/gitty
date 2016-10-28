package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/batatababa/cli"
	"github.com/batatababa/futil"
	"github.com/libgit2/git2go"
)

func trimPath(p string) string {
	p = strings.TrimSpace(p)
	p = strings.TrimRight(p, "/")
	return p
}

func action_clone(c cli.Command) (err error) {
	return err
}

func action_repoAdd(c cli.Command) (err error) {
	err = verifyArgCountEqual(c.Args, 1)
	if err != nil {
		return err
	}

	repoDir := trimPath(c.Args[0].Value)
	fmt.Println(repoDir)

	if fileInfo, err := os.Stat(repoDir); os.IsNotExist(err) || fileInfo.IsDir() == false {
		return errGitty(repoDir + " is not a directory")
	}

	if _, err := git.OpenRepository(repoDir); err != nil {
		return errGitty(repoDir + " is not a Git repo")
	}

	v, err := futil.NewFVector(globals.repos)
	if err != nil {
		return errGitty("Could not open " + globals.repos)
	}

	found, err := v.Contains(repoDir)
	if err != nil {
		return errGitty("Could not read " + globals.repos)
	}

	if found == true {
		return errGitty(repoDir + " is already mangaged by Gitty")
	} else {
		if err = v.Add(repoDir); err != nil {
			return errGitty("Could not write to " + globals.repos)
		}
	}

	return err
}

func action_repoRemove(c cli.Command) (err error) {
	err = verifyArgCountEqual(c.Args, 1)
	if err != nil {
		return err
	}

	v, err := futil.NewFVector(globals.repos)
	if err != nil {
		return errGitty("Could not open " + globals.repos)
	}

	if repoNum, err := strconv.Atoi(c.Args[0].Value); err == nil {
		if _, err = v.RemoveAt(repoNum); err != nil {
			return errGitty("Could not remove repo #" + c.Args[0].Value)
		}
	} else {
		repoDir := trimPath(c.Args[0].Value)
		if err = v.Remove(repoDir); err != nil {
			return errGitty(err.Error()) //"Could not write to " + globals.repos)
		}
	}

	return err
}

func action_repoActive(c cli.Command) (err error) {
	fmt.Println("todo: repoActive")
	return err
}

func action_repoShow(c cli.Command) (err error) {
	err = verifyArgCountEqual(c.Args, 0)
	if err != nil {
		return err
	}

	v, err := futil.NewFVector(globals.repos)
	if err != nil {
		return errGitty("Could not open " + globals.repos)
	}

	sl, err := v.Slice()

	fmt.Println("  Registered Repos")

	for i, repo := range sl {
		fmt.Printf("    %d: %s\n", i, repo)
	}

	return err
}
