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
		parserFns := make([]reflect.Value, 0)
		parserFnsNames := make([]string, 0)

		pmFns := make([]reflect.Value, 0)
		pmFnsNames := make([]string, 0)

		pVal := reflect.ValueOf(p)
		fmt.Println("pValElem.NumMethod", pVal.NumMethod())
		for i := 0; i < pVal.NumMethod(); i++ {
			methodVal := pVal.Method(i)
			parserFns = append(parserFns, methodVal)
		}
		fmt.Println("parserFns", parserFns, len(parserFns))

		tVal := reflect.TypeOf(p)
		for i := 0; i < tVal.NumMethod(); i++ {
			m := tVal.Method(i)
			parserFnsNames = append(parserFnsNames, m.Name)
		}
		fmt.Println("ParserFnsNames", parserFnsNames, len(parserFnsNames))

		//////

		pmVal := reflect.Indirect(pVal).FieldByName("Pm").Elem()
		for i := 0; i < pmVal.NumMethod(); i++ {
			methodVal := pmVal.Method(i)
			pmFns = append(pmFns, methodVal)
			// res := methodVal.Call([]reflect.Value{})
			// name := runtime.FuncForPC(method.Pointer()).Name()
			// fmt.Println(res)
		}
		fmt.Println("pmFns", pmFns, len(pmFns))

		tLamaVal := reflect.TypeOf(p.Pm)
		for i := 0; i < tLamaVal.NumMethod(); i++ {
			m := tLamaVal.Method(i)
			pmFnsNames = append(pmFnsNames, m.Name)
		}
		fmt.Println("pmFnsNames", pmFnsNames, len(pmFnsNames))

		fmt.Printf("Rule %d: %s\n", index, rule)
	}

	return "", errors.New("Match failed")
}
