package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/creack/pty"
	"github.com/dlclark/regexp2"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/json"
	"golang.org/x/text/language"
	"golang.org/x/text/search"
)

var LOG_LEVEL = "DEBUG"

func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

func findFirst(expr *regexp.Regexp, content string) (string, []int) {
	found := expr.FindAllString(content, -1)
	index := expr.FindStringIndex(content)

	if found == nil {
		return "", []int{-1, -1}
	}
	exactMatch := strings.TrimSpace(string(found[0]))
	return exactMatch, index
}

func regexp2FindAllString(re *regexp2.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m.String())
		m, _ = re.FindNextMatch(m)
	}
	return matches
}

func findEach2(expr *regexp2.Regexp, content string) []string {
	matches := make([]string, 0)
	found := regexp2FindAllString(expr, content)

	if found == nil {
		return append(matches, "")
	}

	for _, word := range found {
		matches = append(matches, strings.TrimSpace(string(word)))
	}
	return matches
}

func SubstringIndex(str string, substr string) (int, bool) {
	m := search.New(language.English, search.IgnoreCase)
	start, _ := m.IndexString(str, substr)
	if start == -1 {
		return start, false
	}
	return start, true
}

func findNamedMatch(expr *regexp2.Regexp, content string, groupName string) string {
	match, _ := expr.FindStringMatch(content)
	if match == nil {
		return ""
	}
	gp := match.GroupByName(groupName)
	return gp.Captures[0].String()
}

/*
 * Two requirements:
 * 1. Compile time setting of log level should be possible
 * 2. Runtime overriding of the log level should be possible
 * For (1), pass -ldflags="-X main.LOG_LEVEL=INFO" for example,
 * to hide DEBUG messages in production.
 * For (2), command parsing can be used to support "-v" or
 * "--verbose" options
 */
func configureZeroLog(level string) {
	log_level_map := make(map[string]zerolog.Level)
	log_level_map["DEBUG"] = zerolog.DebugLevel
	log_level_map["INFO"] = zerolog.InfoLevel
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(log_level_map[level])
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

}

func validateCmdArgs() {
	if len(os.Args) < 2 {
		log.Error().Str("Type", "Preprocess").Msg(color.RedString("Argument missing"))
		log.Error().Msg(color.GreenString("Usage: elf <request_file>.http"))
		os.Exit(0)
	}
}

func getHttpFileAsString(path string) string {
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		log.Fatal().Str("Type", "Preprocess").Msg(fmt.Sprint("Couldn't read: ", path))
	}
	return string(b)
}

func getFilePathComponents(name string) (string, string, string) {
	fullpath, _ := filepath.Abs(name)
	dir, fname := path.Split(fullpath)
	return fullpath, dir, fname
}

func changeWorkingDir(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		log.Fatal().
			Str("Type", "Preprocess").
			Str("dir", dir).
			Msg(fmt.Sprint("Moving into dir failed"))
	}
}

func loadElfEnv() {
	err := godotenv.Load("elf.env")
	if err != nil {
		log.Info().Str("Type", "Preprocess").Msg("Error loading elf.env file")
	}
}

func getSpecPieces(content string) (string, string, string, string, []string, []string, []string) {
	r_json_obj := regexp.MustCompile(`(?smi)^([{\[].*[}\]])$`)
	json_obj, json_idx := findFirst(r_json_obj, content)

	r_httpv := regexp.MustCompile(`(?smi)^(GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH)\s*$`)
	r_url := regexp.MustCompile(`(?smi)^https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.?[a-zA-Z0-9()]{1,6}\b?([-a-zA-Z0-9()@:%_\+.~#?&\//=]*).*?$`)
	// r_varjsonorheader := regexp2.MustCompile(`(?sm)^(?![http|#|{])[\S]+[:=](?:(?!\S+[:=])\S+)+\s*$`, 0)
	// Regex test cases: https://regex101.com/r/KDrfDt/3
	r_linetype := regexp2.MustCompile(`(?sm)^(?![http|#|{])^("([^"]*)"|'([^']*)'|(.*?))(?P<type>[:@=])("([^"]*)"|'([^']*)'|(.*?))\s*$`, regexp2.RE2)
	r_multipart := regexp.MustCompile(`(?smi)^MULTIPART\s*$`)
	// r_filefield := regexp2.MustCompile(`(?sm)^(?![http|#|{])[\S]+[@](?:(?!\S+[])\S+)+\s*$`, 0)
	// r_headers := regexp2.MustCompile(`(?!http)(?sm)^[\S]+[:=](?:(?!\S+[:=])\S+)+\s*$`, 0)

	httpv, _ := findFirst(r_httpv, content)
	url, _ := findFirst(r_url, content)
	// varjsonorheader := findEach2(r_varjsonorheader, content)
	multipart, _ := findFirst(r_multipart, content)

	varjson := make([]string, 0)
	headers := make([]string, 0)
	filefields := make([]string, 0)

	log.Debug().
		Str("Type", "Processing").
		Ints("JsonIdx", json_idx).
		Str("JsonObj", json_obj).
		Msg("Removing JSON from content")

	// remove multiline json object
	linecontent := content
	if json_idx[0] != -1 {
		linecontent = content[:json_idx[0]] + content[json_idx[1]:]
	}
	for _, line := range strings.Split(linecontent, "\n") {
		line = strings.TrimSpace(line)
		t := findNamedMatch(r_linetype, line, "type")
		if t == ":" {
			headers = append(headers, line)
		} else if t == "=" {
			varjson = append(varjson, line)
		} else if t == "@" {
			filefields = append(filefields, line)
		}
	}

	log.Debug().
		Str("Type", "Postprocessing").
		Str("httpv", httpv).
		Str("url", url).
		Strs("varjson", varjson).
		Strs("headers", headers).
		Strs("filefields", filefields).
		Msg("Extracted fields")

	return httpv, url, json_obj, multipart, varjson, headers, filefields
}

func constructCommand(httpv string, url string, json_obj string, multipart string,
	varjson []string, headers []string, filefields []string) string {
	command := make([]string, 0)
	if len(multipart) == 0 && json_obj != "" {
		m := minify.New()
		m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
		json_obj, err := m.String("*/json", json_obj)
		if err != nil {
			panic(err)
		}
		command = append(command, "echo '")
		command = append(command, json_obj)
		command = append(command, "' | ")

	}

	command = append(command, "http")
	if len(multipart) > 0 {
		command = append(command, "--multipart")
	}
	command = append(command, httpv)
	command = append(command, url)
	if varjson != nil {
		command = append(command, varjson...)
	}
	if headers != nil {
		command = append(command, headers...)
	}
	if filefields != nil {
		command = append(command, filefields...)
	}

	command_str := strings.Join(command[:], " ")
	log.Info().Str("Type", "Postprocessing").Msg(command_str)
	fmt.Println()
	fmt.Println()
	return command_str
}

func executeCommand(command_str string) string {
	// r, w := io.Pipe()
	c := exec.Command("bash", "-c", os.ExpandEnv(command_str))
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
	ret_str := string(ret)
	return ret_str
}

func processHttpFile(input_f string) string {
	content := getHttpFileAsString(input_f)
	_, dir, _ := getFilePathComponents(input_f)
	changeWorkingDir(dir)
	loadElfEnv()
	content = os.ExpandEnv(content)

	httpv, url, json_obj, multipart, varjson, headers, filefields := getSpecPieces(content)

	command_str := constructCommand(httpv, url, json_obj, multipart, varjson, headers, filefields)
	return executeCommand(command_str)
}

func main() {
	configureZeroLog(LOG_LEVEL)
	validateCmdArgs()
	input_f := os.Args[1]
	processHttpFile(input_f)
}
