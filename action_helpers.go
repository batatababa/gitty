package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

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

func fileContains(f *os.File, s string) (line int, err error) {
	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)
	line = 1

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), s) {
			return line, nil
		}
		line++
	}

	err = scanner.Err()

	if err == nil {
		err = io.EOF
	}

	return line, err
}

func fileGet(filePath string, flag int) (f *os.File, err error) {
	_, err = os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			f, err = os.Create(filePath)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	f, err = os.OpenFile(filePath, flag, 0600)

	return f, err
}
