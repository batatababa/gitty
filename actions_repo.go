package main

import (
	"fmt"
	"os"
	"path/filepath"
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

func action_repoClone(c cli.Command) (err error) {
	err = verifyArgCountEqual(c.Args, 1)
	if err != nil {
		return err
	}
	url := c.Args[0].Value

	if strings.HasPrefix(url, "http") || strings.HasPrefix(url, "ftp") {
		url = strings.TrimSuffix(url, ".git")
	} else {
		url = strings.TrimRight(url, "/")
	}

	_, filename := filepath.Split(url)

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	path := pwd + "/" + filename
	fmt.Println(path)

	opts := git.CloneOptions{
		Bare: false,
	}
	_, err = git.Clone(url, path, &opts)

	return err
}

func action_repoInit(c cli.Command) (err error) {
	var path = ""
	err = verifyArgCountLessThan(c.Args, 2)
	if err != nil {
		return err
	}

	if len(c.Args) > 0 {
		path = strings.TrimRight(c.Args[0].Value, "/")
	} else {
		path, err = os.Getwd()
	}
	_, err = git.InitRepository(path, false)

	return err
}

func action_repoAdd(c cli.Command) (err error) {
	err = verifyArgCountEqual(c.Args, 1)
	if err != nil {
		return err
	}

	repoDir := trimPath(c.Args[0].Value)

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
			return errGitty(err.Error())
		}
	}

	return err
}

func action_repoActive(c cli.Command) (err error) {
	repoDir := ""
	err = verifyArgCountLessThan(c.Args, 2)
	if err != nil {
		return err
	}

	if len(c.Args) == 0 {
		if err = action_repoActiveShow(); err != nil {
			return errGitty(err.Error())
		}
	} else {
		vRepos, err := futil.NewFVector(globals.repos)
		if err != nil {
			return errGitty("Could not open " + globals.repos)
		}

		if repoNum, err := strconv.Atoi(c.Args[0].Value); err == nil {
			if repoDir, err = vRepos.At(repoNum); err != nil {
				return errGitty("Could not remove repo #" + c.Args[0].Value)
			}
		} else {
			repoDir = trimPath(c.Args[0].Value)
			found, err := vRepos.Contains(repoDir)
			if err != nil {
				return errGitty(err.Error())
			}
			if found == false {
				return errGitty("Could not find repo, " + repoDir)
			}
		}

		vActiveRepo, err := futil.NewFVector(globals.activeRepo)
		if err != nil {
			return errGitty("Could not open " + globals.activeRepo)
		}

		err = vActiveRepo.Clear()
		if err != nil {
			return errGitty("Could not set the active repo")
		}

		if err = vActiveRepo.Add(repoDir); err != nil {
			return errGitty(err.Error())
		}
	}

	return err
}

func action_repoActiveShow() (err error) {
	v, err := futil.NewFVector(globals.activeRepo)
	if err != nil {
		return errGitty("Could not open " + globals.activeRepo)
	}

	if v.Size() == 0 {
		fmt.Println("No active repo")
	} else {
		active, err := v.At(0)
		if err != nil {

		} else {
			fmt.Println("Active Repo: " + active)
		}
	}

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
