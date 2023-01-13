package cmdexec

import (
	"fmt"

	"github.com/dop251/goja"
)

func TestJS() {
	vm := goja.New()
	_, err := vm.RunString("let url = 'https://google.com'; let m = 100; obj = JSON.stringify({'a': 1, 'b': 2});")
	// url := vm.Get("url")
	// fmt.Println(url.String())

	var s string
	g := vm.Get("obj")
	fmt.Println(g.String())
	vm.ExportTo(g, s)
	fmt.Println(s)

	vm.RunString("let m2 = m*5")
	g2 := vm.Get("m2")
	fmt.Println(g2.String())

	if err != nil {
		panic(err)
	}
}
