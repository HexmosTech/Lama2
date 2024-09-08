//go:build cli

// Package `cmdexec` provides a facility to execute
// l2 commands, stream output to stdout, while also
// providing ability to retrieve the command output as
// a string.
package cmdexec

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/HexmosTech/httpie-go"
	"github.com/HexmosTech/lama2/utils"
)

// ExecCommand changes directory to the given `apiDir`
// and then executes the command specified in `cmdStr`
// During command execution, ExecCommand streams output
// to stdout.
// Once execution finishes, previous CWD is restored,
// and the command output is returned as a string
func ExecCommand(cmdSlice []string, stdinBody string, apiDir string) (httpie.ExResponse, int64, error) {
	oldDir, _ := os.Getwd()
	utils.ChangeWorkingDir(apiDir)
	start := time.Now()
	resp, err := httpie.Lama2Entry(cmdSlice, strings.NewReader(stdinBody))
	if err != nil {
		return httpie.ExResponse{}, 0, errors.New("Error from API executor: " + err.Error())
	}
	elapsed := time.Since(start)
	utils.ChangeWorkingDir(oldDir)
	responseTime := elapsed.Milliseconds()

	return resp, responseTime, nil
}
