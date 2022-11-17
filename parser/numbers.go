package parser

import (
	"encoding/json"
	"strings"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/utils"
)

// A Number consists of a mandatory integer part,
// and optional Fraction and Exponent parts. The Number
// method "collects" these three elements, converts them
// into a json.Number() type, and finally returns the
// Number wrapped within a gabs Container
func (p *Lama2Parser) Number() (*gabs.Container, error) {
	intPart, e1 := p.Match([]string{"Integer"})
	fracPart, e2 := p.Match([]string{"Fraction"})
	expPart, e3 := p.Match([]string{"Exponent"})

	if e1 != nil {
		return nil, e1
	}
	f := ""
	e := ""

	if e2 == nil {
		f = fracPart.Data().(string)
	}

	if e3 == nil {
		e = expPart.Data().(string)
	}

	temp := gabs.New()
	if f == "" && e == "" {
		num := json.Number(intPart.Data().(string))
		temp.Set(num)
	} else if e == "" {
		iStr := intPart.Data().(string)
		fStr := fracPart.Data().(string)
		temp.Set(json.Number(iStr + fStr))
	} else {
		iStr := intPart.Data().(string)
		temp.Set(json.Number(iStr + f + e))
	}
	return temp, nil
}

// An Exponent consists of mandatory 'e' or 'E',
// optional Sign, followed by Digits
func (p *Lama2Parser) Exponent() (*gabs.Container, error) {
	pows, e := p.CharClass("eE")
	if e != nil {
		return nil, e
	}

	sign, e2 := p.Match([]string{"Sign"})
	if e2 != nil {
		sign = gabs.New()
		sign.Set("", "value")
	}

	d, e3 := p.Match([]string{"Digits"})
	if e3 != nil {
		return nil, e3
	}
	dVal := d.S("value").Data().(string)

	temp := gabs.New()
	temp.Set(string(pows) + sign.S("value").Data().(string) + dVal)
	return temp, nil
}

func (p *Lama2Parser) Fraction() (*gabs.Container, error) {
	s, e := p.Match([]string{"FractionRule1"})
	if e != nil {
		return nil, e
	}
	temp := gabs.New()
	temp.Set(s.Data().(string))
	return temp, nil
}

// A Fraction consists of mandatory "." (dot),
// followed by Digits.
func (p *Lama2Parser) FractionRule1() (*gabs.Container, error) {
	s := make([]string, 0)
	r, e := p.CharClass(".")
	if e != nil {
		return nil, e
	}
	s = append(s, string(r))
	r2, e2 := p.Match([]string{"Digits"})
	if e2 != nil {
		return nil, e2
	}
	digs := r2.Search("value").Data().(string)
	s = append(s, digs)

	temp := gabs.New()
	temp.Set(strings.Join(s, ""))
	return temp, nil
}

func (p *Lama2Parser) Integer() (*gabs.Container, error) {
	s, e := p.Match([]string{
		"IntegerRule4",
		"IntegerRule3", "IntegerRule2", "IntegerRule1",
	})
	temp := gabs.New()
	if e == nil {
		temp.Set(s.Data().(string))
		return temp, nil
	}
	return s, e
}

// InterRule1 matches a Digit
func (p *Lama2Parser) IntegerRule1() (*gabs.Container, error) {
	s := make([]string, 0)
	r, e := p.Match([]string{"Digit"})
	if e == nil {
		s = append(s, r.Search("value").Data().(string))
	} else {
		return nil, e
	}
	temp := gabs.New()
	temp.Set(strings.Join(s, ""))
	return temp, e
}

// IntegerRule2 matches 1-9 mandatorily, and then
// tries to follow it with Digits
func (p *Lama2Parser) IntegerRule2() (*gabs.Container, error) {
	s := make([]string, 0)
	iOneNine, e := p.Match([]string{"OneNine"})
	if e == nil {
		s = append(s, iOneNine.Search("value").Data().(string))
	} else {
		return nil, e
	}

	iDigits, e := p.Match([]string{"Digits"})
	if e == nil {
		s = append(s, iDigits.Search("value").Data().(string))
	}

	temp := gabs.New()
	temp.Set(strings.Join(s, ""))
	return temp, nil
}

// IntegerRule3 starts with a mandatory Sign, and
// follows with IntegerRule1 (Digit)
func (p *Lama2Parser) IntegerRule3() (*gabs.Container, error) {
	s := make([]string, 0)
	r, e := p.Match([]string{"Sign"})
	if e == nil {
		rVal := r.Search("value").Data().(string)
		if rVal == "-" {
			s = append(s, rVal)
		} else {
			return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "Expected negative sign",
				[]string{})
		}
	} else {
		return nil, e
	}

	r, e = p.Match([]string{"IntegerRule1"})
	if e == nil {
		rVal := r.Data().(string)
		s = append(s, rVal)

	} else {
		return nil, e
	}

	temp := gabs.New()
	temp.Set(strings.Join(s, ""))
	return temp, nil
}

// IntegerRule4 starts with a mandatory Sign, and
// follows with IntegerRule2
func (p *Lama2Parser) IntegerRule4() (*gabs.Container, error) {
	s := make([]string, 0)
	r, e := p.Match([]string{"Sign"})
	if e == nil {
		rVal := r.Search("value").Data().(string)
		if rVal == "-" {
			s = append(s, rVal)
		} else {
			return nil, utils.NewParseError(p.Pos+1, p.LineNum+1, "Expected negative sign",
				[]string{})
		}
	} else {
		return nil, e
	}

	r, e = p.Match([]string{"IntegerRule2"})
	if e == nil {
		rVal := r.Data().(string)
		s = append(s, rVal)
	} else {
		return nil, e
	}

	temp := gabs.New()
	temp.Set(strings.Join(s, ""))
	return temp, nil
}

func (p *Lama2Parser) Digits() (*gabs.Container, error) {
	s := make([]string, 0)
	r, e := p.Match([]string{"Digit"})
	if e == nil {
		s = append(s, r.Search("value").Data().(string))
	} else {
		return nil, e
	}
	for {
		r, e = p.Match([]string{"Digit"})
		if e == nil {
			s = append(s, r.Search("value").Data().(string))
		} else {
			break
		}
	}
	temp := gabs.New()
	temp.Set(strings.Join(s, ""), "value")
	return temp, nil
}

func (p *Lama2Parser) Digit() (*gabs.Container, error) {
	s, e := p.CharClass("0")
	temp := gabs.New()
	if e == nil {
		temp.Set(string(s), "value")
		return temp, nil
	}

	d, e2 := p.Match([]string{"OneNine"})
	return d, e2
}

func (p *Lama2Parser) OneNine() (*gabs.Container, error) {
	s, e := p.CharClass("1-9")
	temp := gabs.New()
	if e == nil {
		temp.Set(string(s), "value")
		return temp, nil
	}
	return nil, e
}

func (p *Lama2Parser) Sign() (*gabs.Container, error) {
	s, e := p.CharClass("-+")
	temp := gabs.New()
	if e == nil {
		temp.Set(string(s), "value")
		return temp, e
	}
	return nil, e
}
