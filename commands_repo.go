package main

import "batatababa/cli"

// Clone command
var Clone = &cli.Command{
	Name:        "clone",
	Description: "Copies a Git repo from a server. Does this give a default workspace?",
}

// Repo base command
var Repo = &cli.Command{
	Name:        "repo",
	Description: "Add and remove repos from Gitty control.",
	SubCommands: []cli.Command{*RepoAdd, *RepoRemove},
}

// RepoAdd command
var RepoAdd = &cli.Command{
	Name:        "add",
	Description: "Registers a repo with Gitty, so it can be maintained.",
}

// RepoRemove command
var RepoRemove = &cli.Command{
	Name:        "remove",
	Description: "Un-registers a repo with Gitty",
}
