package main

import (
	"fmt"
	"strings"

	"batatababa/cli"
)

// Gitty command
var Gitty = &cli.Command{
	Name:        "gitty",
	Description: "Making git fun :)",
	SubCommands: []cli.Command{*Clone, *Repo},
}

func main() {
	//runCli(os.Args)
	runCli(toStringArray("gitty -h"))
}

func runCli(s []string) {
	tree := cli.NewCommandTree()
	tree.Author = "Jeff Williams"
	tree.Root = *Gitty

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
