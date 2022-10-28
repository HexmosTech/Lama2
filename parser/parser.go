package parser

import (
	"errors"
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

func (p *Parser) CharClass(charClass string) (rune, error) {
	if p.pos >= p.tlen {
		return rune(0),
			utils.NewParseError(
				p.pos,
				"Expected %s but got end of string",
				[]string{"character"})
	}
	nextChar := p.text[p.pos+1]
	charRangeList, e := p.SplitCharRanges(charClass)
	if e != nil {
		fmt.Errorf("%s", e)
		os.Exit(1)
	}

	for _, charRange := range charRangeList {
		runeCharRange := []rune(charRange)
		if len(runeCharRange) == 1 {
			if nextChar == runeCharRange[0] {
				p.pos += 1
				return nextChar, nil
			}
		} else if runeCharRange[0] <= nextChar && nextChar <= runeCharRange[2] {
			p.pos += 1
			return nextChar, nil
		}
	}

	return rune(0),
		utils.NewParseError(
			p.pos,
			"Expected from character class but no match",
			[]string{})
}

func (p *Parser) SplitCharRanges(charClass string) ([]string, error) {
	runeCharClass := []rune(charClass)
	rv := make([]string, 0)
	index := 0
	length := len(runeCharClass)
	for index < length {
		if index+2 < length && runeCharClass[index+1] == '-' {
			if runeCharClass[index] >= runeCharClass[index+2] {
				return []string{""}, errors.New("bad character range")
			}

			rv = append(rv, string(runeCharClass[index:index+3]))
			index += 3
		} else {
			rv = append(rv, string(runeCharClass[index]))
			index += 1
		}
	}

	return rv, nil
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
