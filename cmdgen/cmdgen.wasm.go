//go:build wasm

package cmdgen

import (
	"encoding/json"
	"fmt"
	"log"
	"syscall/js"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/lama2cmd"
)

func ConstructCommand(parsedInput *gabs.Container, o *lama2cmd.Opts) ([]string, string) {
	httpv, url, jsonObj, headers, multipartBool, formBool := ConstructCommandHelper(parsedInput)
	res, stdinBody := assembleCmdString(httpv, url, jsonObj, headers, multipartBool, formBool, nil)
	LaBearerAuthToken := js.Global().Get("liveapiManifest").String()

	// Parse JSON
	parsedJson, err := gabs.ParseJSON([]byte(LaBearerAuthToken))
	if err != nil {
		log.Fatal("Error parsing JSON:", err)
	}

	// Get project count
	projects, ok := parsedJson.Path("projects").Data().([]interface{})
	if !ok {
		log.Fatal("Failed to parse projects")
	}

	fmt.Println("Project count:", len(projects))

	projectRoot := getProjectRoot()
	if projectRoot == "" {
		fmt.Println("DBGYYY: No project_root meta tag found")
		return res, stdinBody
	}

	fmt.Println("DBGYYY: projectRoot:", projectRoot)

	// Find the matching project
	for _, project := range projects {
		projectMap, ok := project.(map[string]interface{})
		if !ok {
			continue
		}

		if projectMap["project_root"] == projectRoot {
			authData, exists := projectMap["result"].(map[string]interface{})["auth"]
			if !exists {
				fmt.Println("DBGYYY: No auth data found for this project")
			} else {
				authArray, ok := authData.([]interface{})
				if !ok {
					fmt.Println("DBGYYY: Auth data is not in expected format")
					return res, stdinBody
				}

				// Filter the auth methods where "selected" is true
				for _, authEntry := range authArray {
					authMap, ok := authEntry.(map[string]interface{})
					if !ok {
						continue
					}

					selected, exists := authMap["selected"].(bool)
					if exists && selected {
						authJSON, _ := json.MarshalIndent(authMap, "", "  ")
						fmt.Println("DBGYYY: Selected Auth Data:", string(authJSON))
					}
				}
			}
			break
		}
	}

	return res, stdinBody
}



func getProjectRoot() string {
	// Fetch project_root from the <meta> tag in the document
	meta := js.Global().Get("document").Call("querySelector", `meta[name="project_root"]`)
	if meta.IsNull() {
		return ""
	}
	return meta.Call("getAttribute", "content").String()
}
