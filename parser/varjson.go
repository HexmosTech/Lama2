package parser

import (
	"strings"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/utils"
)

// Method VarJSON behaves in two ways depending
// on whether `multipart` or `form` is true or not.
// If there is no multipart, then VarJSON tries
// to match one or more VarJSONPairs
// However, if there is multipart or form, we try to match
// zero or more VarJSON, followed by zero or more
// file fields (separated by `@`). If there is no match
// at all, we return a ParseError; otherwise the
// we return the parsed data.
func (p *Lama2Parser) VarJSON() (*gabs.Container, error) {
	temp := gabs.New()
	hasMultipart := false
	if val, ok := p.Context["multipart"]; ok {
		hasMultipart = val
	}

	pair, e1 := p.Match([]string{"VarJSONPair"})
	// Errors out when there is no key-value pair and multipart is not enabled,
	// Multipart may contain file so thats why its not erroring out.
	if e1 != nil && (!hasMultipart) { 
		return nil, e1
	}
	
	temp.Merge(pair)

	for {
		pair, e1 = p.Match([]string{"VarJSONPair"})
		if e1 != nil {
			break
		}

		temp.Merge(pair)
	}

	if hasMultipart {
		filesObj := gabs.New()
		for {
			pair, e1 = p.Match([]string{"FilesPair"})
			if e1 != nil {
				break
			}
			filesObj.Merge(pair)
		}
		temp.Set(filesObj, "@files")
	}

	if len(temp.ChildrenMap()) == 0 {
		return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "No varjson match", []string{})
	}

	return temp, nil
}

// VarJSONPair tries to match key and value separated
// by `=`. The key and value can either be a quoted
// string, or an unquoted VarJSON unquoted string. If there is no
// match for either, a ParseError is returned.
func (p *Lama2Parser) VarJSONPair() (*gabs.Container, error) {
	temp := gabs.New()
	key, e := p.Match([]string{"QuotedString"})
	if e != nil {
		key, e = p.Match([]string{"VarJSONUnquoted"})
		if e != nil {
			return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "Couldn't match either quoted or unquoted VarJSON string", []string{})
		}
	}

	_, e2 := p.Keyword("=", true, true, true)
	if e2 != nil {
		return nil, e2
	}

	value, err := p.Match([]string{"QuotedString"})
	if err != nil {
		value, err = p.Match([]string{"VarJSONUnquoted"})
		if err != nil {
			return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "In header value, couldn't get string", []string{})
		}
	}
	valueStr, _ := value.Data().(string)
	keyStr, _ := key.Data().(string)
	temp.Set(valueStr, keyStr)

	return temp, nil
}

// FilesPair tries to match key and value separated
// by `@`. The key and value can either be a quoted
// string, or an unquoted Files Unquoted String. If there is no
// match for either, a ParseError is returned.
func (p *Lama2Parser) FilesPair() (*gabs.Container, error) {
	temp := gabs.New()
	key, e := p.Match([]string{"QuotedString"})
	if e != nil {
		key, e = p.Match([]string{"FilesUnquoted"})
		if e != nil {
			return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "Couldn't match either quoted or unquoted Files string", []string{})
		}
	}

	_, e2 := p.Keyword("@", true, true, true)
	if e2 != nil {
		return nil, e2
	}

	value, err := p.Match([]string{"QuotedString"})
	if err != nil {
		value, err = p.Match([]string{"FilesUnquoted"})
		if err != nil {
			return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "In File value, couldn't get string", []string{})
		}
	}
	valueStr, _ := value.Data().(string)
	keyStr, _ := key.Data().(string)
	temp.Set(valueStr, keyStr)

	return temp, nil
}

// VarJSONUnquoted matches a string of characters other than `=`
// and returns them as a String
func (p *Lama2Parser) VarJSONUnquoted() (*gabs.Container, error) {
	acceptableChars := "@0-9A-Za-z \t!$%&()*+./;<>?^_`|~-"
	chars := make([]string, 0)
	c, e := p.CharClass(acceptableChars)
	if e != nil {
		return nil, utils.NewParseError(p.Pos+1, p.LineNum+1,
			"No match for Unquoted Varjson string", []string{})
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

// FilesUnquoted matches a string of characters other than `@`
// and returns them as a String
func (p *Lama2Parser) FilesUnquoted() (*gabs.Container, error) {
	acceptableChars := "0-9A-Za-z \t!$%&()*+./;<>?^_`|~-"
	chars := make([]string, 0)
	c, e := p.CharClass(acceptableChars)
	if e != nil {
		return nil, utils.NewParseError(p.Pos+1, p.LineNum+1,
			"No match for Unquoted Files string", []string{})
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
