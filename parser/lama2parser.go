package parser

import (
	"fmt"
	"strings"

	"github.com/Jeffail/gabs/v2"
	"github.com/hexmos/lama2/utils"
)

type Lama2Parser struct {
	*Parser
	Context map[string]bool
}

func NewLama2Parser() *Lama2Parser {
	p := &Lama2Parser{&Parser{}, make(map[string]bool)}
	p.Pm = p
	p.Init()
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
	res, e := p.Match([]string{"HttpVerb"})
	temp := gabs.New()
	if e == nil {
		temp.Set(res, "verb")
		temp.Set("HttpFile", "type")
	} else {
		return nil, e
	}
	res, e = p.Match([]string{"Multipart"})
	if e == nil {
		temp.Set(res, "multipart")
	}

	res, e = p.Match([]string{"TheUrl"})
	if e == nil {
		temp.Set(res, "url")
	} else {
		return nil, e
	}

	res, e = p.Match([]string{"Details"})
	if e == nil {
		temp.Set(res, "details")
	}
	return temp, nil
}

func (p *Lama2Parser) TheUrl() (*gabs.Container, error) {
	res := []string{}
	kw, e := p.Keyword("http", true, false, true)
	if e == nil {
		res = append(res, string(kw))
	} else {
		return nil, utils.NewParseError(p.Pos+1, "Couldn't find URL (starting with http(s)", []string{})
	}

	_, e = p.CharClass("s")
	if e == nil {
		res = append(res, "s")
	}

	_, e = p.Keyword("://", false, false, true)
	if e == nil {
		res = append(res, "://")
	} else {
		return nil, utils.NewParseError(p.Pos+1, "Couldn't find URL (starting with http(s)://", []string{})
	}

	for {
		up, err := p.CharClass("A-Za-z0-9-._~:/?#[]@!$&'()*+,;%=")
		if err == nil {
			res = append(res, string(up))
		} else {
			break
		}
	}

	val := strings.Join(res, "")
	temp := gabs.New()
	temp.Set(val, "value")
	temp.Set("TheUrl", "type")
	return temp, nil
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
		p.Context["multipart"] = true
		return temp, nil
	}
	return nil, utils.NewParseError(p.Pos+1,
		"Expected 'multipart', but couldn't find", []string{})
}

func (p *Lama2Parser) Details() (*gabs.Container, error) {
	// temp := gabs.New()
	res, e := p.Match([]string{"HeaderData", "DataHeader"})
	if e == nil {
		// temp = utils.SetJson(temp, res, "contents")
		return res, nil
	}

	return nil, utils.NewParseError(p.Pos+1,
		"No details packet found", []string{})
}
