package parser

import "github.com/Jeffail/gabs/v2"

type Lama2Parser struct {
	*Parser
}

func NewLama2Parser() *Lama2Parser {
	p := &Lama2Parser{&Parser{}}
	p.pm = p
	return p
}

func (p *Lama2Parser) Start() *gabs.Container {
	temp := gabs.New()
	p.httpFile()
	return temp
}

func (p *Lama2Parser) httpFile() *gabs.Container {
	temp := gabs.New()
	return temp
}
