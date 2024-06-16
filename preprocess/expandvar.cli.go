//go:build cli

// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// General environment variables.

package preprocess

import (
	"os"
	"strings"

	"github.com/dop251/goja"
	"github.com/rs/zerolog/log"
)

func getJsValue(vm *goja.Runtime, name string, mapping map[string]string, buf []byte) []byte {
	jsVal := vm.Get(name)
	if jsVal != nil {
		buf = append(buf, []byte(jsVal.String())...)
	} else {
		val, ok := mapping[name]
		if ok {
			buf = append(buf, val...)
		} else {
			buf = append(buf, ""...)
			log.Warn().Str("Couldn't find the variable `"+name+"`,  in both Javascript processor block and environment variables. Replacing with empty string", "").Msg("")
		}
	}
	return buf
}

func getEnvironMap() map[string]string {
	m := make(map[string]string)
	for _, e := range os.Environ() {
		if i := strings.Index(e, "="); i >= 0 {
			m[e[:i]] = e[i+1:]
		}
	}
	return m
}
