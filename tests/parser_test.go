package tests

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/HexmosTech/lama2/parser"
)

func TestLama2Parser(t *testing.T) {
	p := parser.NewLama2Parser()
	got, _ := p.Parse("GET http://google.com")
	fmt.Println(got)
	/*
		want := "blah"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	*/
}

func TestCharFunc(t *testing.T) {
	p := parser.NewLama2Parser()
	p.SetText("GET http://google.com")

	for {
		r, e := p.Char()
		if e == nil {
			fmt.Printf("%c", r)
		} else {
			fmt.Println("End of input")
			break
		}
	}
}

func TestKeywordMatch(t *testing.T) {
	p := parser.NewLama2Parser()
	p.SetText("GET http://google.com")
	got, e := p.Keyword("GET", false, false, false)
	want := []rune("GET")
	if e == nil {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q, wanted %q", got, want)
		}
		fmt.Printf("%c\n", got)
		fmt.Println("===")
	} else {
		t.Errorf("Error not expected")
		fmt.Println(e)
	}
}

func TestKeywordWhiteSpaceMatch(t *testing.T) {
	p := parser.NewLama2Parser()
	p.SetText("    GET      http://google.com")
	got, e := p.Keyword("GET", true, true, false)
	want := []rune("GET")
	if e == nil {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q, wanted %q", got, want)
		}
		fmt.Printf("%c\n", got)
		fmt.Println("===")
	} else {
		t.Errorf("Error not expected")
		fmt.Println(e)
	}
}

func TestKeywordCaseInsensitiveMatch(t *testing.T) {
	p := parser.NewLama2Parser()
	p.SetText("    GET      http://google.com")
	got, e := p.Keyword("get", true, true, true)
	want := []rune("get")
	if e == nil {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q, wanted %q", got, want)
		}
		fmt.Printf("%c\n", got)
		fmt.Println("===")
	} else {
		t.Errorf("Error not expected")
		fmt.Println(e)
	}
}

func TestCharClassFunc(t *testing.T) {
	p := parser.NewLama2Parser()
	p.SetText("GET http://google.com")
	got, e := p.CharClass("A-Z0-9&")
	want := 'G'
	if e == nil {
		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
		fmt.Printf("%c\n", got)
		fmt.Println("===")
	} else {
		t.Errorf("Error not expected")
		fmt.Println(e)
	}
	_, e = p.CharClass("x")
	if e == nil || !strings.Contains(e.Error(), "no match") {
		t.Fatalf("Expected a no match error")
	}
}

func TestMatch(t *testing.T) {
	p := parser.NewLama2Parser()
	p.SetText("GET http://google.com")
	got, e := p.Match([]string{"HTTPFile"})
	if e == nil {
		fmt.Println(got)
		fmt.Println("===")
	} else {
		t.Errorf("Error not expected")
		fmt.Println(e)
	}
}

func TestFailMatch(t *testing.T) {
	p := parser.NewLama2Parser()
	p.SetText("BLAH http://google.com")
	got, e := p.Match([]string{"HTTPFile"})
	if e == nil {
		fmt.Println(got)
		t.Errorf("Expected an error!")
	}
	fmt.Println("===")
}

func TestLama2Start(t *testing.T) {
	p := parser.NewLama2Parser()
	p.SetText("GET http://google.com")
	got, e := p.Start()
	fmt.Println("got ", got)
	fmt.Println("e ", e)
	// want := 'G'
	// if e == nil {
	// 	if got != want {
	// 		t.Errorf("got %q, wanted %q", got, want)
	// 	}
	// 	fmt.Printf("%c\n", got)
	// 	fmt.Println("===")
	// } else {
	// 	t.Errorf("Error not expected")
	// 	fmt.Println(e)
	// }
}
