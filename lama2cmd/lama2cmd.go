// Package `lama2cmd` provides CLI argument parsing facilities.
// It hosts the `Opts` structure to record user intentions
package lama2cmd

import (
	"os"

	options "github.com/HexmosTech/lama2/options"
	outputmanager "github.com/HexmosTech/lama2/outputManager"
	"github.com/jessevdk/go-flags"
	"github.com/rs/zerolog/log"
)

// The Opts structure stores user preferences, and is used throughout
// the module to make various decisions.

func getParsedInput(argList []string) (options.Opts, []string) {
	argList = argList[1:] // remove command name
	o := options.Opts{}

	if len(argList) == 0 {
		argList = append(argList, "-h")
	}
	args, err := flags.ParseArgs(&o, argList)
	if err != nil {
		e, _ := err.(*flags.Error)
		if e.Type == flags.ErrHelp {
			os.Exit(0)
		}

		log.Fatal().
			Str("Type", "Preprocess").
			Strs("arglist", argList).
			Msg("Couldn't parse argument list")
	}

	switch len(o.Verbose) {
	case 0:
		outputmanager.ConfigureZeroLog("INFO")
	case 1:
		outputmanager.ConfigureZeroLog("DEBUG")
	case 2:
		outputmanager.ConfigureZeroLog("TRACE")
	default:
		outputmanager.ConfigureZeroLog("DEBUG")
	}

	log.Debug().
		Str("Type", "Preprocess").
		Bools("Verbosity", o.Verbose).
		// Bool("Prettify", o.Prettify).
		// Bool("Sort", o.Sort).
		Bool("NoColor", o.Nocolor).
		Str("Lama API File", o.Positional.LamaAPIFile).
		Msg("Parsed inputs")

	return o, args
}

// GetAndValidateCmd takes in the user's CLI input, and checks
// for validity. If not OK, displays a help message in stdout.
// Otherwise, fills the Opts structure and returns it
// Moreover, based on user input, the outputManager gets configured
// (whether user prefers trace/debug/info level)
func GetAndValidateCmd(ipArgs []string) *options.Opts {
	o, _ := getParsedInput(ipArgs)
	return &o
}
