package parser

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"unicode/utf16"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/utils"
)

func (p *Lama2Parser) readUnicodeHexStr() ([]string, error) {
	codePoint := make([]string, 0)
	for i := 0; i < 4; i++ {
		r, e := p.CharClass("0-9a-fA-F")
		if e == nil {
			codePoint = append(codePoint, string(r))
		} else {
			return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "Invalid code point", []string{})
		}
	}
	return codePoint, nil
}

// QuotedString accepts both single-quoted and double-quoted
// types of strings. Moreover, it can deal with unicode
// escape characters, control characters appropriately
// Ultimately, we get a string wrapped in a gabs container
func (p *Lama2Parser) QuotedString() (*gabs.Container, error) {
	temp := gabs.New()
	chars := []string{}
	quote, e := p.CharClass("\"'")
	if e != nil {
		return nil, utils.NewParseError(p.Pos+1, p.LineNum+1,
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
	escapeSequences[string(quote)] = string(quote)

	for {
		c, _ := p.Char()
		if c == quote {
			break
		} else if c == '\\' {
			escape, _ := p.Char()
			var codePoint []string
			var e error
			if escape == 'u' {
				codePoint, e = p.readUnicodeHexStr()
				if e != nil {
					return nil, e
				}
				hex := strings.Join(codePoint, "")
				n := new(big.Int)
				n.SetString(hex, 16)
				rune1 := rune(n.Int64())
				if utf16.IsSurrogate(rune1) {
					p.CharClass("\\")
					p.CharClass("u")
					codePoint, e = p.readUnicodeHexStr()
					if e != nil {
						return nil, e
					}
					hex2 := strings.Join(codePoint, "")
					n2 := new(big.Int)
					n2.SetString(hex2, 16)
					rune2 := rune(n2.Int64())
					res2 := utf16.DecodeRune(rune1, rune2)
					chars = append(chars, string(res2))

				} else {
					r2 := strconv.FormatInt(n.Int64(), 16)
					r3 := fmt.Sprintf("\\u%04s", r2)
					r4, _ := strconv.Unquote(`"` + r3 + `"`)
					// Uncomment to read the string slice:
					// r4Rune := []rune(r4)
					chars = append(chars, r4)
				}
			} else {
				if val, ok := escapeSequences[string(escape)]; ok {
					chars = append(chars, val)
				} else {
					return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "Invalid escape sequence", []string{})
				}
			}
		} else {
			if utils.UnicodeCategory(c) == "Cc" {
				return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "Unescaped control character", []string{})
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
		return nil, utils.NewParseError(p.Pos+1, p.LineNum+1,
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
