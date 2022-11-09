package parser

import (
	"errors"
	"fmt"
	"os"

	"github.com/HexmosTech/lama2/utils"
)

func (p *Parser) Char() (rune, error) {
	if p.Pos >= p.TotalLen {
		return rune(0),
			utils.NewParseError(
				p.Pos, p.LineNum+1,
				"Expected %s but got end of string",
				[]string{"character"})
	}
	nextChar := p.Text[p.Pos+1]
	p.Pos++
	return nextChar, nil
}

func (p *Parser) CharClass(charClass string) (rune, error) {
	if p.Pos >= p.TotalLen {
		return rune(0),
			utils.NewParseError(
				p.Pos, p.LineNum+1,
				"Expected %s but got end of string",
				[]string{"character"})
	}
	nextChar := p.Text[p.Pos+1]
	charRangeList, e := p.SplitCharRanges(charClass)
	if e != nil {
		fmt.Println(fmt.Errorf("%s", e))
		os.Exit(1)
	}

	for _, charRange := range charRangeList {
		runeCharRange := []rune(charRange)
		if len(runeCharRange) == 1 {
			if nextChar == runeCharRange[0] {
				p.Pos++
				return nextChar, nil
			}
		} else if runeCharRange[0] <= nextChar && nextChar <= runeCharRange[2] {
			p.Pos++
			return nextChar, nil
		}
	}

	return rune(0),
		utils.NewParseError(
			p.Pos, p.LineNum+1,
			"Expected %s from character class but no match",
			[]string{charClass})
}

func (p *Parser) SplitCharRanges(charClass string) ([]string, error) {
	val, prs := p.cache[charClass]
	if prs {
		return val, nil
	}

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
			index++
		}
	}

	p.cache[charClass] = rv
	return rv, nil
}
