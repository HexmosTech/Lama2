package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"syscall/js"
	"time"

	contoller "github.com/HexmosTech/lama2/controller"
)

var version string

func main() {
	if len(version) == 0 {
		version = "vUnset"
	}
	// contoller.Process(version)
	// apiContent := "GET\nhttps://httpbin.org/get"
	// contoller.ProcessWebInput(apiContent, version)
	js.Global().Set("wasmLamaRequest", js.FuncOf(parseWasmRequest))
	js.Global().Set("helloWasm", js.FuncOf(hello))
	js.Global().Set("GoWebRequestFunc", GoWebRequestFunc())
	select {}
}

func parseWasmRequest(this js.Value, args []js.Value) interface{} {
	responseMap := contoller.ProcessWasmInput(args[0].String())
	jsonData, err := json.Marshal(responseMap)
	if err != nil {
		panic(err)
	}
	// contoller.ProcessWasmInput(args[0].String())
	// return js.ValueOf(contoller.ProcessWasmInput(args[0].String()))
	return js.ValueOf(string(jsonData))

}

func hello(this js.Value, p []js.Value) interface{} {
	return js.ValueOf("Hello from Go WASM!")
}

func MyGoFunc() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// Handler for the Promise: this is a JS function
		// It receives two arguments, which are JS functions themselves: resolve and reject
		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			// Commented out because this Promise never fails
			//reject := args[1]

			// Now that we have a way to return the response to JS, spawn a goroutine
			// This way, we don't block the event loop and avoid a deadlock
			go func() {
				// Block the goroutine for 3 seconds
				time.Sleep(3 * time.Second)
				// Resolve the Promise, passing anything back to JavaScript
				// This is done by invoking the "resolve" function passed to the handler
				resolve.Invoke("Trentatré Trentini entrarono a Trento, tutti e trentatré trotterellando")
			}()

			// The handler of a Promise doesn't return any value
			return nil
		})

		// Create and return the Promise object
		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}

func GoWebRequestFunc() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// Get the URL as argument
		// args[0] is a js.Value, so we need to get a string out of it
		requestUrl := args[0].String()
		requestMethod := args[1].String()

		// Handler for the Promise
		// We need to return a Promise because HTTP requests are blocking in Go
		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			reject := args[1]

			// Run this code asynchronously
			go func() {
				// Make the HTTP request
				// res, err := http.DefaultClient.Get(requestUrl)
				req, err := http.NewRequest(requestMethod, requestUrl, nil)
				if err != nil {
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New(err.Error())
					reject.Invoke(errorObject)
					return
				}
				res, err := http.DefaultClient.Do(req)
				if err != nil {
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New(err.Error())
					reject.Invoke(errorObject)
					return
				}
				defer res.Body.Close()

				// Read the response body
				data, err := ioutil.ReadAll(res.Body)
				if err != nil {
					// Handle errors here too
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New(err.Error())
					reject.Invoke(errorObject)
					return
				}

				// "data" is a byte slice, so we need to convert it to a JS Uint8Array object
				arrayConstructor := js.Global().Get("Uint8Array")
				dataJS := arrayConstructor.New(len(data))
				js.CopyBytesToJS(dataJS, data)

				// Create a Response object and pass the data
				responseConstructor := js.Global().Get("Response")
				response := responseConstructor.New(dataJS)

				// Resolve the Promise
				resolve.Invoke(response)
			}()

			// The handler of a Promise doesn't return any value
			return nil
		})

		// Create and return the Promise object
		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}
