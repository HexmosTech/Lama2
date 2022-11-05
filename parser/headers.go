package parser

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
	"github.com/hexmos/lama2/utils"
)

func (p *Lama2Parser) HeaderData() (*gabs.Container, error) {
	headers, e := p.Match([]string{"Headers"})
	fmt.Println("headers", headers)

	jsond, _ := p.Match([]string{"AnyType"})
	fmt.Println("json", jsond)
	temp := gabs.New()
	if e == nil {
		temp = utils.SetJson(temp, headers, "headers")
	}
	temp = utils.SetJson(temp, jsond, "jsond")

	return temp, nil
}

func (p *Lama2Parser) Headers() (*gabs.Container, error) {
	// headMap := make(map[string]string)
	temp := gabs.New()
	res, e := p.Match([]string{"HeaderPair"})
	if e == nil {
		for k, v := range res.ChildrenMap() {
			temp.Set(v, k)
		}
	} else {
		return nil, utils.NewParseError(p.Pos+1, "Couldn't match headers", []string{})
	}
	for {
		res, e = p.Match([]string{"HeaderPair"})
		if e == nil {
			for k, v := range res.ChildrenMap() {
				temp.Set(v, k)
			}
		} else {
			break
		}

	}
	return temp, nil
}

func (p *Lama2Parser) HeaderPair() (*gabs.Container, error) {
	temp := gabs.New()
	key, e := p.Match([]string{"QuotedString"})
	if e != nil {
		key, e = p.Match([]string{"Unquoted"})
		if e != nil {
			return nil, utils.NewParseError(p.Pos+1, "Couldn't match either quoted or unquoted string", []string{})
		}
	}

	p.Keyword(":", true, true, true)

	value, err := p.Match([]string{"QuotedString"})
	if err != nil {
		value, err = p.Match([]string{"Unquoted"})
		if err != nil {
			return nil, utils.NewParseError(p.Pos+1, "In header value, couldn't get string", []string{})
		}
	}
	valueStr, _ := value.Data().(string)
	keyStr, _ := key.Data().(string)
	temp.Set(valueStr, keyStr)

	return temp, nil
}
