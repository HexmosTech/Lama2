package main

import (
	// "fmt"
	"syscall/js"

	controller "github.com/HexmosTech/lama2/controller"
)

var version string

func init() {
	// Initialize version if not set
	if len(version) == 0 {
		version = "vUnset"
	}
}

func main() {
	// Set up the WebAssembly function
	js.Global().Set("lama2Wasm", js.FuncOf(lama2Wasm))
	js.Global().Set("goWebRequestFunc", wasmLamaPromise())
	select {}
}

func lama2Wasm(this js.Value, i []js.Value) interface{} {
	go func() {
		// controller.ProcessWasmInput("POST\nhttps://httpbin.org/post\n{'a': 'b','c': 'd'}")
		// controller.ProcessWasmInput("GET\nhttps://httpbin.org/get")
		// controller.ProcessWasmInput("POST\nhttps://httpbin.org/post\na='b'\n 'c'=d \n X-Parse-Application-Id:'helloworld'\nX-Parse-REST-API-Key:'byeworld'\n")
		// controller.ProcessWasmInput("url = 'http://google.com'\nREMOTE_COORD = 'https://httpbin.org'\n---\nPOST\n${REMOTE_COORD}/anything\n{\n    'username': 'admin',\n    'password': 'Password@123',\n    'from': '${LOCAL_COORD}/anything',\n    'url': '${url}',\n    'Token': 'MySuperSecretToken'\n}\n")
		// controller.ProcessWasmInput("url = 'http://google.com'\nREMOTE_COORD = 'https://httpbin.org'\n---\nPOST\n${REMOTE_COORD}/anything\n{\n    'username': 'admin',\n    'password': 'Password@123',\n    'from': '${LOCAL_COORD}/anything',\n    'url': '${url}',\n    'Token': 'MySuperSecretToken'\n}\n---\nconsole.log('@@Result', result)\nlet TOKEN = result['json']['Token']\nconsole.log(TOKEN)\n")
		controller.ProcessWasmInput("url = 'http://google.com'\nREMOTE_COORD = 'https://httpbin.org'\n---\nPOST\n${REMOTE_COORD}/anything\n{\n    'username': 'admin',\n    'password': 'Password@123',\n    'from': '${LOCAL_COORD}/anything',\n    'url': '${url}',\n    'Token': 'MySuperSecretToken'\n}\n---\nconsole.log('@@Result', result)\nlet TOKEN = result['json']['Token']\nconsole.log(TOKEN)\n---\nGET\n${REMOTE_COORD}/bearer\nAuthorization: 'Bearer ${TOKEN}'\n\n")
	}()

	return js.ValueOf("data from api:")
}

func wasmLamaPromise() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println("Args:", args[0].String())
		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]

			go func() {
				result := controller.ProcessWasmInput(args[0].String())
				resolve.Invoke(response)
			}()
			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}
