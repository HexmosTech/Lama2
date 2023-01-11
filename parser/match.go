package parser

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/utils"
	"github.com/rs/zerolog/log"
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
		log.Trace().Str("Rule", rule).Strs("Rules", rules).Msg("")
		res := p.ruleMethodMap[rule].Call([]reflect.Value{})
		op := res[0].Interface().(*gabs.Container)
		log.Trace().Str("Rule res", op.String()).Msg("")
		e := res[1]
		log.Trace().Str("Rule error", e.String()).Msg("")
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

func (p *Parser) LookAhead(rules []string) (*gabs.Container, error) {
	p.eatWhitespace()
	lastErrorPos := -1
	lastErrorLine := 0
	lastErrorRules := []string{}
	lastError := errors.New("")

	for _, rule := range rules {
		initialPos := p.Pos
		log.Trace().Str("Rule", rule).Strs("Rules", rules).Msg("")
		r, _ := regexp.Compile(rule)
		match := r.FindString(string(p.Text[p.Pos+1:]))
		// op := res[0].Interface().(*gabs.Container)
		// log.Trace().Str("Rule res", op.String()).Msg("")
		// e := res[1]
		// log.Trace().Str("Rule error", e.String()).Msg("")
		fmt.Println(match)
		// if e == nil {
		// p.eatWhitespace()
		// return op, nil
		// }

		/*
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
		*/
	}

	if len(lastErrorRules) == 1 {
		return nil, lastError
	}

	if lastErrorPos >= p.TotalLen {
		lastErrorPos = p.TotalLen - 1
	}
	return nil, utils.NewParseError(lastErrorPos, lastErrorLine, "Expected %s but got %s", []string{strings.Join(lastErrorRules, ","), string(p.Text[lastErrorPos])})

}

func (p *Parser) RegexMatch(rules []string) (*gabs.Container, error) {
	p.eatWhitespace()
	lastErrorPos := -1
	lastErrorLine := 0
	lastErrorRules := []string{}
	lastError := errors.New("")

	for _, rule := range rules {
		// initialPos := p.Pos
		log.Trace().Str("Rule", rule).Strs("Rules", rules).Msg("")
		r, _ := regexp.Compile(rule)
		match := r.FindString(string(p.Text[p.Pos+1:]))
		// op := res[0].Interface().(*gabs.Container)
		// log.Trace().Str("Rule res", op.String()).Msg("")
		// e := res[1]
		// log.Trace().Str("Rule error", e.String()).Msg("")
		fmt.Println(match)
		// if e == nil {
		// p.eatWhitespace()
		// return op, nil
		// }

		/*
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
		*/
	}

	if len(lastErrorRules) == 1 {
		return nil, lastError
	}

	if lastErrorPos >= p.TotalLen {
		lastErrorPos = p.TotalLen - 1
	}
	return nil, utils.NewParseError(lastErrorPos, lastErrorLine, "Expected %s but got %s", []string{strings.Join(lastErrorRules, ","), string(p.Text[lastErrorPos])})

}
