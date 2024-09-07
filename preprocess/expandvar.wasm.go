//go:build wasm

// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// General environment variables.

package preprocess

import (
	"fmt"

	"github.com/HexmosTech/lama2/utils"
	"github.com/rs/zerolog/log"
)

// Expand replaces ${var} or $var in the string based on the mapping function.
// For example, os.ExpandEnv(s) is equivalent to os.Expand(s, os.Getenv).
func Expand(s string, vm interface{}, mapping map[string]string) string {
	var buf []byte
	// ${} is all ASCII, so bytes are fine for this operation.
	i := 0
	for j := 0; j < len(s); j++ {
		if s[j] == '$' && j+1 < len(s) {
			if buf == nil {
				buf = make([]byte, 0, 2*len(s))
			}
			buf = append(buf, s[i:j]...)
			name, w := getShellName(s[j+1:])
			//nolint:all
			if name == "" && w > 0 {
				// Encountered invalid syntax; eat the
				// characters.
			} else if name == "" {
				// Valid syntax, but $ was not followed by a
				// name. Leave the dollar character untouched.
				buf = append(buf, s[j])
			} else {
				// buf = getJsValue(vm, name, mapping, buf)
				buf = getJsValue(name, mapping, buf)
			}
			j += w
			i = j + 1
		}
	}
	res := ""
	if buf == nil {
		res = s
	} else {
		res = string(buf) + s[i:]
	}
	res2 := utils.RemoveUnquotedValueMarker(res)
	return res2
}

func getJsValue(name string, mapping map[string]string, buf []byte) []byte {
	log.Debug().Str("WW Getting from worker", name).Msg("")
	jsVal := GetFromWorker(name)
	fmt.Println("WW: getjsvalue", name, jsVal)
	log.Debug().Str("WW JavaScript Variable Value:", jsVal).Msg("")
	if jsVal != "" {
		buf = append(buf, []byte(jsVal)...)
	} else {
		val, ok := mapping[name]
		if ok {
			buf = append(buf, val...)
		} else {
			buf = append(buf, ""...)
			log.Warn().Str("Couldn't find the variable `"+name+"`, in both Javascript processor block and environment variables. Replacing with empty string", "").Msg("")
		}
	}
	return buf
}
