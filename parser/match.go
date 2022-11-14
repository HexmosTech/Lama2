package parser

import (
	"errors"
	"reflect"
	"strings"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/utils"
	"github.com/rs/zerolog/log"
)

func (p *Parser) Match(rules []string) (*gabs.Container, error) {
	p.eatWhitespace()
	lastErrorPos := -1
	lastErrorLine := 0
	lastErrorRules := []string{}
	lastError := errors.New("")

	for _, rule := range rules {
		initialPos := p.Pos
		log.Trace().Str("Rule", rule).Strs("Rules", rules)
		res := p.ruleMethodMap[rule].Call([]reflect.Value{})
		op := res[0].Interface().(*gabs.Container)
		log.Trace().Str("Rule res", op.String())
		e := res[1]
		log.Trace().Str("Rule error", e.String())
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
