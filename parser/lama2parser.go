package parser

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
	"github.com/hexmos/lama2/utils"
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
	res, e := p.MustMatch([]string{"HttpVerb"})
	if e == nil {
		temp := gabs.New()
		temp.Set(res, "value")
		temp.Set("HttpFile", "type")
		return temp, nil
	} else {
		return nil, e
	}
}

func (p *Lama2Parser) HttpVerb() (*gabs.Container, error) {
	fmt.Println("Withint HttpVerb")
	verbList := []string{"get", "head", "post", "put",
		"delete", "connect", "trace", "patch"}

	for _, k := range verbList {
		r, e := p.Keyword(k, true, true, true)
		if e == nil {
			temp := gabs.New()
			temp.Set(string(r), "value")
			temp.Set("HttpVerb", "type")
			return temp, nil
		}
	}
	return nil, utils.NewParseError(p.Pos+1,
		"Couldn't find the http verb (get/post/put/delete/connect/options/trace/patch/head", []string{})
}

func (p *Lama2Parser) Multipart() (*gabs.Container, error) {
	_, e := p.Keyword("multipart", true, true, true)
	if e == nil {
		temp := gabs.New()
		temp.Set(true, "value")
		temp.Set("Multipart", "type")
		return temp, nil
	}
	return nil, utils.NewParseError(p.Pos+1,
		"Expected 'multipart', but couldn't find", []string{})
}
