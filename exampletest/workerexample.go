package main

import (
	"fmt"
	"time"
)

type WorkerMessage struct {
	Type    string
	Payload string
}

var workerChannel = make(chan WorkerMessage, 20)

func InitWorker() {
	go func() {
		for msg := range workerChannel {
			switch msg.Type {
			case "execute":
				fmt.Println("Executing code:", msg.Payload)
			case "get":
				// Simulate retrieving data
				workerChannel <- WorkerMessage{Type: "response", Payload: "result"}
			}
		}
	}()
}

func RunCodeInWorker(chainCode string) {
	// Ensure the worker is initialized
	InitWorker()
	// Send the message to the worker
	workerChannel <- WorkerMessage{
		Type:    "execute",
		Payload: chainCode,
	}
}

func GetFromWorker(variableName string) string {
	responseChannel := make(chan WorkerMessage, 1) // Using a buffered channel

	workerChannel <- WorkerMessage{
		Type:    "get",
		Payload: variableName,
	}

	go func() {
		for msg := range workerChannel {
			if msg.Type == "response" {
				responseChannel <- msg
				break
			}
		}
	}()

	select {
	case result := <-responseChannel:
		fmt.Println("Result from worker:", result.Payload)
		return result.Payload
	case <-time.After(2 * time.Second): // Timeout if no response
		fmt.Println("Timeout waiting for worker response")
		return ""
	}
}

func main() {
	// Example usage
	variableName := "var1"
	chainCode := "var var1 = 'Hello'"

	// Run some code in the worker
	RunCodeInWorker(chainCode)

	// Get the result from the worker
	result := GetFromWorker(variableName)
	fmt.Println("Result:", result)
    result = GetFromWorker(variableName)
	fmt.Println("Result:", result)
}
