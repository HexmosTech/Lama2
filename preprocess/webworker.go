package preprocess

import (
	"fmt"
	"syscall/js"

	// "sync"

	"github.com/rs/zerolog/log"
)

var worker js.Value
// var responseChannel = make(chan js.Value)
// var responseChannel = make(chan js.Value, 1)
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

// Fix 3: Deadlock fix + Wg.Done()


func GetFromWorker(variableName string) string {
    worker := InitWebWorker()
    responseChannel := make(chan js.Value, 20) // Using a buffered channel

    // var wg sync.WaitGroup
    // wg.Add(1) 

    worker.Call("postMessage", map[string]interface{}{
        "type":    "get",
        "payload": variableName,
    })

    worker.Call("addEventListener", "message", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        // defer wg.Done() 

        select {
        case responseChannel <- args[0].Get("data"):
            log.Debug().Msg("WW: Returning Result from web worker:")
        default:
            log.Error().Msg("WW: responseChannel is full, cannot write data")
            }
        return nil
    }))

    fmt.Println("WW: Webworker call complete:", variableName)
    // wg.Wait() 

    var result js.Value
    result = <-responseChannel;
    fmt.Println("WW: getfromworker webworker.go", variableName, result.Get("value").String())
    return result.Get("value").String()
}