package testutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

type EnvData struct {
	Src string `json:"src"`
	Val string `json:"val"`
}

func checkLocalL2BinaryExists(l2BinPath string) error {
	if _, err := os.Stat(l2BinPath); os.IsNotExist(err) {
		return fmt.Errorf("l2 binary not found in the build folder %s, please change the path", l2BinPath)
	}
	return nil
}

func GetLocalL2BinaryPath() (string, error) {
	l2BinPath := "../build/l2"
	err := checkLocalL2BinaryExists(l2BinPath)
	if err != nil {
		return "", err
	}
	return l2BinPath, nil
}

func runCommand(binPath string, cmdArgs ...string) (string, error) {
	cmd := exec.Command(binPath, cmdArgs...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Error().Str("Error", stderr.String()).Msg("Error running command")
		return "", fmt.Errorf("error running command: %v", err)
	}

	output := stdout.String()
	log.Debug().Str("Test env_command", output).Msg("Output from command")
	return output, nil
}

func RunL2CommandAndGetOutput(cmdArgs ...string) (string, error) {
	l2BinPath, err := GetLocalL2BinaryPath()
	if err != nil {
		return "", err
	}

	return runCommand(l2BinPath, cmdArgs...)
}

func RunL2CommandAndParseJSON(cmdArgs ...string) (map[string]EnvData, error) {
	l2BinPath, err := GetLocalL2BinaryPath()
	if err != nil {
		return nil, err
	}

	output, err := runCommand(l2BinPath, cmdArgs...)
	if err != nil {
		return nil, err
	}

	envMap := make(map[string]EnvData)
	err = json.Unmarshal([]byte(output), &envMap)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON env: %v\nOutput:\n%s", err, output)
	}

	return envMap, nil
}

func GetAbsolutePath(relPath string) (string, error) {
	// GetAbsolutePath takes a relative path and returns its absolute path.
	absPath, err := filepath.Abs(relPath)
	if err != nil {
		return "", err
	}
	return absPath, nil
}
