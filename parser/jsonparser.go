package parser

import (
	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/utils"
	"github.com/rs/zerolog/log"
)

// AnyType is the top-most element of a JSON structure
// It consists of Complex and Primitive Types
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

// CustomPairMerge uses a gabs feature to deal with merge conflicts.
// More here: https://github.com/HexmosTech/gabs/blob/master/gabs.go#L511
func CustomPairMerge(_, source interface{}) interface{} {
	return source
}

// Map is a slightly lenient version of standard JSON
// map. In Lama2 Map, it is OK to have a trailing
// comma after the last element (whereas in strict JSON,
// it is not OK to have trailing comma)
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
		log.Trace().Str("Pair result", item.String()).Msg("")
		temp.MergeFn(item, CustomPairMerge)

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

// List is a slightly lenient version of standard JSON
// list. In Lama2 List, it is OK to have a trailing
// comma after the last element (whereas in strict JSON,
// it is not OK to have trailing comma)
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
		if e2 != nil {
			break
		}

		temp.ArrayAppend(item)

		_, e3 := p.Keyword(",", true, true, true)
		if e3 != nil {
			break
		}
		/* Enable this section to get strict JSON checking
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
	key, e := p.Match([]string{"QuotedString"})
	if e != nil {
		return nil, e
	}
	_, e = p.Keyword(":", true, true, true)
	if e != nil {
		return nil, e
	}

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
			return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "Couldn't find boolean", []string{})
		}
		temp.Set(false)
	} else {
		temp.Set(true)
	}

	return temp, nil
}

func (p *Lama2Parser) Null() (*gabs.Container, error) {
	temp := gabs.New()
	_, e1 := p.Keyword("null", true, true, false)
	if e1 != nil {
		return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "Couldn't find null", []string{})
	}
	temp.Set(nil)
	return temp, nil
}
