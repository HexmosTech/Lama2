// lsp_suggestEnv.go
package methods

import (
	"errors"
	"path/filepath"
	"strings"

	l2envpackege "github.com/HexmosTech/lama2/l2env"
	"github.com/HexmosTech/lama2/l2lsp/request"
	"github.com/HexmosTech/lama2/l2lsp/response"
	// "github.com/rs/zerolog/log"
)

func getSearchQueryString(req request.JSONRPCRequest) string {
	if req.Params.SearchQuery != nil {
		return *req.Params.SearchQuery
	}
	return ""
}

func getRequestURI(req request.JSONRPCRequest) (string, int, error) {
	if req.Params.TextDocument.URI == nil {
		return "", response.ErrInvalidURI, errors.New("URI cannot be empty. Ex: 'file:///path/to/workspace/myapi.l2'")
	}
	uri := *req.Params.TextDocument.URI

	// Handle local files
	if strings.HasPrefix(uri, "file://") {
		return uri[len("file://"):], 0, nil

		// Handle remote files (example for SSH)
	} else if strings.HasPrefix(uri, "ssh://") {
		return "", response.ErrUnsupportedFeature, errors.New("SSH is not supported as of now. To contribute visit here: https://github.com/HexmosTech/Lama2")

		// Handle WSL files
	} else if strings.HasPrefix(uri, "wsl://") {
		return "", response.ErrUnsupportedFeature, errors.New("WSL is not supported as of now. To contribute visit here: https://github.com/HexmosTech/Lama2")

		// Handle Windows files
	} else if strings.Contains(uri, "\\") {
		return "", response.ErrUnsupportedFeature, errors.New("Windows is not supported as of now. To contribute visit here: https://github.com/HexmosTech/Lama2")
	} else {
		// Log the unexpected URI scheme
		// log.Warn().Str("URI", uri).Msg("Encountered unexpected URI scheme.")
		return "", response.ErrUnexpectedURIScheme, errors.New("encountered unexpected URI scheme. Ex: 'file:///path/to/workspace/myapi.l2'")
	}
}

func SuggestEnvironmentVariables(req request.JSONRPCRequest) response.JSONRPCResponse {
	/*
		{
			"jsonrpc": "2.0",
			"id": 2,
			"method": "suggest/environmentVariables",
			"params": {
				"textDocument": {
					"uri": "file:///home/lovestaco/repos/Lama2/elfparser/ElfTestSuite/root_variable_override/api/y_0020_root_override.l2"
				},
				"position": {
					"line": 1,
					"character": 2
				},
				"relevantSearchString": ""
			}
		}
	*/

	// log.Info().Msg("L2 LSP environment variables suggestion requested")
	// log.Info().Str("Method", req.Method).Interface("Params", req.Params)

	searchQuery := getSearchQueryString(req)
	// log.Debug().Str("Method", req.Method).Interface("searchQuery", searchQuery)
	uri, errorCode, err := getRequestURI(req)
	if err != nil {
		return response.ErrorResp(req, errorCode, err.Error())
	}
	// log.Debug().Str("Method", req.Method).Interface("uri", uri)
	parentFolder := filepath.Dir(uri)
	// log.Debug().Str("Method", req.Method).Interface("parentFolder", parentFolder)
	res := l2envpackege.ProcessEnvironmentVariables(searchQuery, parentFolder)
	return response.CreateEnvironmentVariablesResp(req, res)
}
