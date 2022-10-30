package parser

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/Jeffail/gabs/v2"
	"github.com/hexmos/lama2/utils"
)

func (p *Parser) Match(rules []string) (interface{}, error) {
	p.eatWhitespace()
	lastErrorPos := -1
	lastErrorRules := []string{}
	lastError := ""
	fmt.Printf("%d %s %s\n", lastErrorPos, lastErrorRules, lastError)

	for _, rule := range rules {
		res := p.ruleMethodMap[rule].Call([]reflect.Value{})
		fmt.Println("Res = ", res)
		op := res[0].Interface().(*gabs.Container)
		e := res[1]
		if e.IsNil() {
			fmt.Println(op.StringIndent("", "  "))
			return op, nil
		} else {
			pe := e.Interface().(utils.ParseError)
			fmt.Errorf(fmt.Sprintf("%s", pe.Error()))
			return nil, pe
		}
	}

	return "", errors.New("Match failed")
}
