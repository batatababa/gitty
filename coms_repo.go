package main

import "github.com/batatababa/cli"

/* Command Sets */
var repo = cli.Command{
	Name:        "repo",
	Description: "Add and remove repos from Gitty control",
	SubCommands: []cli.Command{repo_init, repo_clone, repo_add, repo_remove, repo_active},
	Args:        []cli.Argument{arg_repo_rtnShowAll},
	Action:      action_repoShow,
}

/* Commands */
var repo_init = cli.Command{
	Name:        "init",
	Description: "Create a Git repo from source",
	Args:        []cli.Argument{arg_srcPath},
	Action:      action_repoInit,
}

var repo_clone = cli.Command{
	Name:        "clone",
	Description: "Copies a Git repo from a server. Does this give a default workspace?",
	Args:        []cli.Argument{arg_urlOrPath},
	Action:      action_repoClone,
}

var repo_add = cli.Command{
	Name:        "add",
	Description: "Registers a repo with Gitty",
	Args:        []cli.Argument{arg_urlOrPath},
	Action:      action_repoAdd,
}

var repo_remove = cli.Command{
	Name:        "remove",
	Description: "Un-registers a repo with Gitty",
	Args:        []cli.Argument{arg_repoNum},
	Action:      action_repoRemove,
}

var repo_active = cli.Command{
	Name:        "active",
	Description: "Set which Repo is actively managed by gitty",
	Args:        []cli.Argument{arg_repoNum, arg_repo_rtnShowActive},
	Action:      action_repoActive,
}

/* Arguments */
var arg_repo_rtnShowAll = cli.Argument{
	Name:        "rtn",
	Description: "(Carriage Return) Show all repos that gitty can see",
}

var arg_repo_rtnShowActive = cli.Argument{
	Name:        "rtn",
	Description: "(Carriage Return) Show the active repo",
}
