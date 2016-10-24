package futil

import (
	"bufio"
	"io"
	"os"
)

func ContainsLine(f *os.File, s string) (lineNum int, err error) {
	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)
	lineNum = 1

	for scanner.Scan() {
		text := scanner.Text()
		if text == s {
			return lineNum, nil
		}
		lineNum++
	}

	err = scanner.Err()

	if err == nil {
		err = io.EOF
	}

	return lineNum, err
}

func GetFile(filePath string) (f *os.File, err error) {
	return os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND|O_TRUNC, 0600)
}

func AppendFile(f *os.File, s string) (err error) {
	info, err := f.Stat()
	appendEnabled := info.Mode() & os.ModeAppend
	if appendEnabled != 0 {
		f.Seek(0, 2)
	}
	_, err = f.WriteString(s)
	return err
}

func ClearFile(f *os.File) (err error) {
	filePath := f.Name()

	if err = os.Remove(filePath); err != nil {
		return nil
	}

	f, err = Get(filePath)

	return err
}

func RemoveLineNumber(f *os.File, lineNum int) (removed string, err error) {
	linesFound := 1
	ScanReplace(f, func(line *string) error {
		if lineNum == linesFound {
			removed = *line
			*line = ""
		}
		linesFound++

		return err
	})

	return removed, err
}

func RemoveContaining(f *os.File, s string) (err error) {

	ScanReplace(f, func(line *string) error {
		if *line == s {
			*line = ""
		}

		return nil
	})

	return err
}

func ScanReplace(f *os.File, op func(line *string) error) (err error) {
	filePath := f.Name()
	tempFilePath := filePath + "_temp"

	scanner := bufio.NewScanner(f)
	os.Remove(tempFilePath)
	tempFile, err := Get(tempFilePath)

	if err != nil {
		return err
	}

	for scanner.Scan() {
		text := scanner.Text()

		if err = op(&text); err != nil {
			os.Remove(tempFilePath)
			return err
		}

		if text != "" {
			tempFile.WriteString(text + "\n")
		}
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	// if err = os.Remove(filePath); err != nil {
	// 	return err
	// }

	// if err = os.Rename(tempFilePath, filePath); err != nil {
	// 	return err
	// }

	return err
}
