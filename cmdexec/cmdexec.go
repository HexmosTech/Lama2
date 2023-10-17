// Package `cmdexec` provides a facility to execute
// l2 commands, stream output to stdout, while also
// providing ability to retrieve the command output as
// a string.
package cmdexec

import (
	"errors"
	// "os"
	"strings"

	"github.com/HexmosTech/httpie-go"
	// "github.com/HexmosTech/lama2/utils"
	"github.com/rs/zerolog/log"
)

// ExecCommand changes directory to the given `apiDir`
// and then executes the command specified in `cmdStr`
// During command execution, ExecCommand streams output
// to stdout.
// Once execution finishes, previous CWD is restored,
// and the command output is returned as a string
func ExecCommand(cmdSlice []string, stdinBody string) (httpie.ExResponse, error) {
	// oldDir, _ := os.Getwd()
	// utils.ChangeWorkingDir(apiDir)
	resp, err := httpie.Lama2Entry(cmdSlice, strings.NewReader(stdinBody))
	if err != nil {
		log.Fatal().Str("Error from the API executor", err.Error()).Msg("")
		return httpie.ExResponse{}, errors.New("Error from API executor: " + err.Error())
	}
	log.Debug().Str("Response body from API executor", resp.Body).Msg("")
	// utils.ChangeWorkingDir(oldDir)
	return resp, nil
}
