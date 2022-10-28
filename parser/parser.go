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
	cache *gabs.Container
	pm    ParserMinimalType
}

func (p *Parser) Start() *gabs.Container {
	temp := gabs.New()
	return temp
}

func (p *Parser) Init() {
	p.cache = gabs.New()
}

func (p *Parser) Parse(text string) string {
	p.text = []rune(text)
	p.pos = -1
	p.tlen = len(text) - 1
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
}

func (p *Parser) Char() (rune, error) {
	if p.pos >= p.tlen {
		return rune(0),
			utils.NewParseError(
				p.pos,
				"Expected %s but got end of string",
				[]string{"character"})
	}
	next_char := p.text[p.pos+1]
	p.pos += 1
	return next_char, nil
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
