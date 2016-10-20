package main

import (
	"fmt"

	"github.com/batatababa/cli"
)

/* Command Sets*/
var changelist = cli.Command{
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
	Usage:  fmt.Sprintf("<%s>|<%s>|<%s>", arg_changelist.Name, arg_dirOrFile.Name, arg_cl_rtnShowAll.Name),
	Args:   []cli.Argument{arg_changelist, arg_dirOrFile, arg_cl_rtnShowAll},
	Action: action_clShow,
}

/* Commands */
var cl_create = cli.Command{
	Name:        "create",
	Description: "Create a changelist",
	Args:        []cli.Argument{arg_cl_createNamed, arg_cl_rtnCreate},
	Action:      action_clCreate,
}

var cl_delete = cli.Command{
	Name:        "delete",
	Description: "Delete a changelist",
	Args:        []cli.Argument{arg_changelist},
	Action:      action_clDelete,
}

var cl_add = cli.Command{
	Name:        "add",
	Description: "Add files to a changelist",
	Usage:       fmt.Sprintf("<%s> <%s>", arg_changelist.Name, arg_dirOrFile.Name),
	Args:        []cli.Argument{arg_changelist, arg_dirOrFile},
	Action:      action_clAdd,
}

var cl_remove = cli.Command{
	Name:        "remove",
	Description: "Remove files from a changelist",
	Usage:       fmt.Sprintf("<%s> <%s>", arg_changelist.Name, arg_dirOrFile.Name),
	Args:        []cli.Argument{arg_changelist, arg_dirOrFile},
	Action:      action_clRemove,
}

var cl_revert = cli.Command{
	Name:        "revert",
	Description: "Revert a changelist or files in a changelist",
	Usage:       fmt.Sprintf("<%s> |<%s> <%s>", arg_changelist.Name, arg_changelist.Name, arg_dirOrFile.Name),
	Args:        []cli.Argument{arg_changelist, arg_dirOrFile},
	Action:      action_clRevert,
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
	Action: action_clMove,
}

var cl_commit = cli.Command{
	Name:        "commit",
	Description: "Commit a changelist to the repo",
	Args:        []cli.Argument{arg_changelist},
	Action:      action_clCommit,
}

var cl_show = cli.Command{
	Name:        "show",
	Description: "Show files in a changelist(s)",
	Args:        []cli.Argument{arg_changelist},
	Action:      action_clShow,
}

/* Arguments */
var arg_cl_rtnShowAll = cli.Argument{
	Name:        "rtn",
	Description: "Carriage Return - Show all changelists",
}

var arg_cl_rtnCreate = cli.Argument{
	Name:        "rtn",
	Description: "Carriage Return - Create a changelist with an auto-generated name",
}

var arg_cl_createNamed = cli.Argument{
	Name:        "name",
	Description: "Create a changelist with the specified name",
}
