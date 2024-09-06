package request

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
	FilePath              string            `json:"filePath,omitempty"`
}

type ClientInfo struct {
	Name    string  `json:"name"`
	Version *string `json:"version,omitempty"`
}

type Uri struct {
	URI *string `json:"uri,omitempty"`
}

type WorkspaceFolder struct {
	URI  string `json:"uri"`
	Name string `json:"name"`
}

type JSONRPCRequest struct {
	ID      int    `json:"id"`
	Method  string `json:"method"`
	Params  Params `json:"params,omitempty"`
	JSONRPC string `json:"jsonrpc"`
}

// clear down
type LSPServer struct {
	IsShutdownRequested bool
}

type TextDocumentClientCapabilities struct {
	Completion *CompletionClientCapabilities `json:"completion,omitempty"`
	Hover      *HoverClientCapabilities      `json:"hover,omitempty"`
	// ... Add other text document capabilities as needed
}

type CompletionClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

type HoverClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

type ClientCapabilities struct {
	TextDocument *TextDocumentClientCapabilities `json:"textDocument,omitempty"`
	// ... Add other capability sections (e.g., workspace) as needed
}
