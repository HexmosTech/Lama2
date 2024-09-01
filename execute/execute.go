package execute

import (
	"fmt"
	"os"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/httpie-go"
	"github.com/HexmosTech/lama2/cmdexec"
	helpers "github.com/HexmosTech/lama2/helpers"
	"github.com/HexmosTech/lama2/l2lsp/request"
	"github.com/HexmosTech/lama2/l2lsp/response"
	"github.com/HexmosTech/lama2/lama2cmd"
	"github.com/HexmosTech/lama2/options"
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
	o := lama2cmd.GetAndValidateCmd([]string{"", filePath})

	apiContent := preprocess.GetLamaFileAsString(filePath)
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
	resp, responseTime, statusCodes, contentSizes := handleParsedFileForLSP(parsedAPI, o, dir)

	jsonResponse, err := outputmanager.ResponseToJSON(resp, responseTime, statusCodes, contentSizes)
	// log.Info().RawJSON("JSON Response", jsonResponse.Bytes()).Msg("")
	if err != nil {
		return nil, err
	}

	return jsonResponse, nil
}

func handleParsedFileForLSP(parsedAPI *gabs.Container, o *options.Opts, dir string) (httpie.ExResponse, []outputmanager.ResponseTime, []outputmanager.StatusCode, []outputmanager.ContentSize) {
	parsedAPIblocks := helpers.GetParsedAPIBlocks(parsedAPI)
	vm := cmdexec.GetJSVm()
	var resp httpie.ExResponse
	var responseTime []outputmanager.ResponseTime
	var statusCodes []outputmanager.StatusCode
	var contentSizes []outputmanager.ContentSize
	var timeInMs int64

	for i, block := range parsedAPIblocks {
		log.Debug().Int("Block num", i).Msg("")
		log.Debug().Str("Block getting processed", block.String()).Msg("")
		blockType := block.S("type").Data().(string)
		if blockType == "processor" {
			helpers.ExecuteProcessorBlock(block, vm)
		} else if blockType == "Lama2File" {
			resp, timeInMs = helpers.ExecuteRequestorBlock(block, vm, o, dir)
			log.Info().Str("ResponseTime", fmt.Sprintf("%dms", timeInMs)).Msg("")
			responseTime, statusCodes, contentSizes = helpers.CalculateMetrics(resp, timeInMs, responseTime, statusCodes, contentSizes)
		}
	}

	return resp, responseTime, statusCodes, contentSizes
}
