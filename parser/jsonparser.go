package parser

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

func (p *Lama2Parser) AnyType() (*gabs.Container, error) {
	r, e := p.Match([]string{"ComplexType", "PrimitiveType"})
	return r, e
}

func (p *Lama2Parser) ComplexType() (*gabs.Container, error) {
	r, e := p.Match([]string{"Map"})
	return r, e
}

func (p *Lama2Parser) PrimitiveType() (*gabs.Container, error) {
	r, e := p.Match([]string{"QuotedString"})
	return r, e
}

func (p *Lama2Parser) Map() (*gabs.Container, error) {
	temp := gabs.New()
	_, e := p.Keyword("{", true, true, true)
	if e != nil {
		return nil, e
	}
	for {
		item, err := p.Match([]string{"Pair"})
		if err != nil {
			break
		}
		fmt.Println("PairResult", item)
		temp.Merge(item)

		_, e = p.Keyword(",", true, true, true)
		if e != nil {
			break
		}
	}

	_, e = p.Keyword("}", true, true, true)
	if e != nil {
		return nil, e
	}

	return temp, nil

}

func (p *Lama2Parser) Pair() (*gabs.Container, error) {
	fmt.Println("Within Pair")
	key, e := p.Match([]string{"QuotedString"})
	if e != nil {
		return nil, e
	}
	_, e = p.Keyword(":", true, true, true)
	if e != nil {
		return nil, e
	}

	fmt.Println("Trying to match AnyType")
	value, err := p.Match([]string{"AnyType"})
	if err != nil {
		return nil, e
	}

	temp := gabs.New()
	temp.Set(value, key.Data().(string))

	return temp, nil
}
