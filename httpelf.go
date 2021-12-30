package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/creack/pty"
	"github.com/dlclark/regexp2"
	"github.com/joho/godotenv"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/json"
	"golang.org/x/text/language"
	"golang.org/x/text/search"
)

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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("\u001B[31mArgument missing\u001B[0m")
		fmt.Println("Usage: elf <request_file>.http")
		os.Exit(0)
	}
	b, err := ioutil.ReadFile(os.Args[1]) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	content := string(b) // convert content to a 'string'

	fullpath, _ := filepath.Abs(os.Args[1])
	dir, _ := path.Split(fullpath)
	err = os.Chdir(dir)
	if err != nil {
		fmt.Print(err)
	}

	err = godotenv.Load("elf.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	photo := os.Getenv("PHOTO")
	fmt.Println(photo)
	content = os.ExpandEnv(content)

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

	fmt.Println("===")
	fmt.Println("JSON IDX = ", json_idx)
	fmt.Println("JSON OBJ = ", json_obj)
	fmt.Println("===")

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

	/*
		for _, word := range varjsonorheader {
			colonIdx, isheader := SubstringIndex(word, ":")
			equalIdx, isequal := SubstringIndex(word, "=")
			if isheader && isequal {
				if colonIdx > equalIdx {
					// equal comes first; color comes next => varjson
					varjson = append(varjson, word)
				} else {
					headers = append(headers, word)
				}
			} else {
				if isheader {
					headers = append(headers, word)
				} else {
					varjson = append(varjson, word)
				}
			}
		}
	*/

	fmt.Println("httpv", httpv)
	fmt.Println("url", url)
	fmt.Println("json_obj", json_obj)
	fmt.Println("varjson", varjson)
	fmt.Println("headers", headers)
	fmt.Println("multipart", multipart)
	fmt.Println("filefields", filefields)

	command := make([]string, 0)
	if len(multipart) == 0 && json_obj != "" {
		m := minify.New()
		m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
		json_obj, err = m.String("*/json", json_obj)
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
	fmt.Println()
	fmt.Println(command_str)

	// fmt.Println("EXP", exp)
	c := exec.Command("bash", "-c", os.ExpandEnv(command_str))
	f, err := pty.Start(c)
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, f)

	Use(httpv, url, json_obj, varjson, headers, command)
}
