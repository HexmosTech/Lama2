package methods

import (
	execute "github.com/HexmosTech/lama2/execute"
	"github.com/HexmosTech/lama2/l2lsp/request"
	"github.com/HexmosTech/lama2/l2lsp/response"
	"github.com/rs/zerolog/log"
)

func ExecuteCommand(req request.JSONRPCRequest) response.JSONRPCResponse {

	filePath := "/home/lince/hexmos/apihub/apihub/fb_backend_v3/search.l2"

	jsonResponse, err := execute.ProcessForLSP(filePath)
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	return response.JSONRPCResponse{
		Result: jsonResponse,
	}
}
