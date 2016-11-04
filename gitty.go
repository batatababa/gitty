package main

import (
	"os"
	"os/user"

	"github.com/batatababa/cli"
	"github.com/batatababa/futil"
)

type config struct {
	appFolder  string
	appDir     string
	comTree    cli.CommandTree
	repos      string
	activeRepo string
}

var globals config

func newConfig() (c config) {
	c.appFolder = ".gitty"

	usr, err := user.Current()
	if err != nil {
		os.Exit(1)
	}

	c.appDir = usr.HomeDir + "/" + c.appFolder
	c.repos = c.appDir + "/repos"
	c.activeRepo = c.appDir + "/active_repo"
	return c
}

func getActiveRepo() (ar string, err error) {
	v, err := futil.NewFVector(globals.activeRepo)
	if err != nil {
		return ar, err
	}

	return v.At(0)
}
