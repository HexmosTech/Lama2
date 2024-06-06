//go:build wasm

// Package `cmdexec` provides a facility to execute
// l2 commands, stream output to stdout, while also
// providing ability to retrieve the command output as
// a string.
package cmdexec

import (
	"errors"
	"fmt"

	"strings"

	"github.com/HexmosTech/httpie-go"
)

// ExecCommand changes directory to the given `apiDir`
// and then executes the command specified in `cmdStr`
// During command execution, ExecCommand streams output
// to stdout.
// Once execution finishes, previous CWD is restored,
// and the command output is returned as a string
func ExecCommand(cmdSlice []string, stdinBody string, apiDir string) (httpie.ExResponse, error) {
	proxyURL := "https://proxyserver.hexmos.com/"
	proxyUserName := "proxyServer"
	proxyUserPassword := "proxy22523146server"
	allowRedirects := true
	resp, err := httpie.Lama2Entry(cmdSlice, strings.NewReader(stdinBody), proxyURL, proxyUserName, proxyUserPassword, allowRedirects)
	if err != nil {
		fmt.Println("Got error while executing", err)
		return httpie.ExResponse{}, errors.New("Error from API executor: " + err.Error())
	}
	return resp, nil
}
