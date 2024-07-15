//go:build cli

package codegen

import (
	_ "embed"
	"fmt"

	"github.com/HexmosTech/gabs/v2"
	"github.com/atotto/clipboard"
)

func GenerateTargetCode(targetLangLib string, parsedAPI *gabs.Container) {
	convertedSnippetFinal := generateConvertedSippet(targetLangLib, parsedAPI)
	fmt.Println("This is the converted snippet:", convertedSnippetFinal)
	clipboard.WriteAll(convertedSnippetFinal)
	fmt.Println("Code copied to clipboard")
}
