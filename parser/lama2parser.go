package parser

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

type Lama2Parser struct {
	*Parser
}

func NewLama2Parser() *Lama2Parser {
	p := &Lama2Parser{&Parser{}}
	p.Pm = p
	return p
}

func (p *Lama2Parser) Start() (*gabs.Container, error) {
	fmt.Println("Within the Start function!")
	temp := gabs.New()
	res, e := p.Match([]string{"HttpFile"})
	if e == nil {
		temp.Set(res, "value")
		temp.Set("Lama2Parser", "type")
		return temp, nil
	} else {
		return nil, e
	}
}

func (p *Lama2Parser) HttpFile() (*gabs.Container, error) {
	fmt.Println("Within HttpFile")
	res, e := p.Keyword("GET", true, true, true)
	if e == nil {
		temp := gabs.New()
		temp.Set(string(res), "value")
		temp.Set("HttpFile", "type")
		return temp, nil
	} else {
		return nil, e
	}
}
