// Package `utils` provides useful functions for
// simplifying various programming tasks
package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/HexmosTech/gabs/v2"
	"github.com/rs/zerolog/log"
)

// SetJSON is a helper function to work with the `gabs` library, which
// in turn is an API on top of the standard JSON library
// The function helps us create `parentObj.key = childObj` through
// using the `Merge` primitive available in `gabs`
func SetJSON(parentObj *gabs.Container, childObj *gabs.Container, key string) *gabs.Container {
	temp := gabs.New()
	temp.Set(childObj, key)
	parentObj.Merge(temp)
	return parentObj
}

// PrettyPrint takes in a generic interface{}
// objects and uses standard JSON capabilities
// to try to print with indentation
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

// ContainsRune searches for rune `e` in
// a slice of runes `s`; returns a boolean
func ContainsRune(s []rune, e rune) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ContainsString searches for string `e` in
// a slice of strings `s`; returns a boolean
func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ContainsStringPartial substring-searches for string `e` in
// a slice of strings `s`; returns a boolean
func ContainsStringPartial(s []string, e string) bool {
	for _, a := range s {
		if strings.Contains(e, a) {
			return true
		}
	}
	return false
}

// UnicodeCategory returns the Unicode Character Category of the given rune.
func UnicodeCategory(r rune) string {
	for name, table := range unicode.Categories {
		if len(name) == 2 && unicode.Is(table, r) {
			return name
		}
	}
	return "Cn"
}

// GetFilePathComponent returns absolute path, directory,
// and filename given a filepath
func GetFilePathComponents(name string) (string, string, string) {
	fullpath, _ := filepath.Abs(name)
	dir := filepath.Dir(fullpath)
	fname := filepath.Base(fullpath)
	return fullpath, dir, fname
}

// ChangeWorkingDirectory tries to set the CWD; on failure
// it exits with a log error message
func ChangeWorkingDir(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		log.Fatal().
			Str("Type", "Preprocess").
			Str("dir", dir).
			Msg(fmt.Sprint("Moving into dir failed"))
	}
}

func downloadFile(filepath string, url string) (err error) {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// UpdateSelf downloads the installation script from
// the official repository, and executes it to update
// the l2 binary to the latest version
func UpdateSelf() {
	fmt.Println("Trying to update lama2")
	downloadFile("/tmp/lama2install.sh", "https://raw.githubusercontent.com/HexmosTech/Lama2/main/install.sh")
	cmd := exec.Command("sh", "-c", "chmod +x /tmp/lama2install.sh")
	cmd.Run()
	cmd = exec.Command("bash", "-c", "/tmp/lama2install.sh")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
