package tests

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/HexmosTech/lama2/cmdexec"
	controller "github.com/HexmosTech/lama2/controller"
	"github.com/HexmosTech/lama2/parser"
	"github.com/HexmosTech/lama2/preprocess"
	"github.com/HexmosTech/lama2/utils"
)

func TestBase64(t *testing.T) {

	// parse file
	l2Path := "../examples/0008_base64_image/0008_base64_image.l2"
	apiContent, _ := os.ReadFile(l2Path)
	_, dir, _ := utils.GetFilePathComponents(l2Path)
	nowPwd, _ := os.Getwd()
	utils.ChangeWorkingDir(dir)
	preprocess.LoadEnvFile(path.Join(dir, "l2.env"))
	p := parser.NewLama2Parser()
	parsedAPI, _ := p.Parse(string(apiContent))
	fmt.Println(parsedAPI)

	// get parsed blocks
	parsedAPIblocks := controller.GetParsedAPIBlocks(parsedAPI)
	fmt.Println(parsedAPIblocks)

	vm := cmdexec.GetJSVm()

	// loop through
	for _, block := range parsedAPIblocks {
		preprocess.ProcessVarsInBlock(block, vm)
		b64Str := os.Getenv("PHOTO")
		fmt.Println(len(b64Str))
		if len(b64Str) <= 25000 {
			t.Errorf("Base64 generation failed")
		}
		fmt.Println(block)
	}
	utils.ChangeWorkingDir(nowPwd)

}
