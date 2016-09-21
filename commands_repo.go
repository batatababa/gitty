package main

import "github.com/batatababa/cli"

// Clone command
var Clone = cli.Command{
	Name:        "clone",
	Description: "Copies a Git repo from a server. Does this give a default workspace?",
	Args:        []cli.Argument{arg_urlOrPath},
}

// Repo base command
var Repo = cli.Command{
	Name:        "repo",
	Description: "Add and remove repos from Gitty control",
	SubCommands: []cli.Command{RepoAdd, RepoRemove},
}

// RepoAdd command
var RepoAdd = cli.Command{
	Name:        "add",
	Description: "Registers a repo with Gitty",
	Args:        []cli.Argument{arg_urlOrPath},
}

// RepoRemove command
var RepoRemove = cli.Command{
	Name:        "remove",
	Description: "Un-registers a repo with Gitty",
	Args:        []cli.Argument{arg_repoNum},
}
