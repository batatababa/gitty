package main

import "batatababa/cli"

var arg_urlOrPath = cli.Argument{
	Name:        "url:path",
	Description: "Url or Path of an existing Git Repo",
}

var arg_changelist = cli.Argument{
	Name:        "changelist",
	Description: "Locally unique identifier of a changelist",
}

var arg_dirOrFile = cli.Argument{
	Name:        "dir:file",
	Description: "Directory or FileName",
}

var arg_branch = cli.Argument{
	Name:        "branch",
	Description: "Name of a branch",
}

var arg_targetBranch = cli.Argument{
	Name:        "target_branch",
	Description: "Name of a branch",
}

var arg_sourceBranch = cli.Argument{
	Name:        "source_branch",
	Description: "Name of a branch",
}

var arg_srcChangelist = cli.Argument{
	Name:        "src_changelist",
	Description: "Name of a changelist",
}

var arg_destChangelist = cli.Argument{
	Name:        "dest_changelist",
	Description: "Name of a changelist",
}
