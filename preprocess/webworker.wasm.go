//go:build wasm

package preprocess

import (
	"fmt"
	"syscall/js"

	"github.com/rs/zerolog/log"
)

var worker js.Value

func InitWebWorker() js.Value {
	fmt.Println("WASM: Initializing web worker")
	if worker.IsUndefined() {
		script := `
			var result;

			self.addEventListener("message", function (e) {
			  const { type, payload } = e.data;

			  if (type === "execute") {
			    console.log("WW: Execute Request!", payload);
			    try {
			      const evalresult = eval(payload);
			      self.postMessage({ status: "success", evalresult });
			    } catch (error) {
			      console.log("WW: Execute Request! Failed:", error.message);
			      self.postMessage({ status: "error", error: error.message });
			    }
			  } else if (type === "get") {
			    const variableName = payload;
			    var value = self[variableName];
			    if (value == null) {
			      value = variableName;
			    }
			    console.log("WW: Execute Get!", payload, value);
			    console.log("WW: getfromworker worker.js", payload, value);
			    self.postMessage({ status: "success", value });
			  } else if (type === "close") {
			    console.log("WW: Closing worker");
			    self.close();
			  } else {
			    console.log("WW: Error Unknown request type");
			    self.postMessage({ status: "error", error: "Unknown request type" });
			  }
			});
		`
		// Convert the script into a Blob and create a new worker
		blob := js.Global().Get("Blob").New([]interface{}{script}, map[string]interface{}{"type": "application/javascript"})
		workerURL := js.Global().Get("URL").Call("createObjectURL", blob)
		worker = js.Global().Get("Worker").New(workerURL)

		worker.Call("addEventListener", "message", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			result := args[0].Get("data").Get("result")
			err := args[0].Get("data").Get("error")
			if err.String() != "null" {
				fmt.Println("Error from web worker:", err)
			} else {
				fmt.Println("Result from web worker:", result)
			}
			return nil
		}))
	}
	fmt.Println("WASM: Web worker initialized")
	return worker
}

func RunCodeInWorker(chainCode string) {
	fmt.Println("WASM: Running code in web worker")
	// Ensure the worker is initialized
	InitWebWorker()
	// Send the message to the worker
	worker.Call("postMessage", map[string]interface{}{
		"type":    "execute",
		"payload": chainCode,
	})
	fmt.Println("WASM: Code sent to web worker")
}

func GetFromWorker(variableName string) string {
	fmt.Println("WASM: Getting from web worker")
	worker := InitWebWorker()
	responseChannel := make(chan js.Value, 20)

	worker.Call("postMessage", map[string]interface{}{
		"type":    "get",
		"payload": variableName,
	})

	worker.Call("addEventListener", "message", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		data := args[0].Get("data")
		if data.Get("value").IsUndefined() {
			log.Error().Msg("WW: data is undefined, skipping write to responseChannel")
			return nil
		}

		select {
		case responseChannel <- data:
			log.Debug().Msg("WW: Returning Result from web worker:")
		default:
			log.Error().Msg("WW: responseChannel is full, cannot write data")
		}
		return nil
	}))

	fmt.Println("WW: Webworker call complete:", variableName)
	var result js.Value
	result = <-responseChannel
	fmt.Println("WASM: getfromworker webworker.go", variableName, result.Get("value").String())
	return result.Get("value").String()
}
