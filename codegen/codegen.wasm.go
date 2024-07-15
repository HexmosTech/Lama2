//go:build wasm

package codegen

import (
	_ "embed"
	"fmt"

	"github.com/HexmosTech/gabs/v2"
)

func GenerateTargetCode(targetLangLib string, parsedAPI *gabs.Container) string {

	convertedSnippetFinal := generateConvertedSippet(targetLangLib, parsedAPI)
	fmt.Println("This is the converted snippet:", convertedSnippetFinal)
	return convertedSnippetFinal
}
