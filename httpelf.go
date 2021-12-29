package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/creack/pty"
	"github.com/dlclark/regexp2"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/json"
)

func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

func findFirst(expr *regexp.Regexp, content string) string {
	found := expr.FindAllString(content, -1)

	if found == nil {
		return ""
	}

	exactMatch := strings.TrimSpace(string(found[0]))
	return exactMatch
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

func main() {
	b, err := ioutil.ReadFile(os.Args[1]) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	content := string(b) // convert content to a 'string'

	r_httpv := regexp.MustCompile(`(?sm)^(GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH)\s*$`)
	r_url := regexp.MustCompile(`(?sm)^https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.?[a-zA-Z0-9()]{1,6}\b?([-a-zA-Z0-9()@:%_\+.~#?&\//=]*).*?$`)
	r_json_obj := regexp.MustCompile(`(?sm)\s([{\[].*[}\]])$`)
	r_varjson := regexp2.MustCompile(`(?!http)(?sm)^[\S]+=(?:(?!\S+=)\S+)+\s*$`, 0)
	r_headers := regexp2.MustCompile(`(?!http)(?sm)^[\S]+:(?:(?!\S+:)\S+)+\s*$`, 0)

	httpv := findFirst(r_httpv, content)
	url := findFirst(r_url, content)
	json_obj := findFirst(r_json_obj, content)
	varjson := findEach2(r_varjson, content)
	headers := findEach2(r_headers, content)

	fmt.Print("URL", url)

	command := make([]string, 0)
	if json_obj != "" {
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
	command = append(command, httpv)
	command = append(command, url)
	if varjson != nil {
		command = append(command, varjson...)
	}
	if headers != nil {
		command = append(command, headers...)
	}

	command_str := strings.Join(command[:], " ")
	fmt.Println(command_str)

	c := exec.Command("bash", "-c", command_str)
	f, err := pty.Start(c)
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, f)

	Use(httpv, url, json_obj, varjson, headers, command)
}
