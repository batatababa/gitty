package main

import (
	"fmt"

	"github.com/batatababa/cli"
)

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
	Usage: fmt.Sprintf("<%s>|<%s>|<%s>", arg_changelist.Name, arg_dirOrFile.Name, cl_arg_cr_show_all.Name),
	Args:  []cli.Argument{arg_changelist, arg_dirOrFile, cl_arg_cr_show_all},
}

var cl_create = cli.Command{
	Name:        "create",
	Description: "Create a changelist",
	Args:        []cli.Argument{cl_arg_create_named, cl_arg_cr_create},
}

var cl_delete = cli.Command{
	Name:        "delete",
	Description: "Delete a changelist",
	Args:        []cli.Argument{arg_changelist},
}

var cl_add = cli.Command{
	Name:        "add",
	Description: "Add files to a changelist",
	Usage:       fmt.Sprintf("<%s> <%s>", arg_changelist.Name, arg_dirOrFile.Name),
	Args:        []cli.Argument{arg_changelist, arg_dirOrFile},
}

var cl_remove = cli.Command{
	Name:        "remove",
	Description: "Remove files from a changelist",
	Usage:       fmt.Sprintf("<%s> <%s>", arg_changelist.Name, arg_dirOrFile.Name),
	Args:        []cli.Argument{arg_changelist, arg_dirOrFile},
}

var cl_revert = cli.Command{
	Name:        "revert",
	Description: "Revert a changelist or files in a changelist",
	Usage:       fmt.Sprintf("<%s> |<%s> <%s>", arg_changelist.Name, arg_changelist.Name, arg_dirOrFile.Name),
	Args:        []cli.Argument{arg_changelist, arg_dirOrFile},
}

var cl_move = cli.Command{
	Name:        "move",
	Description: "Move files from one changelist to another",
	Usage:       fmt.Sprintf("<%s> <%s> <%s>", arg_srcChangelist.Name, arg_destChangelist.Name, arg_dirOrFile.Name),
	Args: []cli.Argument{
		arg_srcChangelist,
		arg_destChangelist,
		arg_dirOrFile,
	},
}

var cl_commit = cli.Command{
	Name:        "commit",
	Description: "Commit a changelist to the repo",
	Args:        []cli.Argument{arg_changelist},
}

var cl_show = cli.Command{
	Name:        "show",
	Description: "Show files in a changelist(s)",
	Args:        []cli.Argument{arg_changelist},
}

var cl_arg_cr_show_all = cli.Argument{
	Name:        "rtn",
	Description: "(Carriage Return) Show all changelists",
}
var cl_arg_cr_create = cli.Argument{
	Name:        "rtn",
	Description: "(Carriage Return) Create a changelist with an auto-generated name",
}

var cl_arg_create_named = cli.Argument{
	Name:        "name",
	Description: "Create a changelist with the specified name",
}
