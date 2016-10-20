package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/batatababa/cli"
)

func action_clone(c cli.Command) (err error) {
	return err
}

func action_repoAdd(c cli.Command) (err error) {
	err = verifyArgCountEqual(c.Args, 1)
	if err != nil {
		return err
	}

	appDir, err := getAppDir()

	if err != nil {
		log.Fatal("Could not get app directory")

		return err
	}

	os.MkdirAll(appDir, 0660)
	return err
}

func action_repoRemove(c cli.Command) (err error) {
	fmt.Println("todo: repoRemove")

	return err
}

func action_repoActive(c cli.Command) (err error) {
	fmt.Println("todo: repoActive")
	return err
}

func action_repoShow(c cli.Command) (err error) {

	fmt.Println("todo: repoShow")
	return err
}

func getAppDir() (dir string, err error) {
	usr, err := user.Current()
	dir = usr.HomeDir + "/.gitty"
	return dir, err
}
