package futil

import (
	"bufio"
	"io"
	"os"
)

func ContainsLine(f *os.File, s string) (lineNum int, err error) {
	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	if _, err = f.Seek(0, 0); err != nil {
		return 0, err
	}

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

func GetLine(f *os.File, lineNum int) (s string, err error) {
	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	if _, err = f.Seek(0, 0); err != nil {
		return "", err
	}

	curLineNum := 0
	for scanner.Scan() {
		if curLineNum == lineNum {
			err = scanner.Err()
			s = scanner.Text()
			return s, err
		}
		curLineNum++
	}

	err = scanner.Err()

	return s, err
}

func GetLineCount(f *os.File) (lineCount int, err error) {
	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)
	lineCount = 0

	if _, err = f.Seek(0, 0); err != nil {
		return lineCount, err
	}

	for scanner.Scan() {
		lineCount++
	}

	err = scanner.Err()

	return lineCount, err
}

func GetFile(filePath string) (f *os.File, err error) {
	return os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND|os.O_SYNC, 0600)
}

func AppendFile(f *os.File, s string) (err error) {
	info, err := f.Stat()
	appendEnabled := info.Mode() & os.ModeAppend
	if appendEnabled != 0 {
		f.Seek(0, 2)
	}
	_, err = f.WriteString(s + "\n")
	f.Sync()
	return err
}

func ClearFile(f *os.File) (err error) {
	if err = f.Truncate(0); err != nil {
		f.Sync()
		return err
	}
	err = f.Sync()
	return err
}

func RemoveLineNumber(f *os.File, lineNum int) (removed string, err error) {
	linesFound := 0
	err = ScanReplace(f, func(line *string) error {
		if lineNum == linesFound {
			removed = *line
			*line = ""
		}
		linesFound++

		return nil
	})

	return removed, err
}

func RemoveContaining(f *os.File, s string) (err error) {

	err = ScanReplace(f, func(line *string) error {
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
	if _, err = f.Seek(0, 0); err != nil {
		return err
	}

	tempFile, err := GetFile(tempFilePath)
	if err != nil {
		return err
	}

	ClearFile(tempFile)

	tempWriter := bufio.NewWriter(tempFile)

	for scanner.Scan() {
		text := scanner.Text()

		if err = op(&text); err != nil {
			os.Remove(tempFilePath)
			return err
		}

		if text != "" {
			tempWriter.WriteString(text + "\n")
		}
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	tempWriter.Flush()
	tempFile.Close()
	f.Close()

	if err = os.Remove(filePath); err != nil {
		return err
	}

	if err = os.Rename(tempFilePath, filePath); err != nil {
		return err
	}

	tempFile, err = GetFile(filePath)
	*f = *tempFile

	return err
}
