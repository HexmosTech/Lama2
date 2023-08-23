package prettify

import (
	"fmt"
	"os"
	"strconv"

	"github.com/HexmosTech/gabs/v2"
)

func Prettify(parsedAPI *gabs.Container, context map[string]bool, markRange map[string]int, content string, fPath string) {
	fmt.Println(markRange)
	parsedAPIblocks := parsedAPI.S("value").Data().(*gabs.Container).Children()
	// Prettification procedure:
	// 1. Scan from bottom of contents
	// 2. For every JSON fragment:
	// 		2.1. reformat
	//      2.2. replace
	//      2.3. write output
	l := len(parsedAPIblocks)
	markMax := (len(markRange) / 2) - 1
	for i := l - 1; i >= 0; i-- {
		block := parsedAPIblocks[i]
		// vm := cmdexec.GetJSVm()
		blockType := block.S("type").Data().(string)
		if blockType == "processor" {
			fmt.Println("Skipping processor block")
		} else if blockType == "Lama2File" {
			// preprocess.ProcessVarsInBlock(block, vm)
			isJSON, ok := context["JSON"]
			if !ok {
				return
			}

			if !isJSON {
				return
			}

			fmt.Println("Prettifying JSON in the l2 file")
			jsonObj := block.S("details", "ip_data")

			idxStr := strconv.Itoa(markMax)

			content = content[:markRange["DataStart"+idxStr]] + jsonObj.StringIndent("", "  ") + "\n" + content[markRange["DataEnd"+idxStr]:]
			markMax -= 1
		}
		os.WriteFile(fPath, []byte(content), 0644)
	}
}
