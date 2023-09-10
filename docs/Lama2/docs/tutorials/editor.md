## Useful Options

The `l2` command provides some helpful options for
extension developers. The options are:

1. `--nocolor` or `-n` disables colored output in httpie-go (in CLI);
2. `--output=<target.json` or `-o` writes a structured JSON
   output to the target file (without colors). The following is the content
   structure:

```json
{
    "logs": ...,
    "headers": ...,
    "body": ...
}
```

- `logs`: String value. Contains ASCII-color-coded log output. Control the logging level through `-v` option
- `headers`: A string consisting of `HTTP header: Value` pairs separated by newline.
- `body`: A string containing the HTTP response. Usually a JSON or HTML response.

!!! note

    Right now, all the three values in the JSON are strings. In the future, we may transform the values further to provide a more parse-friendly structure.

## The Commands

### Execute current file

Combining the options `-n` and `-o`, we get:

```bash
l2 -n -o /tmp/lama2.json my_api.l2
```

The command mentioned above disables HTTPie colors,
writes the whole transaction to a structured JSON,
while also printing details into `stdout`.

The extension author can simply read the file, and
display the contents to users appropriately. For an
example, see [Lama2 for VSCode](https://github.com/HexmosTech/Lama2Code)
(also see [Marketplace page](https://marketplace.visualstudio.com/items?itemName=hexmos.Lama2)).

# The LSP Methods

## Initilize server

**Overview:**

The server initialization is the first step to establish communication with the LSP.
This process involves starting a server that listens for incoming JSON RPC 2.0 requests.

**Steps:**

### 1. Starting the Server:

Launch the server using the following command:

```bash
l2 --lsp
```

### 2. Making a Request:

Send a JSON RPC 2.0 request to the server's stdin. Here's an example of an initialization request:

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "initialize",
  "params": {
    "processId": null,
    "clientInfo": {
      "name": "MyEditor",
      "version": "1.0.0"
    },
    "rootUri": "file:///path/to/workspace"
  }
}
```

### 3. Receiving a Response:

After processing, the server will send a response back. This response can be read from the server's stdout.

### Environment variable autocompletion

**Overview:**

The LSP server offers a custom method, `suggest/environmentVariables``, which provides autocompletion suggestions for environment variables.
This method combines data from both l2.env and l2config.env to present a unified JSON representation of the environment variables.

**Steps:**

### 1. Making a Request

Here's an example request to obtain all environment variables:

```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "method": "suggest/environmentVariables",
  "params": {
    "textDocument": {
      "uri": "file:///home/Lama2/elfparser/ElfTestSuite/root_variable_override/api/y_0020_root_override.l2"
    },
    "position": {
      "line": 1,
      "character": 2
    },
    "searchQuery": ""
  }
}
```

### 2. Receiving a Response:

The server will respond with a JSON object containing the environment variables.
Extension authors can then read this response from the server's stdout and present the variables to users in a suitable format.

```json
{
  "id": 2,
  "result": {
    "AHOST": {
      "src": "l2env",
      "val": "http://127.0.0.1:8000"
    },
    "BHOST": {
      "src": "l2configenv",
      "val": "https://httpbin.org"
    }
  },
  "jsonrpc": "2.0"
}
```

To filter the results based on a specific prefix, modify the searchQuery parameter. For instance, to retrieve variables relevant to "BH":

```json
{
  "searchQuery": "BH"
}
```

When `l2.env` is present

![l2envvariable variable](l2envvariable.png)

When both `l2.env` and `l2condig.env` is present
![l2configvariable variable](l2configvariable.png)

Go to [Example](../tutorials/examples.md#case-3-override-root-variable-with-local-variable)

## Syntax Highlighting

The VSCode plugin implements a rudimentary syntax highlighting for `.l2` files. We use [Iro](https://eeyo.io/iro/documentation/) for syntax grammar generation. Find [more details](https://github.com/HexmosTech/Lama2/tree/main/syntax/README.md) if interested.
