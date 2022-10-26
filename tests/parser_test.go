package tests

import (
	"testing"

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
