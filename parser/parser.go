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
	text          []rune
	pos           int
	tlen          int
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

func (p *Parser) Parse(text string) string {
	p.text = []rune(text)
	p.pos = -1
	p.tlen = len(text) - 1
	p.cache = make(map[string][]string)
	p.Pm.Start()
	_, err := p.assertEnd()
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		os.Exit(1)
	}
	return ""
}

func (p *Parser) eatWhitespace() {
	for p.pos < p.tlen {
		spaces := []rune{' ', '\f', '\v', '\r', '\t', '\n'}
		if utils.ContainsRune(spaces, p.text[p.pos+1]) {
			p.pos += 1
		} else {
			break
		}
	}
}

func (p *Parser) SetText(text string) {
	p.text = []rune(text)
	p.pos = -1
	p.tlen = len(p.text) - 1
	p.cache = make(map[string][]string)
}

func (p *Parser) assertEnd() (bool, error) {
	if p.pos < p.tlen {
		return false,
			utils.NewParseError(
				p.pos+1,
				"Expected end of string but got %s",
				[]string{string(p.text[p.pos+1])})
	} else {
		return true, nil
	}
}
