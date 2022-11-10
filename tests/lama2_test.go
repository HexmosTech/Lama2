package tests

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/parser"
	"github.com/HexmosTech/lama2/utils"
)

func FileToString(filePath string) (string, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	str := string(b) // convert content to a 'string'
	return str, nil
}

func jsonFileToGabs(jsonPath string) (*gabs.Container, error) {
	jsonContent, e1 := FileToString(jsonPath)
	if e1 != nil {
		return nil, e1
	}

	j, e2 := gabs.ParseJSON([]byte(jsonContent))
	if e2 != nil {
		return nil, e2
	}
	return j, nil
}

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

func PerformParserMatch(text string) (*gabs.Container, error) {
	p := parser.NewLama2Parser()
	got, e := p.Parse(text)
	if e == nil {
		fmt.Println(got)
		fmt.Println("===")
	} else {
		// t.Errorf("Error not expected")
		fmt.Println(e)
	}
	return got, e
}

func TestValidFiles(t *testing.T) {
	matchFiles, _ := getDataFiles("../elfparser/ElfTestSuite", "y_*")
	// matchFiles, _ := getDataFiles("../elfparser/ElfTestSuite", "y_0009_varjson_basic.http")
	// matchFiles, _ := getDataFiles("../elfparser/ElfTestSuite", "y_0012_varjson_multipart.http")
	// matchFiles, _ := getDataFiles("../elfparser/ElfTestSuite", "y_0014_at_equal_ambiguity.http")
	// matchFiles, _ := getDataFiles("../elfparser/ElfTestSuite", "y_0015_number_vars.http")
	// matchFiles, _ := getDataFiles("../elfparser/ElfTestSuite", "y_0016_simple_array.http")
	for _, m := range matchFiles {
		b, err := os.ReadFile(m) // just pass the file name
		if err != nil {
			fmt.Print(err)
		}

		str := string(b) // convert content to a 'string'
		fmt.Println(str)
		myOp, _ := PerformParserMatch(str)

		fmt.Println(myOp)
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

func TestJsonParserExhaustive(t *testing.T) {
	rootPath := "../elfparser/JSONTestSuite/test_parsing"
	matchFiles, _ := getDataFiles(rootPath, "y_*")
	ignoreNames := []string{
		"y_string_accepted_surrogate_pair.json",
		"y_string_accepted_surrogate_pairs.json",
		"y_string_last_surrogates_1_and_2.json",
		"y_string_surrogates_U+1D11E_MUSICAL_SYMBOL_G_CLEF.json",
		"y_string_unescaped_char_delete.json",
		"y_string_with_del_character.json",
		"y_string_unicode_U+10FFFE_nonchar.json",
		"y_string_unicode_U+1FFFE_nonchar.json",
	}
	for _, m := range matchFiles {
		fmt.Println("### === === === === ===")
		if utils.ContainsStringPartial(ignoreNames, m) {
			continue
		}
		fmt.Println(m)
		jsonText, e := FileToString(m)
		fmt.Println(jsonText)
		if e != nil {
			fmt.Println("fileToString failed")
			return
		}

		gj, e2 := jsonFileToGabs(m)
		if e2 != nil {
			fmt.Println("jsonFileToGabs failed")
			return
		}
		preamble := "POST\nhttp://google.com\n"
		lamaText := preamble + jsonText
		fmt.Println(lamaText)
		jj, e3 := PerformParserMatch(lamaText)
		if e3 != nil {
			fmt.Println("performvalidMatch failed")
			return
		}
		fmt.Println("gj = ", gj)
		jp := jj.S("value", "details", "ip_data")
		fmt.Println("jj = ", jp)

		var v1, v2 interface{}
		json.Unmarshal([]byte(gj.String()), &v1)
		jpStr := jp.String()
		fmt.Println("### jpStr", jpStr)
		json.Unmarshal([]byte(jpStr), &v2)
		if reflect.DeepEqual(v1, v2) {
			fmt.Println("The two JSON structures are equal")
			fmt.Println(m)
		} else {
			// They are different
			fmt.Println("The two JSON structures are different")
			fmt.Println(m)
			break
		}
		fmt.Println("*** === === === === ===")
	}
}

func TestNegativeJsonParserExhaustive(t *testing.T) {
	rootPath := "../elfparser/JSONTestSuite/test_parsing"
	matchFiles, _ := getDataFiles(rootPath, "n_*")
	ignoreNames := []string{
		"n_array_extra_comma.json",
		"n_number_1_000.json",
		"n_number_minus_space_1.json",
		"n_array_number_and_comma",
		"n_object_single_quote.json",
		"n_object_with_trailing_garbage.json",
		"n_string_single_quote.json",
		"n_object_lone_continuation_byte_in_key_and_trailing_comma.json",
		"n_object_trailing_comma.json",
		"n_structure_100000_opening_arrays.json",
		"n_single_space.json",
		"n_structure_no_data.json",
		"n_structure_open_array_object.json",
		"n_structure_trailing_#.json",
		"n_structure_whitespace_formfeed.json",
	}
	for _, m := range matchFiles {
		fmt.Println("### === === === === ===")
		if utils.ContainsStringPartial(ignoreNames, m) {
			continue
		}
		fmt.Println(m)
		jsonText, e := FileToString(m)
		fmt.Println(jsonText)
		if e != nil {
			fmt.Println("fileToString failed")
			return
		}

		preamble := "POST\nhttp://google.com\n"
		lamaText := preamble + jsonText
		fmt.Println(lamaText)
		_, e3 := PerformParserMatch(lamaText)
		if e3 == nil {
			t.Fatalf("Expected parser to fail for %s", m)
		}
		fmt.Println("*** === === === === ===")
	}
}
