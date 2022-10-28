package tests

import (
	"testing"

	"fmt"

	"github.com/hexmos/lama2/parser"
)

func TestParser(t *testing.T) {
	p := parser.Parser{}
	got := p.Parse("GET http://google.com")
	want := "blah"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestLama2Parser(t *testing.T) {
	p := parser.NewLama2Parser()
	got := p.Parse("GET http://google.com")
	want := "blah"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestCharFunc(t *testing.T) {
	p := parser.NewLama2Parser()
	p.SetText("GET http://google.com")

	for true {
		r, e := p.Char()
		if e == nil {
			fmt.Printf("%c", r)
		} else {
			fmt.Println("End of input")
			break
		}
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
	got, e = p.CharClass("x")
	if e == nil {
		t.Errorf("Expected an error")
	} else {
		fmt.Printf("Got the error: %s", e)
	}
	got, e = p.CharClass("A-Z0-9&")
}
