package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/batatababa/cli"
	"github.com/ryanuber/columnize"
)

func toStringArray(s string) []string {
	return strings.Split(s, " ")
}

func main() {
	globals = newConfig()
	initGitty(&globals)

	cli.Run(os.Args, &globals.comTree)
	// cli.Run(toStringArray("gitty branch"), &globals.comTree)
}

func initGitty(c *config) {
	if _, err := os.Stat(c.appDir); os.IsNotExist(err) {
		if os.Mkdir(c.appDir, 0755) != nil {
			fmt.Println("Unable to create Gitty directory")
		}
	}
	c.comTree = cli.NewCommandTree()
	c.comTree.Author = "Jeff Williams"
	c.comTree.ToHelpString = toHelpString

	c.comTree.Root = cli.Command{
		Name:        "gitty",
		Description: "Making git fun :)",
		SubCommands: []cli.Command{
			repo,
			branch,
			branchMap,
			changelist,
		},
	}
}

func toHelpString(c cli.Command, pathToRoot []string) (help string) {
	var helpBuf bytes.Buffer
	config := columnize.DefaultConfig()
	config.Prefix = "  "

	if len(pathToRoot) != 0 {
		helpBuf.WriteString(fmt.Sprintf(" %s", strings.Join(pathToRoot, " ")))
	}

	helpBuf.WriteString(fmt.Sprintf(" [%s]\n", c.Name))
	helpBuf.WriteString(fmt.Sprintf(" Description: %s\n", c.Description))

	if c.Usage == "" {
		if c.Args != nil {
			helpBuf.WriteString(fmt.Sprintf(" Usage:\n"))
			for _, a := range c.Args {
				if a.Name != "?" {
					helpBuf.WriteString(fmt.Sprintf("  %s <%s>\n", c.Name, a.Name))
				}
			}
			helpBuf.WriteString("\n")
		}
	} else {
		helpBuf.WriteString(fmt.Sprintf(" Usage:\n"))

		usageSl := strings.Split(c.Usage, "|")

		for _, usageSet := range usageSl {
			helpBuf.WriteString(fmt.Sprintf("  %s %s\n", c.Name, usageSet))
		}
		helpBuf.WriteString("\n")
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
