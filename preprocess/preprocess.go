// Package preprocess provides facilities to expand
// environment variables in `.l2` API files and return
// the contents
package preprocess

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/HexmosTech/godotenv"
	"github.com/HexmosTech/lama2/utils"
)

func debugOp(str string) {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(str)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("String written to file successfully.")
}

func escapeString(input string) string {
	output, err := json.Marshal(input)
	if err != nil {
		fmt.Println("Error marshaling JSON: escapeString()")
	}
	return string(output)
}

func SearchL2ConfigEnv(dir string) (string, error) {
	parentDir := dir
	for parentDir != string(filepath.Separator) {
		l2ConfigPath := filepath.Join(parentDir, "l2config.env")
		_, err := os.Stat(l2ConfigPath)
		if err == nil {
			return l2ConfigPath, nil // Found the l2config.env file
		}
		parentDir = filepath.Dir(parentDir)
	}
	return "", errors.New("Didn't find l2config.env in the API directory")
}

func LoadEnvFile(l2path string) {
	envFileName := filepath.Base(l2path)
	err := godotenv.Load(l2path)
	if err != nil {
		fmt.Println("Didn't find " + envFileName + " in the API directory")
	}
}

func LoadEnvironments(dir string) {
	l2ConfigPath, err := SearchL2ConfigEnv(dir)
	if err != nil {
		fmt.Println("Type", "Preprocess", err)
	} else {
		LoadEnvFile(l2ConfigPath) // Loads global variables from l2config.env
	}
	LoadEnvFile(path.Join(dir, "l2.env")) // Overwrites the global variables if declared again in l2.env
}

func readFile(filename string) (envMap map[string]string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	return godotenv.Parse(file)
}

func getEnvMap(envPath string, source string) (map[string]map[string]interface{}, error) {
	envs, err := readFile(envPath)
	if err != nil {
		return nil, err
	}
	envMap := make(map[string]map[string]interface{})
	for key, value := range envs {
		variable := make(map[string]interface{})
		variable["val"] = value
		variable["src"] = source
		envMap[key] = variable
	}
	return envMap, nil
}

func combineEnvMaps(envMaps ...map[string]map[string]interface{}) map[string]map[string]interface{} {
	// First l2config varibales are added to the map.
	// Then l2env varibales, so if a variable is redeclared in l2env then it will overwritten and taken to consideration.
	finalEnvMap := make(map[string]map[string]interface{})
	for _, envMap := range envMaps {
		for key, value := range envMap {
			finalEnvMap[key] = value
		}
	}
	return finalEnvMap
}

func GetL2EnvVariables(dir string) (map[string]map[string]interface{}, error) {
	l2ConfigEnvMap := make(map[string]map[string]interface{})

	l2ConfigPath, err := SearchL2ConfigEnv(dir)
	if err != nil {
		// If l2config file is not found, assign an empty map to l2EnvMap and continue
		l2ConfigEnvMap = make(map[string]map[string]interface{})
	} else {
		l2ConfigEnvMap, err = getEnvMap(l2ConfigPath, "l2configenv")
		if err != nil {
			// If an error occurs, assign an empty map to l2EnvMap and continue
			l2ConfigEnvMap = make(map[string]map[string]interface{})
		}
	}

	l2EnvPath := path.Join(dir, "l2.env")
	l2EnvMap, err := getEnvMap(l2EnvPath, "l2env")
	if err != nil {
		// If an error occurs, assign an empty map to l2EnvMap and continue
		l2EnvMap = make(map[string]map[string]interface{})
	}

	finalEnvMap := combineEnvMaps(l2ConfigEnvMap, l2EnvMap)

	return finalEnvMap, nil
}

func GetLamaFileAsString(path string) string {
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		fmt.Println("Type: Preprocess , Couldn't read: ", path)
	}
	return string(b)
}

// LamaFile takes in a path to an API file.
// It moves into the API file directory, reads the
// API contents, loads the `l2.env` file if available,
// and finally substitutes environment vars in the API contents
// Once done, it reverts back to the original directory,
// and returns the processed l2 file.
func LamaFile(inputFile string) (string, string) {
	content := GetLamaFileAsString(inputFile)
	_, dir, _ := utils.GetFilePathComponents(inputFile)
	oldDir, _ := os.Getwd()

	utils.ChangeWorkingDir(dir)
	LoadEnvFile("l2.env")
	content = os.ExpandEnv(content)
	utils.ChangeWorkingDir(oldDir)

	return content, dir
}
