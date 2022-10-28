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
