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
	contoller "github.com/HexmosTech/lama2/controller"
	"github.com/HexmosTech/lama2/parser"
	"github.com/HexmosTech/lama2/utils"
	"github.com/rs/zerolog/log"
)

func FileToString(filePath string) (string, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		log.Debug().Str("Error", err.Error()).Msg("")
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
		log.Debug().Str("Error", err.Error()).Msg("")
		os.Exit(1)
	}
	res := filepath.Join(pwd, relativeAppend, globPattern)
	matches, err := filepath.Glob(res)
	log.Debug().Strs("Matches", matches).Msg("")
	if err == nil {
		sort.Strings(matches)
		return matches, nil
	}
	log.Debug().Msg("Couldn't get data file matches")
	return nil, errors.New("Couldn't get data file matches")
}

func PerformParserMatch(text string) (*gabs.Container, error) {
	p := parser.NewLama2Parser()
	got, e := p.Parse(text)
	if e == nil {
		log.Debug().Str("Got", got.String()).Msg("")
	}
	return got, e
}

func TestValidFiles(_ *testing.T) {
	matchFiles, _ := getDataFiles("../elfparser/ElfTestSuite", "y_*")
	// matchFiles, _ := getDataFiles("../elfparser/ElfTestSuite", "y_0009_varjson_basic.l2")
	// matchFiles, _ := getDataFiles("../elfparser/ElfTestSuite", "y_0012_varjson_multipart.l2")
	// matchFiles, _ := getDataFiles("../elfparser/ElfTestSuite", "y_0014_at_equal_ambiguity.l2")
	// matchFiles, _ := getDataFiles("../elfparser/ElfTestSuite", "y_0015_number_vars.l2")
	// matchFiles, _ := getDataFiles("../elfparser/ElfTestSuite", "y_0016_simple_array.l2")
	// matchFiles, _ := getDataFiles("../elfparser/ElfTestSuite", "y_0017_processor.l2")
	for _, m := range matchFiles {
		b, err := os.ReadFile(m) // just pass the file name
		if err != nil {
			log.Fatal().Str("Error:", err.Error()).Msg("")
		}

		str := string(b) // convert content to a 'string'
		myOp, _ := PerformParserMatch(str)
		log.Debug().Str("Execution result myOp", myOp.String()).Msg("")
	}
}

func TestInvalidFiles(_ *testing.T) {
	matchFiles, _ := getDataFiles("../elfparser/ElfTestSuite", "n_*")
	for _, m := range matchFiles {
		b, err := os.ReadFile(m) // just pass the file name
		if err != nil {
			log.Fatal().Str("Error:", err.Error()).Msg("")
		}

		str := string(b) // convert content to a 'string'
		log.Debug().Str("str", str).Msg("")
	}
}

func TestJsonParserExhaustive(t *testing.T) {
	rootPath := "../elfparser/JSONTestSuite/test_parsing"
	matchFiles, _ := getDataFiles(rootPath, "y_*")
	ignoreNames := []string{
		// Enable these files to ignore failing test cases
		// for some complicated Unicode cases:

		// "y_string_accepted_surrogate_pair.json",
		// "y_string_accepted_surrogate_pairs.json",
		// "y_string_last_surrogates_1_and_2.json",
		// "y_string_surrogates_U+1D11E_MUSICAL_SYMBOL_G_CLEF.json",
		"y_string_unescaped_char_delete.json",
		"y_string_with_del_character.json",
		// "y_string_unicode_U+10FFFE_nonchar.json",
		// "y_string_unicode_U+1FFFE_nonchar.json",
	}
	for _, m := range matchFiles {
		// fmt.Println("### === === === === ===")
		if utils.ContainsStringPartial(ignoreNames, m) {
			continue
		}
		jsonText, e := FileToString(m)
		if e != nil {
			log.Fatal().Str("Error:", e.Error()).Msg("")
		}

		gj, e2 := jsonFileToGabs(m)
		if e2 != nil {
			return
		}
		preamble := "POST\nhttp://google.com\n"
		lamaText := preamble + jsonText
		jj, e3 := PerformParserMatch(lamaText)
		if e3 != nil {
			return
		}

		blocks := contoller.GetParsedAPIBlocks(jj)
		var jp *gabs.Container
		for _, block := range blocks {
			jp = block
		}
		jp = jp.S("details", "ip_data")

		var v1, v2 interface{}
		json.Unmarshal([]byte(gj.String()), &v1)
		jpStr := jp.String()
		json.Unmarshal([]byte(jpStr), &v2)
		if reflect.DeepEqual(v1, v2) {
			fmt.Println("The two JSON structures are equal")
			fmt.Println(m)
		} else {
			t.Errorf("The two JSON structures are different")
			fmt.Println(m)
			break
		}
		// fmt.Println("*** === === === === ===")
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
		log.Trace().Msg("### === === === === ===")
		if utils.ContainsStringPartial(ignoreNames, m) {
			continue
		}
		log.Trace().Str("m", m).Msg("")
		jsonText, e := FileToString(m)
		log.Trace().Str("JSONText", jsonText).Msg("")
		if e != nil {
			log.Trace().Msg("fileToString failed")
			return
		}

		preamble := "POST\nhttp://google.com\n"
		lamaText := preamble + jsonText
		log.Trace().Str("LamaText", lamaText).Msg("")
		_, e3 := PerformParserMatch(lamaText)
		if e3 == nil {
			t.Fatalf("Expected parser to fail for %s", m)
		}
		log.Trace().Msg("*** === === === === ===")
	}
}
