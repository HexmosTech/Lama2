// Package `cmdexec` provides a facility to execute
// l2 commands, stream output to stdout, while also
// providing ability to retrieve the command output as
// a string.
package cmdexec

import (
	"bytes"
	"io"
	"os"
	"os/exec"

	"github.com/HexmosTech/lama2/utils"
	"github.com/creack/pty"
)

// ExecCommand changes directory to the given `apiDir`
// and then executes the command specified in `cmdStr`
// During command execution, ExecCommand streams output
// to standard output.
// Once execution finishes, previous CWD is restored,
// and the command output is returned as a string
func ExecCommand(cmdStr string, apiDir string) string {
	oldDir, _ := os.Getwd()
	utils.ChangeWorkingDir(apiDir)
	c := exec.Command("bash", "-c", cmdStr)
	f, err := pty.Start(c)
	if err != nil {
		panic(err)
	}
	var buffer1 bytes.Buffer
	writer := io.MultiWriter(&buffer1, os.Stdout)
	io.Copy(writer, f)
	ret, _ := io.ReadAll(&buffer1)
	retStr := string(ret)
	utils.ChangeWorkingDir(oldDir)
	return retStr
}
