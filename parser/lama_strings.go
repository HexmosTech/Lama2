package parser

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/utils"
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
	temp.Set(res)
	return temp, nil
}

func (p *Lama2Parser) Unquoted() (*gabs.Container, error) {
	acceptableChars := "0-9A-Za-z \t!$%&()*+./;<=>?^_`|~-"
	chars := make([]string, 0)
	c, e := p.CharClass(acceptableChars)
	if e != nil {
		return nil, utils.NewParseError(p.Pos+1,
			"No match for Unquoted string", []string{})
	}
	chars = append(chars, string(c))
	for {
		c, e = p.CharClass(acceptableChars)
		if e != nil {
			break
		}
		chars = append(chars, string(c))
	}
	s := strings.TrimSpace(strings.Join(chars, ""))
	temp := gabs.New()
	temp.Set(s)
	return temp, nil
}
