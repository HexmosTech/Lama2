package main

import (
	"regexp"

	"github.com/dlclark/regexp2"
)

type ConstantsStruct struct {
	r_json_obj  *regexp.Regexp
	r_httpv     *regexp.Regexp
	r_url       *regexp.Regexp
	r_linetype  *regexp2.Regexp
	r_multipart *regexp.Regexp
}

func Constants() ConstantsStruct {
	c := ConstantsStruct{}
	c.r_json_obj = regexp.MustCompile(`(?smi)^([{\[].*[}\]])$`)
	c.r_httpv = regexp.MustCompile(`(?smi)^(GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH)\s*$`)
	c.r_url = regexp.MustCompile(`(?smi)^https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.?[a-zA-Z0-9()]{1,6}\b?([-a-zA-Z0-9()@:%_\+.~#?&\//=]*).*?$`)
	c.r_linetype = regexp2.MustCompile(`(?sm)^(?![http|#|{])^("([^"]*)"|'([^']*)'|(.*?))(?P<type>[:@=])("([^"]*)"|'([^']*)'|(.*?))\s*$`, regexp2.RE2)
	c.r_multipart = regexp.MustCompile(`(?smi)^MULTIPART\s*$`)
	return c
}
