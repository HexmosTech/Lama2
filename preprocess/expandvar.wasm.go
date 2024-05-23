//go:build wasm

package preprocess

import (

	"syscall/js"

	"github.com/HexmosTech/lama2/utils"
)

// Expand replaces ${var} or $var in the string based on the mapping function.
// For example, os.ExpandEnv(s) is equivalent to os.Expand(s, os.Getenv).
func Expand(s string, mapping map[string]string) string {
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
				// jsVal := vm.Get(name)
				jsVal := js.Global().Get(name)

				if jsVal.Truthy() {
					buf = append(buf, []byte(jsVal.String())...)
				} else {
					val, ok := mapping[name]
					if ok {
						buf = append(buf, val...)
					} else {
						buf = append(buf, ""...)
					}
				}
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

func getEnvironMap() map[string]string {
	result := make(map[string]string)

	keysArray :=  js.Global().Get("Object").Call("keys", js.Global())
	keys := make([]string, keysArray.Length())
	for i := 0; i < keysArray.Length(); i++ {
		keys[i] = keysArray.Index(i).String()
	}
	for _, name := range keys {
		if value := js.Global().Get(name); value.Truthy() {
			result[name] = value.String()
		}
	}


	return result
}

// ExpandEnv replaces ${var} or $var in the string according to the values
// of the current environment variables. References to undefined
// variables are replaced by the empty string.
func ExpandEnv(s string) string {
	return Expand(s, getEnvironMap())
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
