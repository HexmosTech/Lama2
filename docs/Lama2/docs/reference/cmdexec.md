<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# cmdexec

```go
import "github.com/HexmosTech/lama2/cmdexec"
```

Package \`cmdexec\` provides a facility to execute l2 commands, stream output to stdout, while also providing ability to retrieve the command output as a string.

## Index

- [func ExecCommand(cmdStr string, apiDir string) string](<#func-execcommand>)


## func [ExecCommand](<https://github.com/HexmosTech/Lama2/blob/master/cmdexec/cmdexec.go#L24>)

```go
func ExecCommand(cmdStr string, apiDir string) string
```

ExecCommand changes directory to the given \`apiDir\` and then executes the command specified in \`cmdStr\` During command execution, ExecCommand streams output to stdout. Once execution finishes, previous CWD is restored, and the command output is returned as a string



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
