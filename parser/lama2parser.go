package parser

import (
	"errors"
	"strings"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/utils"
)

type Lama2Parser struct {
	*Parser
	Context   map[string]bool
	MarkRange map[string]int
}

// NewLama2Parser creates a new Lama2Parser
// and initializes it properly
func NewLama2Parser() *Lama2Parser {
	p := &Lama2Parser{&Parser{}, make(map[string]bool), make(map[string]int)}
	p.Pm = p
	p.Init()
	return p
}

// Start primarily calls the Lama2File method
func (p *Lama2Parser) Start() (*gabs.Container, error) {
	temp := gabs.New()
	res, e := p.Match([]string{"Lama2File"})
	if e == nil {
		temp.Set(res, "value")
		temp.Set("Lama2Parser", "type")
		return temp, nil
	}
	return nil, e
}

func (p *Lama2Parser) Lama2File() (*gabs.Container, error) {
	// Trying to get:
	// PSBlock? Requestor [SPSBlock Requestor]*
	temp := gabs.New()
	tempArr, e1 := temp.Array()
	if e1 != nil {
		return nil, errors.New("couldn't create Array for the parsed data")
	}
	// optionally match processor
	res2, procE1 := p.Match([]string{"Processor"})
	if procE1 == nil {
		tempArr.ArrayAppend(res2)
	}

	_, sepE1 := p.Match([]string{"Separator"})
	if procE1 != nil && sepE1 == nil {
		return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "Separator without preceding processor block found", []string{})
	} else if procE1 == nil && sepE1 != nil {
		return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "Processor without subsequent requestor block found", []string{})
	}

	// match requester
	res3, reqE1 := p.Match([]string{"Requester"})
	if reqE1 != nil {
		return nil, reqE1
	}

	tempArr.ArrayAppend(res3)

	// until file is done:
	var res4, res5 *gabs.Container
	var procE2, e5 error
	for p.Pos < p.TotalLen {
		_, sepE2 := p.Match([]string{"Separator"})
		if sepE2 != nil {
			return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "Separator expected (1)", []string{})
		}
		// match processor
		res4, procE2 = p.Match([]string{"Processor"})
		if procE2 != nil {
			return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "Processor expected (2)", []string{})
		}

		if p.Pos < p.TotalLen {
			_, sepE3 := p.Match([]string{"Separator"})
			if sepE3 != nil {
				return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "Separator exepcted (3)", []string{})
			}

		}

		// match requester
		if p.Pos < p.TotalLen {
			res5, e5 = p.Match([]string{"Requester"})
			if e5 != nil {
				return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "Request block exepcted (4)", []string{})
			}

			tempArr.ArrayAppend(res4)
			tempArr.ArrayAppend(res5)
		}
	}
	return tempArr, nil
}

func (p *Lama2Parser) Processor() (*gabs.Container, error) {
	// A Processor cannot start with any of the HTTP Verbs
	res := p.LookAhead([]string{"HTTPVerb"})
	if res {
		return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "HTTPVerb found at start of block; cannot be a Requestor block", []string{})
	}
	temp := gabs.New()
	res2, _ := p.MatchUntil("\n---\n")
	temp.Set("processor", "type")
	temp.Set(res2, "value")
	return temp, nil
}

func (p *Lama2Parser) Separator() (*gabs.Container, error) {
	temp := gabs.New()
	s, e := p.Keyword("---\n", false, false, false)
	temp.Set(string(s))
	if e != nil {
		return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "Couldnt' find separator", []string{})
	}
	return temp, nil
}

// Requester applies the rule:
// HTTPVerb Multipart? TheURL Details?
func (p *Lama2Parser) Requester() (*gabs.Container, error) {
	res, e := p.Match([]string{"HTTPVerb"})
	temp := gabs.New()
	if e == nil {
		temp.Set(res, "verb")
		temp.Set("Lama2File", "type")
	} else {
		return nil, e
	}
	res, e = p.Match([]string{"Multipart"})
	if e == nil {
		temp.Set(res, "multipart")
	}

	res, e = p.Match([]string{"Form"})
	if e == nil {
		temp.Set(res, "form")
	}

	res, e = p.Match([]string{"TheURL"})
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

func (p *Lama2Parser) TheURL() (*gabs.Container, error) {
	res := []string{}
	for {
		up, err := p.CharClass("A-Za-z0-9-._~:/?#[]@!$&'()*+,;%=}{")
		if err == nil {
			res = append(res, string(up))
		} else {
			break
		}
	}

	val := strings.Join(res, "")
	temp := gabs.New()
	temp.Set(val, "value")
	temp.Set("TheURL", "type")
	return temp, nil
}

func (p *Lama2Parser) HTTPVerb() (*gabs.Container, error) {
	verbList := []string{
		"get", "head", "post", "put",
		"delete", "connect", "trace", "patch",
	}

	for _, k := range verbList {
		r, e := p.Keyword(k, true, true, true)
		if e == nil {
			temp := gabs.New()
			temp.Set(string(r), "value")
			temp.Set("HTTPVerb", "type")
			return temp, nil
		}
	}
	return nil, utils.NewParseError(p.Pos+1, p.LineNum+1,
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
	return nil, utils.NewParseError(p.Pos+1, p.LineNum+1,
		"Expected 'multipart', but couldn't find", []string{})
}

func (p *Lama2Parser) Form() (*gabs.Container, error) {
	_, e := p.Keyword("form", true, true, true)
	if e == nil {
		temp := gabs.New()
		temp.Set(true, "value")
		temp.Set("Form", "type")
		p.Context["form"] = true
		return temp, nil
	}
	return nil, utils.NewParseError(p.Pos+1, p.LineNum+1,
		"Expected 'form', but couldn't find", []string{})
}

func (p *Lama2Parser) Details() (*gabs.Container, error) {
	res, e := p.Match([]string{"HeaderData", "DataHeader"})
	if e == nil {
		return res, nil
	}

	return nil, utils.NewParseError(p.Pos+1, p.LineNum+1,
		"No details packet found", []string{})
}
