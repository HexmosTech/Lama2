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

// The following string fragment is used to mark unquoted
// variables within a JSON body in the lama2 file.
// For example, Lama2 parser will mark an unquoted variable
// `hello` as:
//
// "<626f4c60-${hello}>"
//
// The above string is actually representative of the user intent:
//
// ${hello}
//
// Note the lack of quotes around the variable above. This hack
// is needed to work around Gabs, which is unaware of variables
// required in the Lama2 system
const UNQUOTED_VAR_MARKER = "626f4c60"

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
