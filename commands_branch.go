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
// var ChangelistArg = cli.Argument{
//     Name:        "changelist",
//     Description: "Locally unique identifier of a changelist",
// }

// var DirFileArg = cli.Argument{
//     Name:        "dir|file",
//     Description: "Directory or FileName",
// }

// RepoAdd command
var Map = cli.Command{
	Name:        "Map",
	Description: "Map Repo Branches To Directories",
	SubCommands: []cli.Command{
		map_remove,
	},
	Args: []cli.Argument{
		{
			Name:        "cr",
			Description: "Show branch mappings",
		},
	},
}

var map_remove = cli.Command{
	Name:        "remove",
	Description: "Delete a changelist",
	Args:        []cli.Argument{arg_changelist, arg_dirOrFile},
}

var Branch = cli.Command{
	Name:        "branch",
	Description: "Per changelist command set / Show a specific changelist",
	SubCommands: []cli.Command{
		cl_create,
		cl_delete,
		cl_add,
		cl_remove,
		cl_revert,
		cl_move,
		cl_show,
		cl_commit,
	},
}

var Pull = cli.Command{
	Name:        "pull",
	Description: "Create a changelist",
	Args:        []cli.Argument{arg_changelist, arg_dirOrFile},
}
