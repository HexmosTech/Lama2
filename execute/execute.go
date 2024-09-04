package execute

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

type Executor struct{}

func NewExecutor() *Executor {
	return &Executor{}
}

func (e *Executor) ExecuteCommand(req request.JSONRPCRequest) response.JSONRPCResponse {
	filePath := "/home/lince/hexmos/apihub/apihub/fb_backend_v3/search.l2"

	jsonResponse, err := ProcessForLSP(filePath)
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	return response.JSONRPCResponse{
		Result: jsonResponse,
	}
}

func ProcessForLSP(filePath string) (*gabs.Container, error) {
	log.Debug().Str("filePath", filePath).Msg("")

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

	log.Debug().Str("Parsed API", parsedAPI.String()).Msg("")
	resp, _, responseTime, statusCodes, contentSizes, err := controller.HandleParsedFileHelper(parsedAPI)
	if err != nil {
		return nil, err
	}

	jsonResponse, err := outputmanager.ResponseToJSON(resp, responseTime, statusCodes, contentSizes)
	// log.Info().RawJSON("JSON Response", jsonResponse.Bytes()).Msg("")
	if err != nil {
		return nil, err
	}

	return jsonResponse, nil
}
