package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/ryanuber/columnize"

	"github.com/batatababa/cli"
)

// Gitty command
var Gitty = cli.Command{
	Name:        "gitty",
	Description: "Making git fun :)",
	SubCommands: []cli.Command{
		Clone,
		Repo,
		Changelist,
	},
}

func main() {
	runCli(os.Args)
	// runCli(toStringArray("gitty repo -h"))
}

func runCli(s []string) {
	tree := cli.NewCommandTree()
	tree.Author = "Jeff Williams"
	tree.Root = Gitty
	tree.ToHelpString = toHelpString

	// fmt.Println(os.Args)
	// err := cli.Run(strings.Split("gitty repo add -h", " "), &tree)

	err := cli.Run(s, &tree)

	if err != nil {
		fmt.Println(err)
	}
}

func toStringArray(s string) []string {
	return strings.Split(s, " ")
}

func toHelpString(c cli.Command, pathToRoot []string) (help string) {
	var helpBuf bytes.Buffer
	config := columnize.DefaultConfig()
	config.Prefix = "  "

	helpBuf.WriteString(fmt.Sprintf("Description: %s\n", c.Description))

	if c.Usage == "" {
		helpBuf.WriteString(fmt.Sprintf("Usage: %s ", c.Name))

		if c.Args != nil {
			for _, a := range c.Args {
				if a.Name != "?" {
					helpBuf.WriteString(fmt.Sprintf("<%s> ", a.Name))
				}
			}
			helpBuf.WriteString("\n\n")
		}
	} else {
		helpBuf.WriteString(fmt.Sprintf("Usage: %s\n\n", c.Usage))
	}

	if c.Args != nil {
		helpBuf.WriteString(" Arguments:\n")
		var args []string
		for _, a := range c.Args {
			args = append(args, fmt.Sprintf("<%s>|%s", a.Name, a.Description))
		}
		helpBuf.WriteString(columnize.Format(args, config))
		helpBuf.WriteString("\n\n")
	}

	if c.Flags != nil {
		helpBuf.WriteString(" Flags:\n")
		var flags []string
		for _, f := range c.Flags {
			flags = append(flags, toShortLongDescString(f.ShortName, f.LongName, f.Description))
		}
		helpBuf.WriteString(columnize.Format(flags, config))
		helpBuf.WriteString("\n\n")
	}

	if c.Opts != nil {
		helpBuf.WriteString(" Options:\n")
		var opts []string
		for _, o := range c.Opts {
			opts = append(opts, toShortLongDescString(o.ShortName, o.LongName, o.Description))
		}
		helpBuf.WriteString(columnize.Format(opts, config))
		helpBuf.WriteString("\n\n")
	}

	if c.SubCommands != nil {
		helpBuf.WriteString(" SubCommands:\n")
		var subs []string
		for _, s := range c.SubCommands {
			subs = append(subs, fmt.Sprintf("%s:|%s", s.Name, s.Description))
		}
		helpBuf.WriteString(columnize.Format(subs, config))
		helpBuf.WriteString("\n")
	}

	help = helpBuf.String()
	return
}

func toShortLongDescString(short string, long string, description string) (str string) {
	var buf bytes.Buffer
	if short != "" {
		buf.WriteString(fmt.Sprintf("-%s", short))
	}
	buf.WriteString("|")
	if long != "" {
		buf.WriteString(fmt.Sprintf("--%s", long))
	}
	buf.WriteString(fmt.Sprintf(",|%s", description))
	str = buf.String()
	return
}
