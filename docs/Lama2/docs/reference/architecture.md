```mermaid
%%{init: {'securityLevel': 'loose', 'theme':'base'}}%%
graph TD
	A["CLI <br/>(lama2cmd)"]
	B["Parser <br/>(parser)"]	
	C[Extension API]
	D["Request Executor <br/>(cmdexec)"]
	E["Output Format Manager"]
	F["Error Reporting"]
	G["Environment Var Expansion <br/>(preprocess)"]
	H["Request Command Generator <br/>(cmdgen)"]
	I["Lama2 Prettifier"]
	A --> G
	C --> G
	G --> B
	B --> F
	B --> H
	H --> D
	D --> E
	B --> I
```
