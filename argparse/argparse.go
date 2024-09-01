package argparse

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/HexmosTech/lama2/importer"
	"github.com/HexmosTech/lama2/l2lsp"
	"github.com/HexmosTech/lama2/options"
	"github.com/HexmosTech/lama2/utils"
)

func ArgParsing(o *options.Opts, version string) {
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
