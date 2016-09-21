package main

import (
	"testing"

	"github.com/batatababa/cli"
)

func TestAddRepo(t *testing.T) {
}

func TestRemoveRepo(t *testing.T) {
}

func TestCloneRepo(t *testing.T) {
}

func TestPrintTreeHelp(t *testing.T) {
	tree := cli.NewCommandTree()
	tree.Author = "Jeff Williams"
	tree.Root = Gitty
	tree.ToHelpString = toHelpString

	cli.PrintTreeHelp(&tree.Root)
}
