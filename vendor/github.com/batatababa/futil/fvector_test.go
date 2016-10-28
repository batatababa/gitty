package futil

import (
	"fmt"
	"os"
	"testing"
)

var testPath = "/tmp/fvectorBacking"

func TestFVectorAdd(t *testing.T) {
	v, err := NewFVector(testPath)
	v.Clear()

	assertFVectorLen(t, v, 0)
	s1 := "test string 1"
	s2 := "test string 2"
	s3 := "test string 3"
	v.Add(s1)
	v.Add(s2)
	v.Add(s3)

	v1, err := v.At(0)
	checkErr(t, err)

	v2, err := v.At(1)
	checkErr(t, err)

	v3, err := v.At(2)
	checkErr(t, err)

	assertStringEqual(t, v1, s1, "")
	assertStringEqual(t, v2, s2, "")
	assertStringEqual(t, v3, s3, "")

	assertFVectorLen(t, v, 3)
}

func TestFVectorClear(t *testing.T) {
	v, _ := NewFVector(testPath)
	v.Clear()

	s1 := "test string 1"
	v.Add(s1)
	v.Add(s1)
	v.Add(s1)

	assertFVectorLen(t, v, 3)
	v.Clear()
	assertFVectorLen(t, v, 0)
}

func TestFVectorRemoveAt(t *testing.T) {
	v, _ := NewFVector(testPath)
	v.Clear()

	assertFVectorLen(t, v, 0)

	s1 := "test string 1"
	s2 := "test string 2"
	s3 := "test string 3"
	v.Add(s1)
	v.Add(s2)
	v.Add(s3)

	if _, err := v.RemoveAt(4); err == nil {
		failTest(t, "RemoveAt(index > len)")
	}

	if _, err := v.RemoveAt(-1); err == nil {
		failTest(t, "RemoveAt(-1)")

	}

	v1, err := v.RemoveAt(0)
	checkErr(t, err)
	assertFVectorLen(t, v, 2)

	v2, err := v.RemoveAt(0)
	checkErr(t, err)
	assertFVectorLen(t, v, 1)

	v3, err := v.RemoveAt(0)
	checkErr(t, err)
	assertFVectorLen(t, v, 0)

	assertStringEqual(t, v1, s1, "")
	assertStringEqual(t, v2, s2, "")
	assertStringEqual(t, v3, s3, "")
}

func TestFVectorRemoveString(t *testing.T) {
	v, _ := NewFVector(testPath)
	v.Clear()

	assertFVectorLen(t, v, 0)

	s1 := "test string 1"
	s2 := "test string 2"
	s3 := "test string 3"
	v.Add(s1)
	v.Add(s2)
	v.Add(s3)

	err := v.Remove(s2)
	checkErr(t, err)
	assertFVectorLen(t, v, 2)
	v1, err := v.At(0)
	v3, err := v.At(1)
	assertStringEqual(t, v1, s1, "")
	assertStringEqual(t, v3, s3, "")

	err = v.Remove(s3)
	checkErr(t, err)
	assertFVectorLen(t, v, 1)
	v1, err = v.At(0)
	assertStringEqual(t, v1, s1, "")

	err = v.Remove(s1)
	checkErr(t, err)
	assertFVectorLen(t, v, 0)

	if err = v.Remove(s3); err == nil {
		t.Errorf("Vector should be empty.")
	}
}

func TestFVectorContains(t *testing.T) {
	v, _ := NewFVector(testPath)
	v.Clear()

	assertFVectorLen(t, v, 0)

	s1 := "test string 1"
	s2 := "test string 2"
	s3 := "test string 3"
	v.Add(s1)
	v.Add(s2)
	v.Add(s3)

	found, err := v.Contains(s1)
	checkErr(t, err)
	if found == false {
		failTest(t, "TestFVectorContains: should be found")
	}

	found, err = v.Contains(s2)
	checkErr(t, err)
	if found == false {
		failTest(t, "TestFVectorContains: should be found")
	}
	found, err = v.Contains(s3)
	checkErr(t, err)
	if found == false {
		failTest(t, "TestFVectorContains: should be found")
	}

	found, err = v.Contains("afdsf")
	checkErr(t, err)
	if found == true {
		failTest(t, "TestFVectorContains: should be Not be found")
	}

	_, err = v.Contains("")
	if err == nil {
		failTest(t, "TestFVectorContains: should fail on empty string")
	}

}

func assertStringEqual(t *testing.T, s1 string, s2 string, msg string) {
	if s1 != s2 {
		if msg == "" {
			msg = fmt.Sprintf("\"%s\" != \"%s\" ", s1, s2)
		}
		failTest(t, msg)
	}
}

func failTest(t *testing.T, msg string) {
	// debug.PrintStack()
	// t.Errorf(msg)
	fmt.Println(msg)
	os.Exit(1)
}

func assertFVectorLen(t *testing.T, v fvector, val int) {
	length, err := GetLineCount(v.file)
	checkErr(t, err)
	assertIntEqual(t, v.Size(), val, "")
	assertIntEqual(t, length, val, "")
}

func assertIntEqual(t *testing.T, first int, sec int, msg string) {
	if first != sec {
		if msg == "" {
			msg = fmt.Sprintf("\"%d\" != \"%d\"", first, sec)
		}
		failTest(t, msg)
	}
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		failTest(t, err.Error())

	}
}
