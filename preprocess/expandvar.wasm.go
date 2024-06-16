//go:build wasm

package preprocess

import (
	"syscall/js"

	"github.com/dop251/goja"
)

func getJsValue(vm *goja.Runtime, name string, mapping map[string]string, buf []byte) []byte {
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

	return buf
}

func getEnvironMap() map[string]string {
	result := make(map[string]string)

	keysArray := js.Global().Get("Object").Call("keys", js.Global())
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
