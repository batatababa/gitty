package main

import "testing"

func debugEntry(t *testing.T) {
	runCli([]string{"gitty", "-h"})
}
