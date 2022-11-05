package tests

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/hexmos/lama2/parser"
)

func getDataFiles(relativeAppend string, globPattern string) ([]string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	res := filepath.Join(pwd, relativeAppend, globPattern)
	matches, err := filepath.Glob(res)
	fmt.Println(matches)
	if err == nil {
		sort.Strings(matches)
		return matches, nil
	} else {
		fmt.Println("Couldn't get data file matches")
		return nil, errors.New("Couldn't get data file matches")
	}
}

func PerformValidMatch(text string) {
	p := parser.NewLama2Parser()
	got, e := p.Parse(text)
	if e == nil {
		fmt.Println(got)
		fmt.Println("===")
	} else {
		// t.Errorf("Error not expected")
		fmt.Println(e)
	}
}

func TestValidFiles(t *testing.T) {
	// matchFiles, _ := getDataFiles("../elfparser/ElfTestSuite", "y_*")
	// matchFiles, _ := getDataFiles("../elfparser/ElfTestSuite", "y_0009_varjson_basic.http")
	matchFiles, _ := getDataFiles("../elfparser/ElfTestSuite", "y_0012_varjson_multipart.http")
	// matchFiles, _ := getDataFiles("../elfparser/ElfTestSuite", "y_0014_at_equal_ambiguity.http")
	for _, m := range matchFiles {
		b, err := os.ReadFile(m) // just pass the file name
		if err != nil {
			fmt.Print(err)
		}

		str := string(b) // convert content to a 'string'
		fmt.Println(str)
		PerformValidMatch(str)
	}
}

func TestInvalidFiles(t *testing.T) {
	matchFiles, _ := getDataFiles("../elfparser/ElfTestSuite", "n_*")
	for _, m := range matchFiles {
		b, err := os.ReadFile(m) // just pass the file name
		if err != nil {
			fmt.Print(err)
		}

		str := string(b) // convert content to a 'string'
		fmt.Println(str)
	}
}
