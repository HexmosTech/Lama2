<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# preprocess

```go
import "github.com/HexmosTech/lama2/preprocess"
```

Package preprocess provides facilities to expand environment variables in \`.l2\` API files and return the contents

## Index

- [func LamaFile(inputFile string) (string, string)](<#func-lamafile>)


## func [LamaFile](<https://github.com/HexmosTech/Lama2/blob/master/preprocess/preprocess.go#L37>)

```go
func LamaFile(inputFile string) (string, string)
```

LamaFile takes in a path to an API file. It moves into the API file directory, reads the API contents, loads the \`l2.env\` file if available, and finally substitutes environment vars in the API contents Once done, it reverts back to the original directory, and returns the processed l2 file.



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
