package parser

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
	"github.com/hexmos/lama2/utils"
)

func (p *Lama2Parser) AnyType() (*gabs.Container, error) {
	r, e := p.Match([]string{"ComplexType", "PrimitiveType"})
	return r, e
}

func (p *Lama2Parser) ComplexType() (*gabs.Container, error) {
	r, e := p.Match([]string{"List", "Map"})
	return r, e
}

func (p *Lama2Parser) PrimitiveType() (*gabs.Container, error) {
	r, e := p.Match([]string{"Null", "Boolean", "QuotedString", "Number"})
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

func (p *Lama2Parser) List() (*gabs.Container, error) {
	temp := gabs.New()
	temp.Set([]interface{}{})
	_, e1 := p.Keyword("[", true, true, true)
	if e1 != nil {
		return nil, e1
	}
	// comma := false
	for {
		item, e2 := p.Match([]string{"AnyType"})
		fmt.Println("In List: got = ", item, string(p.Text[p.Pos:p.Pos+10]))
		if e2 != nil {
			break
		}

		temp.ArrayAppend(item)

		_, e3 := p.Keyword(",", true, true, true)
		if e3 != nil {
			break
		}
		/*
			if e3 != nil {
				comma = false
				break
			} else {
				comma = true
			}
		*/
	}
	_, e4 := p.Keyword("]", true, true, true)
	if e4 != nil {
		return nil, e4
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

func (p *Lama2Parser) Boolean() (*gabs.Container, error) {
	temp := gabs.New()
	_, e1 := p.Keyword("true", true, true, false)
	if e1 != nil {
		_, e2 := p.Keyword("false", true, true, false)
		if e2 != nil {
			return nil, utils.NewParseError(p.Pos+1, "Couldn't find boolean", []string{})
		} else {
			temp.Set(false)
		}
	} else {
		temp.Set(true)
	}

	return temp, nil
}

func (p *Lama2Parser) Null() (*gabs.Container, error) {
	temp := gabs.New()
	_, e1 := p.Keyword("null", true, true, false)
	if e1 != nil {
		return nil, utils.NewParseError(p.Pos+1, "Couldn't find null", []string{})
	}
	temp.Set(nil)
	return temp, nil
}
