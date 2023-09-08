// lsp_suggestEnv.go
package l2lsp

import (
	"path/filepath"
	"strings"

	l2envpackege "github.com/HexmosTech/lama2/l2env"
	"github.com/rs/zerolog/log"
)

func getSearchQueryString(request JSONRPCRequest) string {
	if request.Params.SearchQuery != nil {
		return *request.Params.SearchQuery
	}
	return ""
}

func getRequestURI(request JSONRPCRequest) string {
	if request.Params.TextDocument.Uri == nil {
		return ""
	}
	uri := *request.Params.TextDocument.Uri

	// Handle local files
	if strings.HasPrefix(uri, "file://") {
		return uri[len("file://"):]

		// Handle remote files (example for SSH)
	} else if strings.HasPrefix(uri, "ssh://") {
		// Extract and return the path after the SSH scheme
		// You might need additional logic to handle SSH URIs properly
		return uri[len("ssh://"):]

		// Handle WSL files
	} else if strings.HasPrefix(uri, "wsl://") {
		// Extract and return the path after the WSL scheme
		// Additional logic might be needed for WSL URIs
		return uri[len("wsl://"):]

	} else {
		// Log the unexpected URI scheme
		log.Warn().Str("URI", uri).Msg("Encountered unexpected URI scheme.")
	}
	return ""
}

func SuggestEnvironmentVariables(request JSONRPCRequest) JSONRPCResponse {
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

	log.Info().Msg("L2 LSP environment variables suggestion requested")
	log.Info().Str("Method", request.Method).Interface("Params", request.Params)

	relevantSearchString := getSearchQueryString(request)
	uri := getRequestURI(request)
	if uri == "" {
		return ErrorResp(request, "Document uri is improper. Ex: 'file:///path/to/workspace/myapi.l2'")
	}
	parentFolder := filepath.Dir(uri)
	res := l2envpackege.ProcessEnvironmentVariables(relevantSearchString, parentFolder)
	return createEnvironmentVariablesResponse(request, res)
}
