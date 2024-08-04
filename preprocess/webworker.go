package preprocess

import (
	"fmt"
	"syscall/js"

	"github.com/rs/zerolog/log"

	"encoding/json"
)

var worker js.Value
var responseChannel = make(chan js.Value)

// func initWebWorker() {
// 	worker = js.Global().Get("Worker").New("worker.js")

// 	worker.Call("addEventListener", "message", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
// 		result := args[0].Get("data")
// 		// Handle the result from the web worker
// 		fmt.Println("Result from web worker:", result)
// 		return nil
// 	}))
// }

func InitWebWorker() js.Value {
	if worker.IsUndefined() {
		log.Debug().Msg("WW Initializing Web worker")
		worker = js.Global().Get("Worker").New("worker.js")
		worker.Call("addEventListener", "message", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			//
			data := args[0].Get("data")
			dataJSON, err := json.Marshal(data)
			if err != nil {
				log.Error().Err(err).Msg("Failed to marshal data")
			} else {
				log.Debug().Msgf("WW: args[0].Get(\"data\"): %s", dataJSON)
			}

			//
			responseChannel <- args[0].Get("data")

			log.Debug().Msg("WW: Returning Result from web worker:")
			return nil
		}))
	}
	return worker
}

func RunCodeInWorker(chainCode string) {
	// Ensure the worker is initialized
	InitWebWorker()
	// Send the message to the worker
	worker.Call("postMessage", map[string]interface{}{
		"type":    "execute",
		"payload": chainCode,
	})
}

func GetFromWorker(variableName string) string {
	// Ensure the worker is initialized
	log.Debug().Msg("WW: Check webworker")
	InitWebWorker()
	// Send the message to the worker
	worker.Call("postMessage", map[string]interface{}{
		"type":    "get",
		"payload": variableName,
	})

	log.Debug().Msg("WW: Webworker call complete")

	// Wait for the response from the worker
	result := <-responseChannel
	fmt.Println("WW: getfromworker webworker.go", variableName, result.Get("value").String())
	log.Debug().Msg("WW: Returning back the data")
	return result.Get("value").String()
}
