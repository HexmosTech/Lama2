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
	"fmt"
	"os"
	"reflect"

	"github.com/Jeffail/gabs/v2"
	"github.com/hexmos/lama2/utils"
)

type ParserMinimalType interface {
	Start() (*gabs.Container, error)
}

type Parser struct {
	Text          []rune
	Pos           int
	TotalLen      int
	cache         map[string][]string
	Pm            ParserMinimalType
	ruleMethodMap map[string]reflect.Value
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
	fmt.Println(p.ruleMethodMap)
}

func (p *Parser) Parse(text string) (*gabs.Container, error) {
	p.Text = []rune(text)
	p.Pos = -1
	p.TotalLen = len(text) - 1
	p.cache = make(map[string][]string)
	res, _ := p.Pm.Start()
	_, err := p.assertEnd()
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		os.Exit(1)
	}
	return res, nil
}

/*
func (p *Parser) eatWhitespace() {
	for p.Pos < p.TotalLen {
		spaces := []rune{' ', '\f', '\v', '\r', '\t', '\n'}
		if utils.ContainsRune(spaces, p.Text[p.Pos+1]) {
			p.Pos += 1
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
		p.Pos += 1
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
				"Expected end of string but got %s",
				[]string{string(p.Text[p.Pos+1])})
	} else {
		return true, nil
	}
}
