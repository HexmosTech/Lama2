// Package `cmdexec` provides a facility to execute
// l2 commands, stream output to stdout, while also
// providing ability to retrieve the command output as
// a string.
package cmdexec

import (
	"errors"
	"fmt"

	// "os"
	"strings"

	"github.com/HexmosTech/httpie-go"
	// "github.com/HexmosTech/lama2/utils"
	// "github.com/rs/zerolog/log"
	// "github.com/HexmosTech/httpie-go"
)

// ExecCommand changes directory to the given `apiDir`
// and then executes the command specified in `cmdStr`
// During command execution, ExecCommand streams output
// to stdout.
// Once execution finishes, previous CWD is restored,
// and the command output is returned as a string
func ExecCommand(cmdSlice []string, stdinBody string) (httpie.ExResponse, error) {
	fmt.Println("Inside cmd exec code")
	fmt.Println("Inside cmd exec code httpie Integer", httpie.Int(24601))
	proxyURL := "https://proxyserver.hexmos.com/"	
	proxyUserName := "proxyServer"
	proxyUserPassword := "proxy22523146server"
	allowRedirects := true
	resp, err := httpie.Lama2Entry(cmdSlice, strings.NewReader(stdinBody), proxyURL, proxyUserName, proxyUserPassword, allowRedirects)
	// resp, err := httpie.Lama2Entry(cmdSlice, strings.NewReader(stdinBody), "", "", "", false)
	fmt.Println("error:", err)
	fmt.Println("Data Response:", resp)
	if err != nil {
		// log.Fatal().Str("Error from the API executor", err.Error()).Msg("")
		fmt.Println("Got error while executing", err)
		return httpie.ExResponse{}, errors.New("Error from API executor: " + err.Error())
	}
	// log.Debug().Str("Response body from API executor", resp.Body).Msg("")
	// utils.ChangeWorkingDir(oldDir)
	// fmt.Println("Reponse fro httpie", resp)
	return resp, nil
}
