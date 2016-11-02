package futil

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type FVector struct {
	path   string
	file   *os.File
	length int
}

func NewFVector(path string) (v FVector, err error) {
	v.path = path
	if v.file, err = GetFile(path); err != nil {
		return v, err
	}

	if v.length, err = GetLineCount(v.file); err != nil {
		return v, err
	}
	return v, err
}

func (v *FVector) Add(s string) (err error) {
	v.length++
	err = AppendFile(v.file, s)
	if err != nil {
		v.length, _ = GetLineCount(v.file)
	}
	return err
}

func (v *FVector) Remove(s string) (err error) {
	if v.length == 0 {
		return fmt.Errorf("Can't remove from empty vector.")
	}
	v.length--
	err = RemoveContaining(v.file, s)
	if err != nil {
		v.length, _ = GetLineCount(v.file)
	}
	return err
}

func (v *FVector) RemoveAt(index int) (removed string, err error) {
	if index >= v.length || index < 0 {
		return "", fmt.Errorf("Invalid remove index, %d", index)
	}
	v.length--
	removed, err = RemoveLineNumber(v.file, index)
	if err != nil {
		v.length, _ = GetLineCount(v.file)
	}

	return removed, err
}

func (v *FVector) Contains(s string) (found bool, err error) {
	if s == "" {
		return false, fmt.Errorf("Invalid contains string, the empty string")
	}
	_, err = ContainsLine(v.file, s)

	if err == io.EOF {
		return false, nil
	} else if err != nil && err != io.EOF {
		return false, err
	} else {
		return true, nil
	}
}

func (v *FVector) Trim() (err error) {
	ScanReplace(v.file, func(line *string) error {
		*line = strings.TrimSpace(*line)
		return nil
	})

	return err
}

func (v *FVector) Clear() (err error) {
	v.length = 0
	err = ClearFile(v.file)
	if err != nil {
		v.length, _ = GetLineCount(v.file)
	}

	return err
}

func (v *FVector) At(index int) (s string, err error) {
	s, err = GetLine(v.file, index)
	return s, err
}

func (v *FVector) Size() (size int) {
	return v.length
}

func (v *FVector) Slice() (sl []string, err error) {
	sl = make([]string, 0, 20)

	scanner := bufio.NewScanner(v.file)

	if _, err = v.file.Seek(0, 0); err != nil {
		return nil, err
	}

	for scanner.Scan() {
		text := scanner.Text()
		sl = append(sl, text)
	}
	err = scanner.Err()

	return sl, err
}
