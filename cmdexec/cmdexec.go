package cmdexec

import (
	"bytes"
	"io"
	"os"
	"os/exec"

	"github.com/HexmosTech/lama2/utils"
	"github.com/creack/pty"
)

func ExecCommand(cmdStr string, apiDir string) string {
	oldDir, _ := os.Getwd()
	utils.ChangeWorkingDir(apiDir)
	c := exec.Command("bash", "-c", cmdStr)
	f, err := pty.Start(c)
	if err != nil {
		panic(err)
	}

	// Defining two buffers
	var buffer1 bytes.Buffer

	// Calling MultiWriter method with its parameters
	writer := io.MultiWriter(&buffer1, os.Stdout)
	io.Copy(writer, f)
	ret, _ := io.ReadAll(&buffer1)
	retStr := string(ret)
	utils.ChangeWorkingDir(oldDir)
	return retStr
}
