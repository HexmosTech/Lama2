// jsonRpcUtils.go
package l2lsp

type Params struct {
	ProcessID             *int64             `json:"processId"`
	ClientInfo            *ClientInfo        `json:"clientInfo,omitempty"`
	Locale                *string            `json:"locale,omitempty"`
	RootPath              *string            `json:"rootPath,omitempty"`
	RootURI               *string            `json:"rootUri,omitempty"`
	InitializationOptions *interface{}       `json:"initializationOptions,omitempty"`
	Capabilities          ClientCapabilities `json:"capabilities"`
	Trace                 *string            `json:"trace,omitempty"`
	WorkspaceFolders      *[]WorkspaceFolder `json:"workspaceFolders,omitempty"`
	SearchQuery           *string            `json:"searchQuery,omitempty"`
	TextDocument          *Uri               `json:"textDocument,omitempty"`
}

type LSPServer struct {
	IsShutdownRequested bool
}

type ClientInfo struct {
	Name    string  `json:"name"`
	Version *string `json:"version,omitempty"`
}

type TextDocumentClientCapabilities struct {
	Completion *CompletionClientCapabilities `json:"completion,omitempty"`
	Hover      *HoverClientCapabilities      `json:"hover,omitempty"`
	// ... Add other text document capabilities as needed
}

type CompletionClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
type Uri struct {
	Uri *string `json:"uri,omitempty"`
}
type HoverClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

type WorkspaceFolder struct {
	URI  string `json:"uri"`
	Name string `json:"name"`
}
type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
}

type ClientCapabilities struct {
	TextDocument *TextDocumentClientCapabilities `json:"textDocument,omitempty"`
	// ... Add other capability sections (e.g., workspace) as needed
}

type ServerCapabilities struct {
	TextDocumentSync int                `json:"textDocumentSync"`
	Completion       *CompletionOptions `json:"completion,omitempty"`
	HoverProvider    bool               `json:"hoverProvider,omitempty"`
	SuggestL2Envs    bool               `json:"suggestL2Env,omitempty"`
}

type CompletionOptions struct {
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`
	ResolveProvider   bool     `json:"resolveProvider,omitempty"`
}

type JSONRPCRequest struct {
	ID      int    `json:"id"`
	Method  string `json:"method"`
	Params  Params `json:"params,omitempty"`
	JSONRPC string `json:"jsonrpc"`
}

type JSONRPCSuggestEnvRequest struct {
	ID      int         `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	JSONRPC string      `json:"jsonrpc"`
}

type JSONRPCResponse struct {
	ID      int         `json:"id"`
	Result  interface{} `json:"result"`
	JSONRPC string      `json:"jsonrpc"`
	Error   interface{} `json:"error,omitempty"`
}
