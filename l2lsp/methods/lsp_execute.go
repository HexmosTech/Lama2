package methods

import (
	"os"

	"github.com/HexmosTech/gabs/v2"
	controller "github.com/HexmosTech/lama2/controller"
	"github.com/HexmosTech/lama2/l2lsp/request"
	"github.com/HexmosTech/lama2/l2lsp/response"
	outputmanager "github.com/HexmosTech/lama2/outputManager"
	"github.com/HexmosTech/lama2/parser"
	"github.com/HexmosTech/lama2/preprocess"
	"github.com/HexmosTech/lama2/utils"
)

/*
ExecuteCommand is the LSP method that handles the execution of the API file
filePath: the path to the API file
*/
func ExecuteCommand(req request.JSONRPCRequest) response.JSONRPCResponse {
	jsonResponse, err := ProcessForLSP(req.Params.FilePath)
	if err != nil {
		return response.ErrorResp(req, 500, err.Error())
	}
	resp := response.CreateSuccessResponse(req.ID, jsonResponse)

	return resp
}

/*
ProcessForLSP is the method that processes the API file
filePath: the path to the API file
*/
func ProcessForLSP(filePath string) (*gabs.Container, error) {
	apiContent, err := preprocess.GetLamaFileAsString(filePath)
	if err != nil {
		return nil, err
	}
	_, dir, _ := utils.GetFilePathComponents(filePath)
	oldDir, _ := os.Getwd()
	utils.ChangeWorkingDir(dir)

	preprocess.LoadEnvironments(dir)
	utils.ChangeWorkingDir(oldDir)
	p := parser.NewLama2Parser()
	parsedAPI, e := p.Parse(apiContent)
	if e != nil {
		return nil, e
	}

	resp, _, responseTime, statusCodes, contentSizes, err := controller.HandleParsedFileHelper(parsedAPI)
	if err != nil {
		return nil, err
	}

	jsonResponse, err := outputmanager.ResponseToJSON(resp, responseTime, statusCodes, contentSizes)
	if err != nil {
		return nil, err
	}

	return jsonResponse, nil
}