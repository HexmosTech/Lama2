//go:build wasm
package preprocess

import (
	"fmt"
	"syscall/js"

	"github.com/rs/zerolog/log"
)

var worker js.Value

func InitWebWorker() js.Value {
	if worker.IsUndefined() {
		log.Debug().Msg("WW Initializing Web worker")
		worker = js.Global().Get("Worker").New("worker.js")
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

    worker := InitWebWorker()
    responseChannel := make(chan js.Value,20)

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
    fmt.Println("WW: getfromworker webworker.go", variableName, result.Get("value").String())
    return result.Get("value").String()
}
