package futil

import (
	"fmt"
	"os"
	"strings"
)

type fvector struct {
	path   string
	file   *os.File
	length int
}

func NewFVector(path string) (v fvector, err error) {
	v.path = path
	if v.file, err = GetFile(path); err != nil {
		return v, err
	}

	if v.length, err = GetLineCount(v.file); err != nil {
		return v, err
	}
	return v, err
}

func (v *fvector) Add(s string) (err error) {
	v.length++
	err = AppendFile(v.file, s)
	if err != nil {
		v.length, _ = GetLineCount(v.file)
	}
	return err
}

func (v *fvector) Remove(s string) (err error) {
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

func (v *fvector) RemoveAt(index int) (removed string, err error) {
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

func (v *fvector) Contains(s string) (found bool, err error) {
	return found, err
}

func (v *fvector) Trim() (err error) {
	ScanReplace(v.file, func(line *string) error {
		*line = strings.TrimSpace(*line)
		return nil
	})

	return err
}

func (v *fvector) Clear() (err error) {
	v.length = 0
	err = ClearFile(v.file)
	if err != nil {
		v.length, _ = GetLineCount(v.file)
	}

	return err
}

func (v *fvector) At(index int) (s string, err error) {
	s, err = GetLine(v.file, index)
	return s, err
}

func (v *fvector) Size() (size int) {
	return v.length
}
