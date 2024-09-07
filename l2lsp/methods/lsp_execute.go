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
	"github.com/rs/zerolog/log"
)
	
func ExecuteCommand(req request.JSONRPCRequest) response.JSONRPCResponse {

	jsonResponse, err := ProcessForLSP(req.Params.FilePath)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	resp := response.CreateExecuteCommandResponse(req, jsonResponse)
	log.Info().Interface("resp", resp).Msg("")

	return resp
}

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