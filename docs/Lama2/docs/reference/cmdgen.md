<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# cmdgen

```go
import "github.com/HexmosTech/lama2/cmdgen"
```

Package \`cmdgen\` provides an API to generate API request commands \(by default based on HTTPie\) based on the parsed API file contents and the \`l2\` command invocation parameters

## Index

- [func ConstructCommandHelper\(parsedInput \*gabs.Container\) \(string, string, \*gabs.Container, \*gabs.Container, bool, bool\)](<#ConstructCommandHelper>)


<a name="ConstructCommandHelper"></a>
## func [ConstructCommandHelper](<https://github.com/HexmosTech/Lama2/blob/main/cmdgen/cmdgen.go#L96>)

```go
func ConstructCommandHelper(parsedInput *gabs.Container) (string, string, *gabs.Container, *gabs.Container, bool, bool)
```

ConstructCommand extracts the HTTP verb, url and other API file inputs, figures out the type of target command and finally generates a string representing the generated command

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
