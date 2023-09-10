# How to Add a New Method to the LSP Binary

This guide will walk you through the process of adding a new method to the LSP binary in the Lama2 project.

## 1. Understanding the Entry Point at [lama2cmd.go](https://github.com/HexmosTech/Lama2/blob/main/lama2cmd/lama2cmd.go)

The LSP server is initiated by `l2 --lsp`. This is handled in the `ArgParsing` function:

```go
func ArgParsing(o *Opts, version string) {
    if o.Lsp {
        l2lsp.StartLspServer()
        // Incoming requests to the LSP will be processed by l2lsp.StartLspServer()
    }
}
```

## 2. Directing the Input

The input is directed to the `HandleMethod` function:

```go
rpcResponse := HandleMethod(rpcRequest)
```

You can view the implementation of this function in the [identify_method.go](https://github.com/HexmosTech/Lama2/blob/main/l2lsp/identify_method.go) file.

## 3. Adding Your Method

To add your custom method, modify the `HandleMethod` function by adding a new case to the switch statement.

## 4. Implementing the Logic

Write the logic for your method in the [methods](https://github.com/HexmosTech/Lama2/blob/main/l2lsp/methods/) directory.

## 5. Using and Defining Structs

For general request structures, refer to `JSONRPCRequest` in [general_request.go](https://github.com/HexmosTech/Lama2/blob/main/l2lsp/request/general_request.go).

If you need to define a custom struct for your method, add it in the [request](https://github.com/HexmosTech/Lama2/blob/main/l2lsp/request) directory.

For general response structures, refer to `JSONRPCResponse` in [general_response.go](https://github.com/HexmosTech/Lama2/blob/main/l2lsp/response/general_response.go).

## 6. Utilizing Helper Functions

The [respond.go](https://github.com/HexmosTech/Lama2/blob/main/l2lsp/response/respond.go) file provides helper functions to create success and error responses:

```go
func CreateSuccessResponse(requestID int, result interface{}) JSONRPCResponse {...}
func ErrorResp(req request.JSONRPCRequest, errorCode int, errorMsg string) JSONRPCResponse {...}
```

Use these functions to streamline the response creation process.
