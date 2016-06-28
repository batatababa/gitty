package main

import "batatababa/cli"

var ChangelistArg = &cli.Argument{
	Name:        "Changelist Number",
	Description: "Locally unique identifier of a changelist",
}

var DirFileArg = &cli.Argument{
	Name:        "Dir | File",
	Description: "Directory or FileName",
}

// Clone command
var Commit = &cli.Command{
	Name:        "commit",
	Description: "Commits all changelists to the repo",
	Args:        []cli.Argument{*ChangelistArg},
}

// Repo base command
var Changelist = &cli.Command{
	Name:        "changelist",
	Description: "Show all changelists",
	SubCommands: []cli.Command{*Changelist_add, *Changelist_cl},
}

// RepoAdd command
var Changelist_add = &cli.Command{
	Name:        "add",
	Description: "Create a New changelist",
	Args: []cli.Argument{{
		Name:        "name",
		Description: "The name of new changelist",
	}},
}

// RepoAdd command
var Changelist_cl = &cli.Command{
	Name:        "Changelist Number",
	Description: "Per changelist command set / Show a specific changelist",
	SubCommands: []cli.Command{
		*Changelist_cl_add,
		*Changelist_cl_revert,
		*Changelist_cl_move,
		*Changelist_cl_commit,
	},
	Wildcard: true,
}

var Changelist_cl_add = &cli.Command{
	Name:        "add",
	Description: "Add a file or all the files in a directory to a changelist",
	Args:        []cli.Argument{*DirFileArg},
}

var Changelist_cl_revert = &cli.Command{
	Name:        "revert",
	Description: "Revert all changes in a changelist",
	Args:        []cli.Argument{*DirFileArg},
}
var Changelist_cl_move = &cli.Command{
	Name:        "move",
	Description: "Move files from one changelist to another",
	Args: []cli.Argument{
		*DirFileArg,
		{
			Name:        "destination changelist",
			Description: "The changelist to move the files to.",
		},
	},
}
var Changelist_cl_commit = &cli.Command{
	Name:        "commit",
	Description: "Commits a changelist to the repo",
}
