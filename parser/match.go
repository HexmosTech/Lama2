package parser

import (
	"errors"
	"fmt"
	"reflect"
)

func (p *Parser) Match(rules []string) (interface{}, error) {
	p.eatWhitespace()
	lastErrorPos := -1
	lastErrorRules := []string{}
	lastError := ""
	fmt.Printf("%d %s %s\n", lastErrorPos, lastErrorRules, lastError)

	for index, rule := range rules {
		res := p.ruleMethodMap[rule].Call([]reflect.Value{})
		fmt.Println("Res = ", res)
		fmt.Printf("Rule %d: %s\n", index, rule)
	}

	return "", errors.New("Match failed")
}
