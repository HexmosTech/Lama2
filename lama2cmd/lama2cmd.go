package lama2cmd

import (
	"fmt"
	"os"

	"github.com/HexmosTech/lama2/utils"
	"github.com/jessevdk/go-flags"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

type Opts struct {
	Verbose  []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	Prettify bool   `short:"p" long:"prettify" description:"Prettify specified .lama file"`
	Sort     bool   `short:"s" long:"sort" description:"Sort specification into recommended order"`
	Nocolor  bool   `short:"n" long:"nocolor" description:"Disable color in httpie output"`
	Help     bool   `short:"h" long:"help" group:"AddHelp" description:"Usage help for Lama2"`

	Positional struct {
		LamaAPIFile string `required:"yes"`
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
			Msg(fmt.Sprint("Couldn't parse argument list"))
	}

	if o.Nocolor {
		utils.ConfigureZeroLog("ERROR")
	} else {
		switch len(o.Verbose) {
		case 0:
			utils.ConfigureZeroLog("INFO")
		case 1:
			utils.ConfigureZeroLog("DEBUG")
		case 2:
			utils.ConfigureZeroLog("TRACE")
		default:
			utils.ConfigureZeroLog("TRACE")
		}
	}

	log.Debug().
		Str("Type", "Preprocess").
		Bools("Verbosity", o.Verbose).
		Bool("Prettify", o.Prettify).
		Bool("Sort", o.Sort).
		Bool("NoColor", o.Nocolor).
		Str("Lama API File", o.Positional.LamaAPIFile).
		Msg("Parsed inputs")

	return o, args
}

func GetAndValidateCmd(ipArgs []string) *Opts {
	o, _ := getParsedInput(ipArgs)
	return &o
}
