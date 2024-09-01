
// Package `lama2cmd` provides CLI argument parsing facilities.
// It hosts the `Opts` structure to record user intentions
package lama2cmd

import (
	"fmt"
	"os"

	"github.com/HexmosTech/lama2/importer"
	"github.com/HexmosTech/lama2/l2lsp"
	outputmanager "github.com/HexmosTech/lama2/outputManager"
	"github.com/HexmosTech/lama2/utils"
	"github.com/jessevdk/go-flags"
	"github.com/rs/zerolog/log"
)

// The Opts structure stores user preferences, and is used throughout
// the module to make various decisions.
type Opts struct {
	Output   string `short:"o" long:"output" description:"Path to output JSON file to store logs, headers and result"`
	Verbose  []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	Prettify bool   `short:"b" long:"prettify" description:"Prettify specified .l2 file"`
	// Sort     bool   `short:"s" long:"sort" description:"Sort specification into recommended order"`
	Convert     string `short:"c" long:"convert" description:"Generate code in given language and library (ex: python.requests); reference: tinyurl.com/l2codegen"`
	Nocolor     bool   `short:"n" long:"nocolor" description:"Disable color in httpie output"`
	Update      bool   `short:"u" long:"update" description:"Update l2 binary to the latest released version (Linux/MacOS only)"`
	PostmanFile string `short:"p" long:"postmanfile" description:"JSON export from Postman (Settings -> Data -> Export Data)"`
	LamaDir     string `short:"l" long:"lama2dir" description:"Output directory to put .l2 files after conversion from Postman format"`
	Help        bool   `short:"h" long:"help" group:"AddHelp" description:"Usage help for Lama2"`
	Lsp         bool   `short:"z" long:"lsp" description:"Start the lsp server"`
	Version     bool   `long:"version" description:"Print Lama2 binary version"`

	Positional struct {
		LamaAPIFile string
	} `positional-args:"yes"`
}

func getParsedInput(argList []string) (Opts, []string) {
	argList = argList[1:] // remove command name
	o := Opts{}

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

func ArgParsing(o *Opts, version string) {
	if o.Version {
		fmt.Println(version)
		os.Exit(0)
	}
	if o.Update {
		utils.UpdateSelf()
		os.Exit(0)
	}
	if o.Lsp {
		l2lsp.StartLspServer()
		// Incoming requests to the LSP will be handled by l2lsp.Process()
	}
	if len(o.PostmanFile) > 0 {
		if len(o.LamaDir) > 0 {
			importer.PostmanImporter(o.PostmanFile, o.LamaDir)
			os.Exit(0)
		}
		log.Fatal().Msg("To convert Postman export to Lama2, try: l2 -p PostmanFile -l Lama2Dir")
		os.Exit(1)
	}
	if len(o.LamaDir) > 0 {
		if len(o.PostmanFile) > 0 {
			importer.PostmanImporter(o.PostmanFile, o.LamaDir)
			os.Exit(0)
		}
		log.Fatal().Msg("To convert Postman export to Lama2, try: l2 -p PostmanFile -l Lama2Dir")
		os.Exit(1)
	}
}

// GetAndValidateCmd takes in the user's CLI input, and checks
// for validity. If not OK, displays a help message in stdout.
// Otherwise, fills the Opts structure and returns it
// Moreover, based on user input, the outputManager gets configured
// (whether user prefers trace/debug/info level)
func GetAndValidateCmd(ipArgs []string) *Opts {
	o, _ := getParsedInput(ipArgs)
	return &o
}
