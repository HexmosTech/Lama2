package helper

import (
	"github.com/HexmosTech/gabs/v2"
	outputmanager "github.com/HexmosTech/lama2/outputManager"
	"github.com/HexmosTech/httpie-go"
	controller "github.com/HexmosTech/lama2/controller"
)

func HandleParsedFileHelper(parsedAPI *gabs.Container, args ...interface{}) (httpie.ExResponse, []outputmanager.ResponseTime, []outputmanager.StatusCode, []outputmanager.ContentSize) {
    parsedAPIblocks := GetParsedAPIBlocks(parsedAPI)
    o, dir := extractArgs(args)
    return processBlocks(parsedAPIblocks, o, dir)
}
