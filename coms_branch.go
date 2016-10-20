package main

import (
	"fmt"

	"github.com/batatababa/cli"
)

/* Command Sets */
var branchMap = cli.Command{
	Name:        "map",
	Description: "Map branches from the active repo to system directories / Show branch mappings",
	SubCommands: []cli.Command{map_remove},
	Usage:       fmt.Sprintf("<%s> <%s>|<%s>", arg_branch.Name, arg_dir.Name, arg_map_rtnShowAll.Name),
	Args:        []cli.Argument{arg_branch, arg_dir, arg_map_rtnShowAll},
	Action:      action_mapShow,
}

var branch = cli.Command{
	Name:        "branch",
	Description: "Create a branch / Show all branches",
	SubCommands: []cli.Command{
		br_remove,
		br_rename,
	},
	Usage:  fmt.Sprintf("<%s> <%s>|<%s>", arg_br_sourceBranch.Name, arg_br_newBranch.Name, arg_br_rtnShowAll.Name),
	Args:   []cli.Argument{arg_br_sourceBranch, arg_br_newBranch, arg_br_rtnShowAll},
	Action: action_branchShow,
}

/* Commands */
var map_remove = cli.Command{
	Name:        "remove",
	Description: "Remove a branch mapping (Does not remove files)",
	Args:        []cli.Argument{arg_map},
	Action:      action_mapRemove,
}

var br_rename = cli.Command{
	Name:        "rename",
	Description: "Rename a branch",
	Usage:       fmt.Sprintf("<%s> <%s>", arg_branch.Name, arg_br_newName.Name),
	Args:        []cli.Argument{arg_branch, arg_br_newName},
	Action:      action_branchRename,
}

var br_remove = cli.Command{
	Name:        "remove",
	Description: "Remove a branch",
	Args:        []cli.Argument{arg_branch},
	Action:      action_branchRemove,
}

/* Arguments */
var arg_map_rtnShowAll = cli.Argument{
	Name:        "rtn",
	Description: "Carriage Return - Show all branch mappings",
}

var arg_br_newBranch = cli.Argument{
	Name:        "br-name",
	Description: "Name for the new branch",
}

var arg_br_newName = cli.Argument{
	Name:        "br-name",
	Description: "New name for the branch",
}

var arg_br_sourceBranch = cli.Argument{
	Name:        "source_branch",
	Description: "Name or ID of a branch",
}

var arg_br_rtnShowAll = cli.Argument{
	Name:        "rtn",
	Description: "Carriage Return - Show all branches",
}
