package parser

import (
	"fmt"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/utils"
)

func (p *Lama2Parser) HeaderData() (*gabs.Container, error) {
	headers, e := p.Match([]string{"Headers"})
	fmt.Println("headers", headers)
	if e != nil {
		return nil, e
	}

	jsond, e2 := p.Match([]string{"DataInput"})
	fmt.Println("json", jsond)
	temp := gabs.New()
	temp = utils.SetJson(temp, headers, "headers")
	if e2 == nil {
		temp = utils.SetJson(temp, jsond, "ip_data")
	}
	return temp, nil
}

func (p *Lama2Parser) DataHeader() (*gabs.Container, error) {
	jsond, e := p.Match([]string{"DataInput"})
	if e != nil {
		return nil, e
	}

	headers, e2 := p.Match([]string{"Headers"})
	fmt.Println("json", jsond)
	temp := gabs.New()
	temp = utils.SetJson(temp, jsond, "ip_data")
	if e2 == nil {
		temp = utils.SetJson(temp, headers, "headers")
	}

	return temp, nil
}

func (p *Lama2Parser) DataInput() (*gabs.Container, error) {
	s, e1 := p.Match([]string{"VarJSON"})
	if e1 != nil {
		s, e1 = p.Match([]string{"AnyType"})
	}
	return s, e1
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
		return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "Couldn't match headers", []string{})
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
			return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "Couldn't match either quoted or unquoted string", []string{})
		}
	}

	_, e2 := p.Keyword(":", true, true, true)
	if e2 != nil {
		return nil, e2
	}

	value, err := p.Match([]string{"QuotedString"})
	if err != nil {
		value, err = p.Match([]string{"Unquoted"})
		if err != nil {
			return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "In header value, couldn't get string", []string{})
		}
	}
	valueStr, _ := value.Data().(string)
	keyStr, _ := key.Data().(string)
	temp.Set(valueStr, keyStr)

	fmt.Println("header pair", temp)

	return temp, nil
}
