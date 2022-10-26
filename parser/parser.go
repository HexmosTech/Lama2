package parser

import (
	"github.com/Jeffail/gabs/v2"
)

type Parser struct {
	text  string
	pos   int
	tlen  int
	cache *gabs.Container
}

func (p Parser) Init() {
	p.cache = gabs.New()
}

func (p Parser) Parse(text string) string {
	p.text = text
	p.pos = -1
	p.tlen = len(text) - 1
	// TODO: self.start()??
	return "blah2"
}

func (p Parser) assertEnd() {
	if p.pos < p.tlen {

	}
}
