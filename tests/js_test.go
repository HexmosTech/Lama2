package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/HexmosTech/lama2/cmdexec"
)

func TestRunVMCode(t *testing.T) {
	vm := cmdexec.GetJSVm()
	cmdexec.RunVMCode("let meaning=42", vm)
}

func TestJSVmNegative(t *testing.T) {
	vm := cmdexec.GetJSVm()
	_, e := vm.RunString("faulty JS code")
	if e == nil {
		t.Errorf("Expected failure on faulty JS code")
	}
}

func TestJSVm(t *testing.T) {
	vm := cmdexec.GetJSVm()
	_, e := vm.RunString("let x=10")
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
