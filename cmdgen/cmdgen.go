package cmdgen

import (
	"fmt"

	"github.com/HexmosTech/gabs/v2"
)

func ConstructCommand(parsedInput *gabs.Container) string {
	fmt.Println("==", parsedInput.S("value", "details", "headers"))
	fmt.Println(parsedInput)
	return "result"
}
