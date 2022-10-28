package parser

import (
	"fmt"
	"os"

	"github.com/Jeffail/gabs/v2"
	"github.com/hexmos/lama2/utils"
)

type ParserMinimalType interface {
	Start() *gabs.Container
}

type Parser struct {
	text  []rune
	pos   int
	tlen  int
	cache map[string][]string
	pm    ParserMinimalType
}

func (p *Parser) Start() *gabs.Container {
	temp := gabs.New()
	return temp
}

func (p *Parser) Init() {
}

func (p *Parser) Parse(text string) string {
	p.text = []rune(text)
	p.pos = -1
	p.tlen = len(text) - 1
	p.cache = make(map[string][]string)
	p.pm.Start()
	_, err := p.assertEnd()
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		os.Exit(1)
	}
	return "blah2"
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
