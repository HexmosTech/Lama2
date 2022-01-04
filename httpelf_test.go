package main

import (
	"strings"
	"testing"
)

func TestBasicGet(t *testing.T) {
	res := processHttpFile("./sample_http/httpbin/get.http")
	if !(strings.Contains(res, "world") && strings.Contains(res, "hello")) {
		t.Errorf("Expected the result arguments to contain \"hello\" : \"world\"")
	}
}
