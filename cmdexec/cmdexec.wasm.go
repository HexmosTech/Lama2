//go:build wasm

// Package `cmdexec` provides a facility to execute
// l2 commands, stream output to stdout, while also
// providing ability to retrieve the command output as
// a string.
package cmdexec

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/HexmosTech/httpie-go"
)

// ExecCommand changes directory to the given `apiDir`
// and then executes the command specified in `cmdStr`
// During command execution, ExecCommand streams output
// to stdout.
// Once execution finishes, previous CWD is restored,
// and the command output is returned as a string
func ExecCommand(cmdSlice []string, stdinBody string, apiDir string) (httpie.ExResponse, int64, error) {
	fmt.Println("WASM: Executing command")
	proxyURL := os.Getenv("PROXY_URL")
	proxyUserName := os.Getenv("PROXY_USERNAME")
	proxyUserPassword := os.Getenv("PROXY_PASSWORD")
	allowRedirects := true
	start := time.Now()
	resp, err := httpie.Lama2Entry(cmdSlice, strings.NewReader(stdinBody), proxyURL, proxyUserName, proxyUserPassword, allowRedirects)
	elapsed := time.Since(start)
	responseTime := elapsed.Milliseconds()
	if err != nil {
		fmt.Println("Got error while executing", err)
		return httpie.ExResponse{}, 0, errors.New("Error from API executor: " + err.Error())
	}
	fmt.Println("WASM: Command executed successfully")
	return resp, responseTime, nil
}
