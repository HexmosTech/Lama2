package methods

import (
	"os"

	"github.com/HexmosTech/lama2/codegen"
	"github.com/HexmosTech/lama2/l2lsp/request"
	"github.com/HexmosTech/lama2/l2lsp/response"
	"github.com/HexmosTech/lama2/parser"
	"github.com/HexmosTech/lama2/preprocess"
	"github.com/HexmosTech/lama2/utils"
)

func CodeGeneration(req request.JSONRPCRequest) response.JSONRPCResponse {
	filePath:= req.Params.FilePath
	Language:= req.Params.Language
	client := req.Params.Client
	generatedCode, err := ProcessCodeGeneration(filePath, Language, client)
	if err != nil {
		return response.ErrorResp(req, 500, err.Error())
	}
	return response.CreateSuccessResponse(req.ID, generatedCode)
}

func ProcessCodeGeneration(filePath, Language, client string) (string, error) {
	languageMethod := Language + "." + client
	apiContent, err := preprocess.GetLamaFileAsString(filePath)
	if err != nil {
		return "", err
	}
	_, dir, _ := utils.GetFilePathComponents(filePath)
	oldDir, _ := os.Getwd()
	utils.ChangeWorkingDir(dir)

	preprocess.LoadEnvironments(dir)
	utils.ChangeWorkingDir(oldDir)
	p := parser.NewLama2Parser()
	parsedAPI, e := p.Parse(apiContent)
	if e != nil {
		return "", e
	}
	generatedCode := codegen.GenerateTargetCode(languageMethod, parsedAPI)
	
	return generatedCode, nil


}