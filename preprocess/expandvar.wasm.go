//go:build wasm

// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// General environment variables.

package preprocess

import (
	"github.com/HexmosTech/lama2/utils"
	"github.com/rs/zerolog/log"
	"syscall/js"
)

var worker js.Value

func initWebWorker() js.Value {
    if worker.IsUndefined() {
        worker = js.Global().Get("Worker").New("worker.js")
        worker.Call("addEventListener", "message", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
            result := args[0].String()
			log.Debug().Str("CResult from web worker:", result).Msg("")
            return nil
        }))
    }
    return worker
}


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
    worker := initWebWorker()
    
    // Use a channel to synchronize with the worker response
    responseChan := make(chan string)
    
    // Setup a listener to capture the worker's response
    js.Global().Get("addEventListener").Invoke("message", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        result := args[0].String()
        responseChan <- result
        return nil
    }))

    // Send the request to the worker
    worker.Call("postMessage", name)

    // Wait for the worker response
    jsVal := <-responseChan
    
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
