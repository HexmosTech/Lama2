package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/HexmosTech/lama2/cmdexec"
	"github.com/dop251/goja"
)

func TestRunVMCode(_ *testing.T) {
	vm := cmdexec.GetJSVm()
	cmdexec.RunVMCode("let meaning=42", vm)
}

func TestJSVmNegative(t *testing.T) {
	vm := cmdexec.GetJSVm()
	_, e := vm.(*goja.Runtime).RunString("faulty JS code")
	if e == nil {
		t.Errorf("Expected failure on faulty JS code")
	}
}

func TestJSVm(t *testing.T) {
	vm := cmdexec.GetJSVm()
	_, e := vm.(*goja.Runtime).RunString("let x=10")
	if e != nil {
		t.Errorf("Expected nil on valid JS code")
	}
}

func TestGenerateChainCode(t *testing.T) {
	generatedStr := cmdexec.GenerateChainCode(`{"hello": "world"}`)
	if strings.HasPrefix(generatedStr, "try") == true {
		fmt.Println("Try/Catch in TestGenerateChain")
	} else {
		t.Fatalf("TestGenerateChain doesn't start with a try")
	}
}
