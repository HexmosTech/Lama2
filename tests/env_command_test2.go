// env_command_test.go
package tests

import (
	"testing"

	testutils "github.com/HexmosTech/lama2/tests/utils"
)

func runL2Command(t *testing.T, cmdArgs ...string) map[string]testutils.EnvData {
	envMap, err := testutils.RunL2CommandAndParseJSON(cmdArgs...)
	if err != nil {
		t.Fatalf("Error running L2 command: %v", err)
	}
	return envMap
}

func TestL2EnvCommand(t *testing.T) {
	fpath := "../elfparser/ElfTestSuite/root_variable_override/api/y_0020_root_override.l2"
	cmdArgs := []string{"-e=", fpath}
	envMap := runL2Command(t, cmdArgs...)

	// Check the "AHOST" key is present
	checkAHost(t, envMap)

	// Check the "BHOST" key
	checkBHost(t, envMap)
}

func TestL2RelevantEnvForAString(t *testing.T) {
	fpath := "../elfparser/ElfTestSuite/root_variable_override/api/y_0020_root_override.l2"
	cmdArgs := []string{"-e=A", fpath}
	envMap := runL2Command(t, cmdArgs...)

	// Check the "AHOST" key is present
	checkAHost(t, envMap)

	// Check the "BHOST" key is absent
	checkBHostDoesNotExist(t, envMap)
}

func TestL2RelevantEnvForBString(t *testing.T) {
	fpath := "../elfparser/ElfTestSuite/root_variable_override/api/y_0020_root_override.l2"
	cmdArgs := []string{"-e=B", fpath}
	envMap := runL2Command(t, cmdArgs...)

	// Check the "BHOST" key is present
	checkBHost(t, envMap)

	// Check the "AHOST" key is absent
	checkAHostDoesNotExist(t, envMap)
}

func TestL2EnvCommandVerbose(t *testing.T) {
	fpath := "../elfparser/ElfTestSuite/root_variable_override/api/y_0020_root_override.l2"
	cmdArgs := []string{"-e=", "-v", fpath}
	envMap := runL2Command(t, cmdArgs...)

	// Check the "AHOST" key is present
	checkAHost(t, envMap)

	// Check the "BHOST" key is present
	checkBHost(t, envMap)
}

func TestL2EnvWithoutL2config(t *testing.T) {
	fpath := "../elfparser/ElfTestSuite/no_l2config/api/y_0021_no_l2config.l2"
	cmdArgs := []string{"-e=", fpath}
	envMap := runL2Command(t, cmdArgs...)

	// Check the "AHOST" key is present
	checkAHost(t, envMap)
}

func TestL2EnvWithoutL2env(t *testing.T) {
	fpath := "../elfparser/ElfTestSuite/no_l2env/api/y_0022_no_l2env.l2"
	cmdArgs := []string{"-e=", fpath}
	envMap := runL2Command(t, cmdArgs...)

	// Check the "BHOST" key is present
	checkBHost(t, envMap)
}

// checkAHost checks the "AHOST" key in the JSON map
func checkAHost(t *testing.T, envMap map[string]testutils.EnvData) {
	if ahost, ok := envMap["AHOST"]; !ok {
		t.Error("Expected 'AHOST' key in the JSON, but it was not found")
	} else {
		// Example assertion: Check the "AHOST" src and val values
		if ahost.Src != "l2env" {
			t.Errorf(`Expected "src" value to be "l2env" for "AHOST", but got: %v`, ahost.Src)
		}
		if ahost.Val != "`echo http://httpbin.org`" {
			t.Errorf(`Expected "val" value to be "echo http://httpbin.org" for "AHOST", but got: %v`, ahost.Val)
		}
	}
}

// checkBHost checks the "BHOST" key in the JSON map
func checkBHost(t *testing.T, envMap map[string]testutils.EnvData) {
	if bhost, ok := envMap["BHOST"]; !ok {
		t.Error("Expected 'BHOST' key in the JSON, but it was not found")
	} else {
		// Example assertion: Check the "BHOST" src and val values
		if bhost.Src != "l2configenv" {
			t.Errorf(`Expected "src" value to be "l2configenv" for "BHOST", but got: %v`, bhost.Src)
		}
		if bhost.Val != "https://httpbin.org" {
			t.Errorf(`Expected "val" value to be "https://httpbin.org" for "BHOST", but got: %v`, bhost.Val)
		}
	}
}

func checkBHostDoesNotExist(t *testing.T, envMap map[string]testutils.EnvData) {
	if _, ok := envMap["BHOST"]; ok {
		t.Error("Expected 'BHOST' key not to be present in the JSON, but it was found")
	}
}

func checkAHostDoesNotExist(t *testing.T, envMap map[string]testutils.EnvData) {
	if _, ok := envMap["AHOST"]; ok {
		t.Error("Expected 'AHOST' key not to be present in the JSON, but it was found")
	}
}
