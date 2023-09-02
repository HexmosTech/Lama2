 curl -X POST -H "Content-Type: application/json" -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "LSPService.Process",
    "params": {
        "jsonrpc": "2.0",
        "id": 1,
        "method": "initialize"
    }
}' http://localhost:9999/lsp

