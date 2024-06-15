//go:build wasm

package preprocess

import (
	"syscall/js"

	"github.com/dop251/goja"
)

func getJsValue(vm *goja.Runtime, name string) goja.Value {
	return js.Global().Get(name)
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
