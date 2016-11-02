package main

import (
	"errors"
	"fmt"

	"github.com/batatababa/cli"
)

func errGitty(s string) error {
	return errors.New("Error: " + s)
}

func verifyArgCountEqual(args []cli.Argument, count int) (err error) {
	argCount := len(args)

	if argCount < count {
		errStr := fmt.Sprintf("Not enough args, expected %d", count)
		return errGitty(errStr)
	} else if argCount > count {
		errStr := fmt.Sprintf("Too many args, expected %d", count)
		return errGitty(errStr)
	} else {
		return nil
	}
}

func verifyArgCountLessThan(args []cli.Argument, count int) (err error) {
	argCount := len(args)

	if argCount >= count {
		errStr := fmt.Sprintf("Too many args, expected %d", count)
		return errGitty(errStr)
	} else {
		return nil
	}
}
