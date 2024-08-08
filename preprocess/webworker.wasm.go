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

// Fix 3: Using unique identifiers
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




// Bugged version

// func GetFromWorker(variableName string) string {
//     // Ensure the worker is initialized
//     log.Debug().Msg("WW: Check webworker")
//     InitWebWorker()
//     // Send the message to the worker
//     worker.Call("postMessage", map[string]interface{}{
//         "type":    "get",
//         "payload": variableName,
//     })

//     log.Debug().Msg("WW: Webworker call complete")


//     // Wait for the response from the worker
//     result := <-responseChannel
//     fmt.Println("WW: getfromworker webworker.go",variableName,result.Get("value").String())
//     log.Debug().Msg("WW: Returning back the data")
//     return result.Get("value").String()
// }

// Fix 1 

// func GetFromWorker(variableName string) string {
//     worker = InitWebWorker()
//     var responseChannel = make(chan js.Value) // Ensure this is set to a new channel or appropriately reset
//     worker.Call("postMessage", map[string]interface{}{
//         "type":    "get",
//         "payload": variableName,
//     })

//     worker.Call("addEventListener", "message", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
//         responseChannel <- args[0].Get("data")
     
//         log.Debug().Msg("WW: Returning Result from web worker:");
//         return nil
//     }))

//     fmt.Println("WW: Webworker call complete:",variableName)
//     var result js.Value
//     result = <-responseChannel // This will block until a message is received

//     fmt.Println("WW: getfromworker webworker.go", variableName, result.Get("value").String())
//     log.Debug().Msg("WW: Returning back the data")
//     return result.Get("value").String()
// }

// Fix 2: Deadlock fix

// func GetFromWorker(variableName string) string {

//     worker := InitWebWorker()
//     responseChannel := make(chan js.Value,20)

//     worker.Call("postMessage", map[string]interface{}{
//         "type":    "get",
//         "payload": variableName,
//     })

//     worker.Call("addEventListener", "message", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
//         select {
//         case responseChannel <- args[0].Get("data"):
//             log.Debug().Msg("WW: Returning Result from web worker:")
//         default:
//             log.Error().Msg("WW: responseChannel is full, cannot write data")
//         }
//         return nil
//     }))
//     fmt.Println("WW: Webworker call complete:", variableName)
//     var result js.Value
//     result = <-responseChannel
//     fmt.Println("WW: getfromworker webworker.go", variableName, result.Get("value").String())
//     return result.Get("value").String()
// }
