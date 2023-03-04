package prettify

import (
	"fmt"
	"os"

	"github.com/HexmosTech/gabs/v2"
)

func Prettify(parsedAPI *gabs.Container, context map[string]bool, markRange map[string]int, content string, fPath string) {
	parsedAPIblocks := parsedAPI.S("value").Data().(*gabs.Container).Children()
	// vm := cmdexec.GetJSVm()
	for _, block := range parsedAPIblocks {
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

			res := content[:markRange["DataStart"]] + jsonObj.StringIndent("", "  ") + content[markRange["DataEnd"]:]
			os.WriteFile(fPath, []byte(res), 0644)
		}
	}
}
