package main

import "github.com/batatababa/cli"

var arg_urlOrPath = cli.Argument{
	Name:        "url:path",
	Description: "Url or Path of an existing Git repo",
}

var arg_changelist = cli.Argument{
	Name:        "cl-num",
	Description: "Locally unique identifier of a changelist",
}

var arg_dirOrFile = cli.Argument{
	Name:        "dir:file",
	Description: "Directory or file name",
}

var arg_repoNum = cli.Argument{
	Name:        "r-num",
	Description: "Unique identifier for a repo",
}

var arg_dirOrFile_optional = cli.Argument{
	Name:        "dir:file",
	Description: "(Optional) Directory or file name",
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
	Name:        "src_cl",
	Description: "Source changelist",
}

var arg_destChangelist = cli.Argument{
	Name:        "dst_cl",
	Description: "Destination changelist",
}
