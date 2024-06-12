//go:build wasm

// Package `utils` provides useful functions for
// simplifying various programming tasks
package utils

import (
	"encoding/json"
	"fmt"

	"os"
)

// ChangeWorkingDirectory tries to set the CWD; on failure
// it exits with a log error message
func ChangeWorkingDir(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		fmt.Println(string(err.Error()))
	}
}

func MarshalAndPrintJSON(data interface{}) {
	filteredJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		os.Exit(0)
	}
	fmt.Println(string(filteredJSON))
}
