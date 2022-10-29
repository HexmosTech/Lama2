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

func (p *Lama2Parser) Start() *gabs.Container {
	fmt.Println("Within the Start function!")
	temp := gabs.New()
	p.HttpFile()
	return temp
}

func (p *Lama2Parser) HttpFile() *gabs.Container {
	fmt.Println("Within HttpFile")
	temp := gabs.New()
	return temp
}
