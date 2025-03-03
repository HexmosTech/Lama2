//go:build wasm

package cmdgen

import (
	"log"
	"syscall/js"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/lama2cmd"
)
 

func ConstructCommand(parsedInput *gabs.Container, o *lama2cmd.Opts) ([]string, string) {
	httpv, url, jsonObj, headers, multipartBool, formBool := ConstructCommandHelper(parsedInput)
	res, stdinBody := assembleCmdString(httpv, url, jsonObj, headers, multipartBool, formBool, nil)

	log.Printf("WASM: Initial command: %v", res)
	
	manifestObj := js.Global().Get("liveapiManifest")
	if manifestObj.IsUndefined() {
		log.Printf("WASM: liveapiManifest is undefined, returning early")
		return res, stdinBody
	}
	manifestData := manifestObj.String()
	log.Printf("WASM: Found manifest data: %s", manifestData)

	// Parse JSON
	parsedJson, err := gabs.ParseJSON([]byte(manifestData))
	if err != nil {
		log.Printf("WASM: Error parsing JSON: %v", err)
		return res, stdinBody
	}

	// Get project count
	projects, ok := parsedJson.Path("projects").Data().([]interface{})
	if !ok {
		log.Printf("WASM: Failed to parse projects from manifest")
		return res, stdinBody
	}
	log.Printf("WASM: Found %d projects in manifest", len(projects))

	projectRoot := getProjectRoot()
	if projectRoot == "" {
		log.Printf("WASM: Project root is empty")
		return res, stdinBody
	}
	log.Printf("WASM: Current project root: %s", projectRoot)

	var selectedAuthHeader string

	// Find the matching project
	for _, project := range projects {
		projectMap, ok := project.(map[string]interface{})
		if !ok {
			continue
		}

		if projectMap["project_root"] == projectRoot {
			authData, exists := projectMap["result"].(map[string]interface{})["auth"]
			if !exists {
			} else {
				authArray, ok := authData.([]interface{})
				if !ok {
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
						// Construct Authorization header based on type
						authType := authMap["type"].(string)
						authValue := authMap["value"].(map[string]interface{})
						switch authType {
						case "bearer-token":
							if token, ok := authValue["token"].(string); ok {
								selectedAuthHeader = "Authorization: Bearer " + token
							}
						case "jwt":
							if jwt, ok := authValue["jwt"].(string); ok {
								selectedAuthHeader = "Authorization: Bearer " + jwt
							}
						case "api-key":
							if key, ok := authValue["key"].(string); ok {
								if value, ok := authValue["value"].(string); ok {
									selectedAuthHeader = key + ": " + value
								}
							}
						case "basic-auth":
							if username, ok := authValue["username"].(string); ok {
								if password, ok := authValue["password"].(string); ok {
									credentials := username + ":" + password
									encoded := js.Global().Get("btoa").Invoke(credentials).String()
									selectedAuthHeader = "Authorization: Basic " + encoded
								}
							}
						}
						break // Stop after the first selected auth method
					}
				}
			}
			break
		}
	}
	// Add the selected authentication method to headers if not already present
	if selectedAuthHeader != "" {
		authHeaderExists := false
		for _, header := range res {
			if len(header) >= len("Authorization:") && header[:len("Authorization:")] == "Authorization:" {
				authHeaderExists = true
				break
			}
		}

		if !authHeaderExists {
			res = append(res, selectedAuthHeader)
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
