package cmdexec

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
)

func TestJS() {
	vm := goja.New()
	new(require.Registry).Enable(vm)
	console.Enable(vm)
	_, err := vm.RunString(`
	js = JSON.stringify
	let url = 'https://google.com'; 
	let m = 100; 
	obj = js({'a': 1, 'b': 2});
	o2 = 0
	for (let i = 1; i <= 10; i++) {
		o2 += i
	}
	`)
	if ex, ok := err.(*goja.Exception); ok {
		fmt.Println(ex.String())
	}
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

	g3 := vm.Get("o2")
	fmt.Println(g3.String())

	if err != nil {
		panic(err)
	}
}
