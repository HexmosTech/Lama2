//go:build cli

// Package `utils` provides useful functions for
// simplifying various programming tasks
package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
)


// ChangeWorkingDirectory tries to set the CWD; on failure
// it exits with a log error message
func ChangeWorkingDir(dir string) {
	if dir == "" {
		log.Error().
			Str("Type", "Preprocess").
			Msg("Directory path is empty")
		return
	}

	log.Debug().
		Str("Type", "Preprocess").
		Str("dir", dir).
		Msg("Attempting to change directory")

	err := os.Chdir(dir)
	if err != nil {
		log.Fatal().
			Str("Type", "Preprocess").
			Str("dir", dir).
			Msg("Moving into dir failed")
	}
}

func MarshalAndPrintJSON(data interface{}) {
	filteredJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Error().Str("Type", "Preprocess").Msg(fmt.Sprintf("Failed to marshal JSON: %v", err))
		os.Exit(0)
	}
	fmt.Println(string(filteredJSON))
}
