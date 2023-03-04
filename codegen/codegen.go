package codegen

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/rs/zerolog/log"

	"github.com/HexmosTech/lama2/cmdexec"
)

type SnippetArgs struct {
	Language string
	Library  string
}

func PrepareHTTPSnippetGenerator(snippetArgs SnippetArgs) string {
	var templOutput bytes.Buffer
	templStr := `var m = require('./codegen/httpsnippet.js'); 
	const snippet = new window.HTTPSnippet({
  		method: 'GET',
  		url: 'http://mockbin.com/request',
	});
	
	let convertedSnippet = snippet.convert('{{.Language}}'{{if .Library }}, '{{.Library}}'{{end}});
	`
	tmpl, _ := template.New("httpsnippet").Parse(templStr)
	tmpl.Execute(&templOutput, snippetArgs)
	return templOutput.String()
}

func GenerateTargetCode(targetLangLib string) {
	snippetArgs := SnippetArgs{}
	lang, lib := SplitLangLib(targetLangLib)
	snippetArgs.Language = lang
	snippetArgs.Library = lib
	httpsnippetCode := PrepareHTTPSnippetGenerator(snippetArgs)
	vm := cmdexec.GetJSVm()
	_, e := vm.RunString(httpsnippetCode)
	if e != nil {
		log.Fatal().
			Str("Type", "CodeGen").
			Str("Error", e.Error()).
			Msg(fmt.Sprint("Code generator error"))
	}
	fmt.Println(vm.Get("convertedSnippet"))
}
