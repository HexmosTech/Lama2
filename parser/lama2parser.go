package parser

import (
	"errors"
	"strings"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/utils"
	"github.com/rs/zerolog/log"
)

type Lama2Parser struct {
	*Parser
	Context map[string]bool
}

// NewLama2Parser creates a new Lama2Parser
// and initializes it properly
func NewLama2Parser() *Lama2Parser {
	p := &Lama2Parser{&Parser{}, make(map[string]bool)}
	p.Pm = p
	p.Init()
	return p
}

// Start primarily calls the Lama2File method
func (p *Lama2Parser) Start() (*gabs.Container, error) {
	log.Trace().Msg("Within the Start function")
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
	log.Trace().Msg("Within Lama2File")
	temp := gabs.New()
	tempArr, e1 := temp.Array()
	if e1 != nil {
		return nil, errors.New("Couldn't create Array for the parsed data")
	}
	// optionally match processor
	res2, e2 := p.Match([]string{"Processor"})
	if e2 == nil {
		tempArr.ArrayAppend(res2)
	}

	// match requester
	res3, e3 := p.Match([]string{"Requester"})
	if e3 != nil {
		return nil, e3
	}

	tempArr.ArrayAppend(res3)

	// until file is done:
	var res4, res5 *gabs.Container
	var e4, e5 error
	for p.Pos < p.TotalLen {
		// match processor
		res4, e4 = p.Match([]string{"Processor"})
		if e4 == nil {
			tempArr.ArrayAppend(res4)
		} else {
			break
		}

		// match requester
		res5, e5 = p.Match([]string{"Requester"})
		if e5 == nil {
			tempArr.ArrayAppend(res5)
		}
	}
	return tempArr, nil
}

func (p *Lama2Parser) Processor() (*gabs.Container, error) {
	log.Trace().Msg("Within Processor")
	temp := gabs.New()
	return temp, nil
}

func (p *Lama2Parser) Separator() (*gabs.Container, error) {
	log.Trace().Msg("Within Separator")
	temp := gabs.New()
	return temp, nil
}

// Requester applies the rule:
// HTTPVerb Multipart? TheURL Details?
func (p *Lama2Parser) Requester() (*gabs.Container, error) {
	log.Trace().Msg("Within Requester")
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
	kw, e := p.Keyword("http", true, false, true)
	if e == nil {
		res = append(res, string(kw))
	} else {
		return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "Couldn't find URL (starting with http(s)", []string{})
	}

	_, e = p.CharClass("s")
	if e == nil {
		res = append(res, "s")
	}

	_, e = p.Keyword("://", false, false, true)
	if e == nil {
		res = append(res, "://")
	} else {
		return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "Couldn't find URL (starting with http(s)://", []string{})
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
	temp.Set("TheURL", "type")
	return temp, nil
}

func (p *Lama2Parser) HTTPVerb() (*gabs.Container, error) {
	log.Trace().Msg("Within HTTPVerb")
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

func (p *Lama2Parser) Details() (*gabs.Container, error) {
	res, e := p.Match([]string{"HeaderData", "DataHeader"})
	if e == nil {
		return res, nil
	}

	return nil, utils.NewParseError(p.Pos+1, p.LineNum+1,
		"No details packet found", []string{})
}
