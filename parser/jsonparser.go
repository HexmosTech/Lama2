package parser

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/Jeffail/gabs/v2"
	"github.com/hexmos/lama2/utils"
)

func (p *Lama2Parser) QuotedString() (*gabs.Container, error) {
	temp := gabs.New()
	chars := []string{}
	quote, e := p.CharClass("\"'")
	if e != nil {
		return nil, utils.NewParseError(p.Pos+1,
			"Couldn't find quoted string", []string{})
	}
	escapeSequences := map[string]string{
		"b":  "\b",
		"f":  "\f",
		"n":  "\n",
		"r":  "\r",
		"t":  "\t",
		"\\": "\\",
		"/":  "/",
	}

	// TODO: The following is slightly different from
	// the original; may require rework
	escapeSequences[string(quote)] = fmt.Sprintf("\\%s", quote)

	for {
		c, _ := p.Char()
		if c == quote {
			break
		} else if c == '\\' {
			escape, _ := p.Char()
			codePoint := []string{}
			if escape == 'u' {
				for i := 0; i < 4; i++ {
					r, e := p.CharClass("0-9a-fA-F")
					if e == nil {
						codePoint = append(codePoint, string(r))
					} else {
						return nil, utils.NewParseError(p.Pos+1, "Invalid code point", []string{})
					}
					hex := strings.Join(codePoint, "")
					n := new(big.Int)
					n.SetString(hex, 16)
					s := fmt.Sprintf("%s", n)
					chars = append(chars, s)
				}
			} else {
				if val, ok := escapeSequences[string(escape)]; ok {
					chars = append(chars, val)
				} else {
					return nil, utils.NewParseError(p.Pos+1, "Invalid escape sequence", []string{})
				}
			}
		} else {
			if utils.UnicodeCategory(c) == "Cc" {
				return nil, utils.NewParseError(p.Pos+1, "Unescaped control character", []string{})
			}

			chars = append(chars, string(c))
		}
	}

	res := strings.Join(chars, "")
	temp.Set(res, "value")
	return temp, nil
}
