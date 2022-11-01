package parser

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
	"github.com/hexmos/lama2/utils"
)

func (p *Lama2Parser) HeaderData() (*gabs.Container, error) {
	headers, _ := p.Match([]string{"Headers"})
	fmt.Println("headers", headers)
	return headers, nil
}

func (p *Lama2Parser) Headers() (*gabs.Container, error) {
	// headMap := make(map[string]string)
	temp := gabs.New()
	res, e := p.Match([]string{"HeaderPair"})
	if e == nil {
		for k, v := range res.ChildrenMap() {
			temp.Set(v, k)
		}
		return temp, nil
	} else {
		return nil, utils.NewParseError(p.Pos+1, "Couldn't match headers", []string{})
	}
}

func (p *Lama2Parser) HeaderPair() (*gabs.Container, error) {
	temp := gabs.New()
	key, e := p.Match([]string{"QuotedString"})
	if e != nil {
		key, e = p.Match([]string{"Unquoted"})
	}

	p.Keyword(":", true, true, true)

	value, err := p.Match([]string{"QuotedString"})
	if err != nil {
		value, err = p.Match([]string{"Unquoted"})
	}
	valueStr, _ := value.Search("value").Data().(string)
	keyStr, _ := key.Search("value").Data().(string)
	temp.Set(valueStr, keyStr)

	return temp, nil
}
