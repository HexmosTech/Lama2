// Package `cmdexec` provides a facility to execute
// l2 commands, stream output to stdout, while also
// providing ability to retrieve the command output as
// a string.
package cmdexec

import (
	"os"
	"runtime"
	"strings"

	"github.com/HexmosTech/httpie-go"
	"github.com/HexmosTech/lama2/utils"
)

// ExecCommand changes directory to the given `apiDir`
// and then executes the command specified in `cmdStr`
// During command execution, ExecCommand streams output
// to stdout.
// Once execution finishes, previous CWD is restored,
// and the command output is returned as a string
func ExecCommand(cmdSlice []string, stdinBody string, apiDir string) string {
	oldDir, _ := os.Getwd()
	utils.ChangeWorkingDir(apiDir)
	var retStr string

	if runtime.GOOS == "windows" {
		/*
			f, err := exec.Command("cmd", "/C", cmdStr).Output()
			if err != nil {
				panic(err)
			}
			retStr = string(f)
		*/
	} else {
		httpie.Lama2Entry(cmdSlice, strings.NewReader(stdinBody))
		/*
			c := exec.Command("bash", "-c", cmdStr)
			f, err := pty.Start(c)
			if err != nil {
				panic(err)
			}
			var buffer1 bytes.Buffer
			writer := io.MultiWriter(&buffer1, os.Stdout)
			io.Copy(writer, f)
			ret, _ := io.ReadAll(&buffer1)
			retStr = string(ret)
		*/
	}

	utils.ChangeWorkingDir(oldDir)
	return retStr
}
