package main

import "github.com/batatababa/cli"

//#branch commands
//gitty map
//- Shows all of the "Branch Mappings" that Gitty knows about
//- Branches are shown with a reference # that can be used anywhere a branch
//is expected
//gitty map <branch> <path>
//- Maps a branch to a location on the file system
//gitty map remove <branch>
//- Removes a "Branch Mapping"
//gitty branch <source branch> <target branch>
//- Creates a branch
//gitty pull
//- Updates the repo and attempts to merge changes into all branches
//gitty pull <branch>
//- Updates the repo and attempts to merge changes into target branch

var map_arg_cr_show = cli.Argument{
	Name:        "cr",
	Description: "Show branch mappings",
}

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
