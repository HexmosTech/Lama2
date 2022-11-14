/*
The `parser` package provides primitives that
help with writing recursive descent parsers.
This version is a golang port of the original
Python implementation from https://tinyurl.com/rdescent

The `Parser` struct is supposed to be extended
to support parsing a new language. Take a look
at `lama2parser.go` for an example. Essentially
the actual parsing begins from the `Start()` method.
*/
package parser

import (
	"reflect"
	"strings"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/utils"
)

type MinimalParser interface {
	Start() (*gabs.Container, error)
}

type Parser struct {
	Text          []rune
	Pos           int
	TotalLen      int
	cache         map[string][]string
	Pm            MinimalParser
	ruleMethodMap map[string]reflect.Value
	LineNum       int
}

func (p *Parser) Start() *gabs.Container {
	temp := gabs.New()
	return temp
}

func (p *Parser) Init() {
	p.ruleMethodMap = make(map[string]reflect.Value, 0)
	pVal := reflect.ValueOf(p)
	pmVal := reflect.Indirect(pVal).FieldByName("Pm").Elem()
	tLamaVal := reflect.TypeOf(p.Pm)
	for i := 0; i < pmVal.NumMethod(); i++ {
		m := tLamaVal.Method(i)
		methodVal := pmVal.Method(i)
		p.ruleMethodMap[m.Name] = methodVal
	}
}

func (p *Parser) Parse(text string) (*gabs.Container, error) {
	text = strings.ReplaceAll(text, "\r\n", "\n")
	p.Text = []rune(text)
	p.Pos = -1
	p.TotalLen = len(p.Text) - 1
	p.cache = make(map[string][]string)
	p.LineNum = 0
	res, _ := p.Pm.Start()
	_, err := p.assertEnd()
	if err != nil {
		return nil, err
	}
	return res, nil
}

/*
func (p *Parser) eatWhitespace() {
	for p.Pos < p.TotalLen {
		spaces := []rune{' ', '\f', '\v', '\r', '\t', '\n'}
		if utils.ContainsRune(spaces, p.Text[p.Pos+1]) {
			p.Pos++
		} else {
			break
		}
	}
}
*/

func (p *Parser) eatWhitespace() {
	spaces := []rune{' ', '\f', '\v', '\r', '\t', '\n'}
	isProcessingComment := false
	for p.Pos < p.TotalLen {
		c := p.Text[p.Pos+1]
		if c == '\n' {
			p.LineNum += 1
		}
		if isProcessingComment {
			if c == '\n' {
				isProcessingComment = false
			}
		} else {
			if c == '#' {
				isProcessingComment = true
			} else if !utils.ContainsRune(spaces, c) {
				break
			}
		}
		p.Pos++
	}
}

func (p *Parser) SetText(text string) {
	p.Text = []rune(text)
	p.Pos = -1
	p.TotalLen = len(p.Text) - 1
	p.cache = make(map[string][]string)
}

func (p *Parser) assertEnd() (bool, error) {
	if p.Pos < p.TotalLen {
		return false,
			utils.NewParseError(
				p.Pos+1,
				p.LineNum+1,
				"Expected end of string but got %s",
				[]string{string(p.Text[p.Pos+1])})
	}
	return true, nil
}
