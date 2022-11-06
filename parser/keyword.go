package parser

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/HexmosTech/lama2/utils"
)

func (p *Parser) Keyword(kw string, eat_ws_start bool, eat_ws_end bool, case_insensitive bool) ([]rune, error) {
	if p.Pos >= p.TotalLen || p.Pos+len(kw) > p.TotalLen {
		return []rune{rune(0)},
			utils.NewParseError(
				p.Pos,
				"Expected %s; but no such keyword found",
				[]string{kw})
	}

	if eat_ws_start {
		p.eatWhitespace()
	}

	kwRune := []rune(kw)
	low := p.Pos + 1
	high := low + len(kwRune)

	if case_insensitive {
		ipText := strings.ToLower(string(p.Text[low:high]))
		kwLower := strings.ToLower(kw)

		if reflect.DeepEqual(ipText, kwLower) {
			p.Pos += len(kwRune)
			if eat_ws_end {
				p.eatWhitespace()
			}
			return []rune(kw), nil
		}
	} else {
		if reflect.DeepEqual(p.Text[low:high], kwRune) {
			p.Pos += len(kwRune)
			if eat_ws_end {
				p.eatWhitespace()
			}
			return []rune(kw), nil
		}
	}

	return []rune{rune(0)}, utils.NewParseError(
		p.Pos,
		"Expected %s; but no such keyword found",
		[]string{kw})
}

func (p *Parser) MustKeyword(kw string, eat_ws_start bool, eat_ws_end bool, case_insensitive bool) ([]rune, error) {
	res, e := p.Keyword(kw, eat_ws_start, eat_ws_end, case_insensitive)
	if e == nil {
		return res, e
	} else {
		fmt.Println(fmt.Errorf("%s", e.Error()))
		os.Exit(1)
	}
	return []rune{' '}, nil
}
