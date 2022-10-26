package main

import (
	"strings"
	"testing"
)

func TestBasicGet(t *testing.T) {
	o := Opts{}
	o.Nocolor = false
	res := processHttpFile("./sample_http/httpbin/get.http", o)
	if !(strings.Contains(res, "world") && strings.Contains(res, "hello")) {
		t.Errorf("Expected the result arguments to contain \"hello\" : \"world\"")
	}
}
