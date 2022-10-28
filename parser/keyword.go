package parser

import (
	"reflect"
	"strings"

	"github.com/hexmos/lama2/utils"
)

func (p *Parser) Keyword(kw string, eat_ws_start bool, eat_ws_end bool, case_insensitive bool) ([]rune, error) {
	if p.pos >= p.tlen {
		return []rune{rune(0)},
			utils.NewParseError(
				p.pos,
				"Expected %s; but no such keyword found",
				[]string{kw})
	}

	if eat_ws_start {
		p.eatWhitespace()
	}

	kwRune := []rune(kw)
	low := p.pos + 1
	high := low + len(kwRune)

	if case_insensitive {
		ipText := strings.ToLower(string(p.text[low:high]))
		kwLower := strings.ToLower(kw)

		if reflect.DeepEqual(ipText, kwLower) {
			p.pos += len(kwRune)
			if eat_ws_end {
				p.eatWhitespace()
			}
			return []rune(kw), nil
		}
	} else {
		if reflect.DeepEqual(p.text[low:high], kwRune) {
			p.pos += len(kwRune)
			if eat_ws_end {
				p.eatWhitespace()
			}
			return []rune(kw), nil
		}
	}

	return []rune{rune(0)}, utils.NewParseError(
		p.pos,
		"Expected %s; but no such keyword found",
		[]string{kw})
}
