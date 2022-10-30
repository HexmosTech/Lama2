package parser

import (
	"errors"
	"fmt"
	"reflect"

	"strings"

	"github.com/Jeffail/gabs/v2"
	"github.com/hexmos/lama2/utils"
)

func (p *Parser) Match(rules []string) (interface{}, error) {
	p.eatWhitespace()
	lastErrorPos := -1
	lastErrorRules := []string{}
	lastError := errors.New("")
	fmt.Printf("%d %s %s\n", lastErrorPos, lastErrorRules, lastError)

	for _, rule := range rules {
		res := p.ruleMethodMap[rule].Call([]reflect.Value{})
		fmt.Println("Res = ", res)
		op := res[0].Interface().(*gabs.Container)
		e := res[1]
		if e.IsNil() {
			fmt.Println(op.StringIndent("", "  "))
			p.eatWhitespace()
			return op, nil
		} else {
			pe := e.Interface().(utils.ParseError)
			fmt.Println(fmt.Errorf("%s", pe.Error()))
			if pe.Pos > lastErrorPos {
				lastError = pe
				lastErrorPos = pe.Pos
				lastErrorRules = nil
				lastErrorRules = append(lastErrorRules, rule)
			} else if pe.Pos == lastErrorPos {
				lastErrorRules = append(lastErrorRules, rule)
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
	}

	return "", errors.New("Match failed")
}
