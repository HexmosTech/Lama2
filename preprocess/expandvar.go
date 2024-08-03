// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// General environment variables.

package preprocess

import (
	"os"
	"strings"

	"github.com/HexmosTech/lama2/utils"
	// "github.com/dop251/goja"
	"github.com/rs/zerolog/log"
	"syscall/js"
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

// isShellSpecialVar reports whether the character identifies a special
// shell variable such as $*.
func isShellSpecialVar(c uint8) bool {
	switch c {
	case '*', '#', '$', '@', '!', '?', '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	}
	return false
}

// isAlphaNum reports whether the byte is an ASCII letter, number, or underscore
func isAlphaNum(c uint8) bool {
	return c == '_' || '0' <= c && c <= '9' || 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
}

// getShellName returns the name that begins the string and the number of bytes
// consumed to extract it. If the name is enclosed in {}, it's part of a ${}
// expansion and two more bytes are needed than the length of the name.
func getShellName(s string) (string, int) {
	switch {
	case s[0] == '{':
		if len(s) > 2 && isShellSpecialVar(s[1]) && s[2] == '}' {
			return s[1:2], 3
		}
		// Scan to closing brace
		for i := 1; i < len(s); i++ {
			if s[i] == '}' {
				if i == 1 {
					return "", 2 // Bad syntax; eat "${}"
				}
				return s[1:i], i + 1
			}
		}
		return "", 1 // Bad syntax; eat "${"
	case isShellSpecialVar(s[0]):
		return s[0:1], 1
	}
	// Scan alphanumerics.
	var i int
	return s[:i], i
}

// ExpandEnv replaces ${var} or $var in the string according to the values
// of the current environment variables. References to undefined
// variables are replaced by the empty string.
func ExpandEnv(s string, vm interface{}) string {
	return Expand(s, vm, getEnvironMap())
}

// func getJsValue(vm interface{}, name string, mapping map[string]string, buf []byte) []byte {
// 	jsVal := vm.(*goja.Runtime).Get(name)
// 	if jsVal != nil {
// 		buf = append(buf, []byte(jsVal.String())...)
// 	} else {
// 		val, ok := mapping[name]
// 		if ok {
// 			buf = append(buf, val...)
// 		} else {
// 			buf = append(buf, ""...)
// 			log.Warn().Str("Couldn't find the variable `"+name+"`,  in both Javascript processor block and environment variables. Replacing with empty string", "").Msg("")
// 		}
// 	}
// 	return buf
// }

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

func getEnvironMap() map[string]string {
	m := make(map[string]string)
	for _, e := range os.Environ() {
		if i := strings.Index(e, "="); i >= 0 {
			m[e[:i]] = e[i+1:]
		}
	}
	return m
}
