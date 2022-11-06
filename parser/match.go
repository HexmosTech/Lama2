package parser

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"strings"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/utils"
)

func (p *Parser) Match(rules []string) (*gabs.Container, error) {
	p.eatWhitespace()
	lastErrorPos := -1
	lastErrorRules := []string{}
	lastError := errors.New("")
	fmt.Printf("%d %s %s\n", lastErrorPos, lastErrorRules, lastError)

	for _, rule := range rules {
		initialPos := p.Pos
		fmt.Println("Trying rule: ", rule, "rules = ", rules)
		res := p.ruleMethodMap[rule].Call([]reflect.Value{})
		fmt.Println("Res = ", res)
		op := res[0].Interface().(*gabs.Container)
		e := res[1]
		fmt.Println("e = ", e)
		if e.IsNil() {
			fmt.Println("here", rule)
			fmt.Println(op.StringIndent("", "  "))
			p.eatWhitespace()
			return op, nil
		} else {
			p.Pos = initialPos
			fmt.Println("Rule error: ", rule)
			pe := e.Interface().(*utils.ParseError)
			fmt.Println(fmt.Errorf("%s", pe.Error()))
			if pe.Pos > lastErrorPos {
				lastError = pe
				lastErrorPos = pe.Pos
				lastErrorRules = nil
				lastErrorRules = append(lastErrorRules, rule)
			} else if pe.Pos == lastErrorPos {
				lastErrorRules = append(lastErrorRules, rule)
			}
			fmt.Println("After the else stuff")
		}
	}

	if len(lastErrorRules) == 1 {
		return nil, lastError
	} else {
		if lastErrorPos >= p.TotalLen {
			lastErrorPos = p.TotalLen - 1
		}
		return nil, utils.NewParseError(lastErrorPos, "Expected %s but got %s", []string{strings.Join(lastErrorRules, ","), string(p.Text[lastErrorPos])})
	}

	return nil, errors.New("Match failed")
}

func (p *Parser) MustMatch(rules []string) (interface{}, error) {
	res, e := p.Match(rules)
	if e == nil {
		return res, e
	} else {
		fmt.Println(fmt.Errorf("%s", e.Error()))
		os.Exit(1)
	}
	return "", nil
}
