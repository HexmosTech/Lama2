package parser

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/utils"
)

// Method Match is the most important of all in the
// parser package. Match takes in a slice of rules
// (essentially method names), and then executes
// them one by one. On successful match, we return
// a gabs Container with `error` set to `nil`
// When a rule fails to match, we reset the scan
// position to initial position; moreover, we keep
// a continuous track of the farthest/longest match
// till present. The farthest match error is potentially
// the most useful error message to the user; thus,
// for error report, Match returns the farthest matching
// error
func (p *Parser) Match(rules []string) (*gabs.Container, error) {
	p.eatWhitespace()
	lastErrorPos := -1
	lastErrorLine := 0
	lastErrorRules := []string{}
	lastError := errors.New("")

	for _, rule := range rules {
		initialPos := p.Pos
		res := p.ruleMethodMap[rule].Call([]reflect.Value{})
		op := res[0].Interface().(*gabs.Container)
		e := res[1]
		if e.IsNil() {
			p.eatWhitespace()
			return op, nil
		}

		p.Pos = initialPos
		pe := e.Interface().(*utils.ParseError)
		if pe.Pos > lastErrorPos {
			lastError = pe
			lastErrorPos = pe.Pos
			lastErrorLine = pe.LineNum
			lastErrorRules = nil
			lastErrorRules = append(lastErrorRules, rule)
		} else if pe.Pos == lastErrorPos {
			lastErrorRules = append(lastErrorRules, rule)
		}
	}

	if len(lastErrorRules) == 1 {
		return nil, lastError
	}

	if lastErrorPos >= p.TotalLen {
		lastErrorPos = p.TotalLen - 1
	}
	return nil, utils.NewParseError(lastErrorPos, lastErrorLine, "Expected %s but got %s", []string{strings.Join(lastErrorRules, ","), string(p.Text[lastErrorPos])})
}

func (p *Parser) LookAhead(rules []string) bool {
	p.eatWhitespace()

	for _, rule := range rules {
		initialPos := p.Pos
		res := p.ruleMethodMap[rule].Call([]reflect.Value{})
		op := res[0].Interface().(*gabs.Container)
		fmt.Println(op.String())
		e := res[1]
		p.Pos = initialPos
		if e.IsNil() {
			p.eatWhitespace()
			return true
		}
	}

	return false
}

func (p *Parser) MatchUntil(end string) (*gabs.Container, error) {
	chunk := make([]string, 0)
	temp := gabs.New()
	for {
		oldPos := p.Pos
		_, e1 := p.Keyword(end, false, false, false)
		if e1 == nil {
			p.Pos = oldPos
			temp.Set(strings.Join(chunk, ""))
			return temp, nil
		}

		r2, e2 := p.Char()
		if e2 != nil {
			break
		}
		chunk = append(chunk, string(r2))
	}
	temp.Set(strings.Join(chunk, ""))
	return temp, nil
}
