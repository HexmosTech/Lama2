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
)

func getJsValue(vm *goja.Runtime, name string) goja.Value {
	return vm.Get(name)
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
