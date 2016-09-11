package main

import "batatababa/cli"

/* Commands in this file:
gitty changelist create 5
gitty changelist delete 5
gitty changelist add 5 file.txt
gitty changelist remove 5 file.txt
gitty changelist revert 5
gitty changelist revert 5 filt.txt
gitty changelist show 5
gitty changelist 5
gitty changelist move 5 file.txt 6
gitty changelist commit 5
*/

// RepoAdd command
var Changelist = cli.Command{
	Name:        "changelist",
	Description: "Per changelist command set / Show a specific changelist",
	SubCommands: []cli.Command{
		cl_create,
		cl_delete,
		cl_add,
		cl_remove,
		cl_revert,
		cl_move,
		cl_commit,
		cl_show,
	},
}

var cl_create = cli.Command{
	Name:        "create",
	Description: "Create a changelist",
	Args:        []cli.Argument{arg_changelist, arg_dirOrFile},
}

var cl_delete = cli.Command{
	Name:        "delete",
	Description: "Delete a changelist",
	Args:        []cli.Argument{arg_changelist, arg_dirOrFile},
}

var cl_add = cli.Command{
	Name:        "add",
	Description: "Add to a changelist",
	Args:        []cli.Argument{arg_changelist, arg_dirOrFile},
}

var cl_remove = cli.Command{
	Name:        "remove",
	Description: "Remove files from a changelist",
	Args:        []cli.Argument{arg_changelist, arg_dirOrFile},
}

var cl_revert = cli.Command{
	Name:        "revert",
	Description: "Revert files in a changelist",
	Args:        []cli.Argument{arg_changelist, arg_dirOrFile},
}

var cl_move = cli.Command{
	Name:        "move",
	Description: "Move files from one changelist to another",
	Args: []cli.Argument{
		arg_srcChangelist,
		arg_destChangelist,
		arg_dirOrFile,
	},
}

var cl_commit = cli.Command{
	Name:        "commit",
	Description: "Commits a changelist to the repo",
	Args:        []cli.Argument{arg_changelist},
}

var cl_show = cli.Command{
	Name:        "show",
	Description: "Revert changes in a changelist",
	Args:        []cli.Argument{arg_changelist},
}
