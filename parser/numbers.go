package parser

import (
	"math"
	"strconv"
	"strings"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/utils"
)

func (p *Lama2Parser) Number() (*gabs.Container, error) {
	intPart, e1 := p.Match([]string{"Integer"})
	fracPart, e2 := p.Match([]string{"Fraction"})
	expPart, e3 := p.Match([]string{"Exponent"})

	if e1 != nil {
		return nil, e1
	}
	i := intPart.Data().(int)
	multiplier := 1
	if i < 0 {
		multiplier = -1
	}
	f := float64(0)
	e := float64(0)

	if e2 == nil {
		f = fracPart.Data().(float64)
	}

	if e3 == nil {
		e = expPart.Data().(float64)
	}

	temp := gabs.New()
	if f == 0 && e == 0 {
		temp.Set(i)
	} else if e == 0 {
		temp.Set(float64(i) + (float64(multiplier) * f))
	} else {
		temp.Set(
			(float64(i) +
				(float64(multiplier) * f)) * math.Pow(10, e))
	}
	return temp, nil
}

func (p *Lama2Parser) Exponent() (*gabs.Container, error) {
	_, e := p.CharClass("eE")
	if e != nil {
		return nil, e
	}

	sign, e2 := p.Match([]string{"Sign"})
	multiplier := 1
	if e2 == nil && sign.S("value").Data().(string) == "-" {
		multiplier = -1
	}

	d, e3 := p.Match([]string{"Digits"})
	if e3 != nil {
		return nil, e
	}
	dVal := d.S("value").Data().(string)
	dInt, _ := strconv.Atoi(dVal)

	res := float64(multiplier) * float64(dInt)
	temp := gabs.New()
	temp.Set(res)
	return temp, nil
}

func (p *Lama2Parser) Fraction() (*gabs.Container, error) {
	s, e := p.Match([]string{"FractionRule1"})
	if e != nil {
		return nil, e
	}
	f, err := strconv.ParseFloat(s.Data().(string), 64)
	if err != nil {
		return nil, e
	}
	temp := gabs.New()
	temp.Set(f)
	return temp, nil
}

func (p *Lama2Parser) FractionRule1() (*gabs.Container, error) {
	s := make([]string, 0)
	r, e := p.CharClass(".")
	if e != nil {
		return nil, e
	}
	s = append(s, string(r))
	r2, e2 := p.Match([]string{"Digits"})
	if e2 != nil {
		return nil, e
	}
	digs := r2.Search("value").Data().(string)
	s = append(s, digs)

	temp := gabs.New()
	temp.Set(strings.Join(s, ""))
	return temp, nil

}

func (p *Lama2Parser) Integer() (*gabs.Container, error) {
	s, e := p.Match([]string{"IntegerRule4",
		"IntegerRule3", "IntegerRule2", "IntegerRule1"})
	temp := gabs.New()
	if e == nil {
		i, err := strconv.Atoi(s.Data().(string))
		if err == nil {
			temp.Set(i)
			return temp, nil
		}
		return nil, utils.NewParseError(p.Pos+1, "Not able to convert to integer as expected", []string{})
	} else {
		return s, e
	}
}

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

func (p *Lama2Parser) IntegerRule3() (*gabs.Container, error) {
	s := make([]string, 0)
	r, e := p.Match([]string{"Sign"})
	if e == nil {
		rVal := r.Search("value").Data().(string)
		if rVal == "-" {
			s = append(s, rVal)
		} else {
			return nil, utils.NewParseError(p.Pos+1, "Expected negative sign",
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

func (p *Lama2Parser) IntegerRule4() (*gabs.Container, error) {
	s := make([]string, 0)
	r, e := p.Match([]string{"Sign"})
	if e == nil {
		rVal := r.Search("value").Data().(string)
		if rVal == "-" {
			s = append(s, rVal)
		} else {
			return nil, utils.NewParseError(p.Pos+1, "Expected negative sign",
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
	} else {
		return nil, e
	}
}

func (p *Lama2Parser) Sign() (*gabs.Container, error) {
	s, e := p.CharClass("-")
	temp := gabs.New()
	if e == nil {
		temp.Set(string(s), "value")
		return temp, e
	} else {
		return nil, e
	}
}
