package futil

import (
	"os"
	"strings"
)

type fvector struct {
	path string
	file *os.File
}

func NewFVector(path string) (v fvector, err error) {
	v.path = path
	v.file, err = Get(path)
	return v, err
}

func (v *fvector) Add(s string) (err error) {
	err = AppendFile(v.file, s)
	return err
}

func (v *fvector) Remove(s string) (err error) {
	err = RemoveContaining(v.file, s)
	return err
}

func (v *fvector) RemoveAt(index int) (removed string, err error) {
	removed, err = RemoveLineNumber(v.file, index)
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
	err = ClearFile(v.Clear())
	return err
}

func (v *fvector) At(index int, s string) (err error) {
	return err
}
