//go:build wasm

package outputmanager

import (
	"bytes"
	"os"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/httpie-go"
)

var LogBuff bytes.Buffer

func wrapError(requestError string) *gabs.Container {
	temp := gabs.New()
	temp.Set(requestError, "errors")
	return temp
}

func ResponseToJSON(resp httpie.ExResponse) (*gabs.Container, error) {
	body := string(resp.Body)

	var headerMapStr string
	for k, v := range resp.Headers {
		headerMapStr += k + ": " + v + "\n"
	}

	temp := gabs.New()
	temp.Set(headerMapStr, "headers")
	temp.Set(body, "body")
	temp.Set(LogBuff.String(), "logs")

	return temp, nil
}

func WriteJSONOutput(resp httpie.ExResponse, targetPath string) {
	temp, _ := ResponseToJSON(resp)
	err := os.WriteFile(targetPath, []byte(temp.String()), 0o644)
	if err != nil {
	}
}
