package main

import (
	"fmt"
	"io"
	"os"

	"github.com/batatababa/cli"
	"github.com/libgit2/git2go"
)

func action_clone(c cli.Command) (err error) {
	return err
}

func action_repoAdd(c cli.Command) (err error) {
	err = verifyArgCountEqual(c.Args, 1)
	if err != nil {
		return err
	}

	repoDir := c.Args[0].Value

	if fileInfo, err := os.Stat(repoDir); os.IsNotExist(err) || fileInfo.IsDir() == false {
		return errGitty(repoDir + " is not a directory")
	}

	if _, err := git.OpenRepository(repoDir); err != nil {
		return errGitty(repoDir + " is not a Git repo")
	}

	f, err := fileGet(globals.repos, os.O_APPEND|os.O_RDWR)
	if err != nil {
		return errGitty("Could not open " + globals.repos)
	}

	_, err = fileContains(f, repoDir)
	if err == io.EOF {
		if _, err = f.WriteString(repoDir + "\n"); err != nil {
			return errGitty("Could not write to " + globals.repos)
		}
	} else if err == nil {
		return errGitty(repoDir + " is already mangaged by Gitty")
	} else {
		return errGitty("Could not read " + globals.repos)
	}

	return err
}

func action_repoRemove(c cli.Command) (err error) {
	fmt.Println("todo: repoRemove")

	return err
}

func action_repoActive(c cli.Command) (err error) {
	fmt.Println("todo: repoActive")
	return err
}

func action_repoShow(c cli.Command) (err error) {

	fmt.Println("todo: repoShow")
	return err
}
