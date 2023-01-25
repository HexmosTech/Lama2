The following is the *recommended* flow for a **simple** `.l2` file. The grammar offers some additional flexibilities in ordering the various components, but it is preferable to
stick to the following ordering to 
help with consistency. Also, for *chaining requests*, one needs a little bit more syntax -- [See Examples](../tutorials/examples.md#chain-requests-using-javascript)

```mermaid
%%{init: {'securityLevel': 'loose', 'theme':'base'}}%%
graph TD
	Z(Start)
	Z --> A
    A["HTTP Verb (get/post/put/delete)"] --> B[Multipart]
	C["URL (http://blah.com)"]
	A --> C
	B --> C
	D["Headers (header_key: header_value)"]
	E{Payload}
	C --> D
	D --> E
	F["VarJSON (key=value)"]
	G["JSON {'key': 'value'}"]
	H["Multipart files (filename@filepath)"]
	E --> F
	F --> H
	E --> G
	I("End")
	H --> I
	G --> I
```
