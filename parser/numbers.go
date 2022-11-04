package parser

import (
	"strconv"
	"strings"

	"github.com/Jeffail/gabs/v2"
	"github.com/hexmos/lama2/utils"
)

func (p *Lama2Parser) Number() (*gabs.Container, error) {
	intPart, e := p.Match([]string{"Integer"})
	return intPart, e
}

func (p *Lama2Parser) Integer() (*gabs.Container, error) {
	s, e := p.Match([]string{"IntegerRule4",
		"IntegerRule3", "IntegerRule2", "IntegerRule1"})
	temp := gabs.New()
	if e == nil {
		i, err := strconv.Atoi(s.Search("value").Data().(string))
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
	s, e := p.Match([]string{"Digit"})
	return s, e
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
		rVal := r.Search("value").Data().(string)
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
		rVal := r.Search("value").Data().(string)
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
	s, e := p.CharClass("+-")
	temp := gabs.New()
	if e == nil {
		temp.Set(string(s), "value")
		return temp, e
	} else {
		return nil, e
	}
}
