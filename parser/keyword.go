package parser

import (
	"reflect"
	"strings"

	"github.com/HexmosTech/lama2/utils"
)

func (p *Parser) Keyword(kw string, eatWsStart bool, eatWsEnd bool, caseInsensitive bool) ([]rune, error) {
	if p.Pos >= p.TotalLen {
		return []rune{rune(0)},
			utils.NewParseError(
				p.Pos,
				p.LineNum+1,
				"Expected %s; but no such keyword found",
				[]string{kw})
	}

	if eatWsStart {
		p.eatWhitespace()
	}

	kwRune := []rune(kw)
	low := p.Pos + 1
	high := low + len(kwRune)

	if high > p.TotalLen+1 {
		return []rune{rune(0)},
			utils.NewParseError(
				p.Pos,
				p.LineNum+1,
				"Expected %s; but EOF",
				[]string{kw})
	}

	if caseInsensitive {
		ipText := strings.ToLower(string(p.Text[low:high]))
		kwLower := strings.ToLower(kw)

		if reflect.DeepEqual(ipText, kwLower) {
			p.Pos += len(kwRune)
			if eatWsEnd {
				p.eatWhitespace()
			}
			return []rune(kw), nil
		}
	} else {
		if reflect.DeepEqual(p.Text[low:high], kwRune) {
			p.Pos += len(kwRune)
			if eatWsEnd {
				p.eatWhitespace()
			}
			return []rune(kw), nil
		}
	}

	return []rune{rune(0)}, utils.NewParseError(
		p.Pos, p.LineNum+1,
		"Expected %s; but no such keyword found",
		[]string{kw})
}
