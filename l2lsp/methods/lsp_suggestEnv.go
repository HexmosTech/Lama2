package methods

import (
	"errors"
	"path/filepath"
	"strings"

	l2envpackege "github.com/HexmosTech/lama2/l2env"
	"github.com/HexmosTech/lama2/l2lsp/request"
	"github.com/HexmosTech/lama2/l2lsp/response"
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
	}
	return "", response.ErrUnexpectedURIScheme, errors.New("encountered unexpected URI scheme. Ex: 'file:///path/to/workspace/myapi.l2'")

}

func SuggestEnvironmentVariables(req request.JSONRPCRequest) response.JSONRPCResponse {
	searchQuery := getSearchQueryString(req)
	uri, errorCode, err := getRequestURI(req)
	if err != nil {
		return response.ErrorResp(req, errorCode, err.Error())
	}
	parentFolder := filepath.Dir(uri)
	res := l2envpackege.ProcessEnvironmentVariables(searchQuery, parentFolder)
	return response.CreateEnvironmentVariablesResp(req, res)
}
