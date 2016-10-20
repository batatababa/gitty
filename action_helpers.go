package main

import (
	"errors"
	"fmt"

	"github.com/batatababa/cli"
)

func toErrString(s string) string {
	s = fmt.Sprintf(" Error: %s", s)
	return s
}

func verifyArgCountEqual(args []cli.Argument, count int) (err error) {
	argCount := len(args)

	if argCount < count {
		errStr := fmt.Sprintf("Not enough args, expected %d", count)
		errStr = toErrString(errStr)
		return errors.New(errStr)
	} else if argCount > count {
		errStr := fmt.Sprintf("Too many args, expected %d", count)
		errStr = toErrString(errStr)
		return errors.New(errStr)
	} else {
		return nil
	}
}
