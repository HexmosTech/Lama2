package options

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
