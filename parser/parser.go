/*
The `parser` package provides primitives that
help with writing recursive descent parsers.
This version is a golang port of the original
Python implementation from https://tinyurl.com/rdescent

The `Parser` struct is supposed to be extended
to support parsing a new language. Take a look
at `lama2parser.go` for an example. Essentially
the actual parsing begins from the `Start()` method.
*/
package parser

import (
	"reflect"
	"strings"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/utils"
)

// MinimalParser enforces concrete Types
// to have a Start() method, from which
// parsing process begins. In the present
// case, `Lama2Parser` adds up dozens of
// of methods to implement `.l2` syntax
type MinimalParser interface {
	Start() (*gabs.Container, error)
}

// Struct Parser stores information about
// the parsing process throughout.
// 1. Text: Incoming text is stored as an array
// of runes, to correctly handle unicode characters
// 2. Pos: Indicates the index position in Text which
// has already been scanned; starts with -1
// 3. TotalLen: Number of runes in the input
// 4. Pm: Composing an external MinimalParser (such as
// Lama2Parser) which builds upon Parser to provide
// the new language recognition capabilities
// 5. ruleMethodMap: Scans through Pm, and creates a
// mapping from method name to method value through
// reflection
// 6. LineNum: Number of normalized newlines found till
// now. Used in providing useful context in error messages
type Parser struct {
	Text          []rune
	Pos           int
	TotalLen      int
	cache         map[string][]string
	Pm            MinimalParser
	ruleMethodMap map[string]reflect.Value
	LineNum       int
}

// Start() in Parser provides a dummy default
// implementation; the expectation is that the
// higher level Struct (Pm) will implement its
// own version
func (p *Parser) Start() *gabs.Container {
	temp := gabs.New()
	return temp
}

// Method Init creates the most important data stucture
// for parsing: ruleMethodMap. We use reflection to create
// a mapping of each Pm.<method_name> to <method_value>
func (p *Parser) Init() {
	p.ruleMethodMap = make(map[string]reflect.Value, 0)
	pVal := reflect.ValueOf(p)
	pmVal := reflect.Indirect(pVal).FieldByName("Pm").Elem()
	tLamaVal := reflect.TypeOf(p.Pm)
	for i := 0; i < pmVal.NumMethod(); i++ {
		m := tLamaVal.Method(i)
		methodVal := pmVal.Method(i)
		p.ruleMethodMap[m.Name] = methodVal
	}
}

// Method Parse normalizes newlines and then creates
// a rune version of the input data. The Start() method
// proceeds to process the rune version of data
func (p *Parser) Parse(text string) (*gabs.Container, error) {
	text = strings.ReplaceAll(text, "\r\n", "\n")
	p.Text = []rune(text)
	p.Pos = -1
	p.TotalLen = len(p.Text) - 1
	p.cache = make(map[string][]string)
	p.LineNum = 0
	res, _ := p.Pm.Start()
	_, err := p.assertEnd()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *Parser) eatWhitespace() {
	spaces := []rune{' ', '\f', '\v', '\r', '\t', '\n'}
	isProcessingComment := false
	for p.Pos < p.TotalLen {
		c := p.Text[p.Pos+1]
		if c == '\n' {
			p.LineNum++
		}
		if isProcessingComment {
			if c == '\n' {
				isProcessingComment = false
			}
		} else {
			if c == '#' {
				isProcessingComment = true
			} else if !utils.ContainsRune(spaces, c) {
				break
			}
		}
		p.Pos++
	}
}

// Method SetText is a utility used
// primarily in testing, when we don't
// want to call Start() automatically
// as in Parse
func (p *Parser) SetText(text string) {
	p.Text = []rune(text)
	p.Pos = -1
	p.TotalLen = len(p.Text) - 1
	p.cache = make(map[string][]string)
}

func (p *Parser) assertEnd() (bool, error) {
	if p.Pos < p.TotalLen {
		return false,
			utils.NewParseError(
				p.Pos+1,
				p.LineNum+1,
				"Expected end of string but got %s",
				[]string{string(p.Text[p.Pos+1])})
	}
	return true, nil
}
