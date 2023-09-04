package l2envpackege

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/HexmosTech/lama2/preprocess"
	trie "github.com/Vivino/go-autocomplete-trie"
	"github.com/rs/zerolog/log"
)

func ProcessEnvironmentVariables(relevantSearchString, directory string) interface{} {
	envMap, err := preprocess.GetL2EnvVariables(directory)
	if err != nil {
		// Potential Errors:
		// - JSON marshalling errors.
		log.Error().Str("Type", "Preprocess").Msg(err.Error())
		os.Exit(0)
	}
	if relevantSearchString == "" { // Handle empty -e=''
		return envMap
	} else { // Handle non-empty -e
		relevantEnvs := GetRelevantEnvs(envMap, relevantSearchString)
		return relevantEnvs
	}

}

func GetRelevantEnvs(envMap map[string]map[string]interface{}, relevantSearchString string) map[string]interface{} {
	envTrie := trie.New()
	for key := range envMap {
		envTrie.Insert(key)
	}

	searchQuery := relevantSearchString
	suggestions := envTrie.SearchAll(searchQuery)
	filteredEnvs := make(map[string]interface{})
	for _, suggestion := range suggestions {
		if env, found := envMap[suggestion]; found {
			filteredEnvs[suggestion] = env
		}
	}
	return filteredEnvs
}

func createJsonENv(data interface{}) []byte {
	filteredJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Error().Str("Type", "Preprocess").Msg(fmt.Sprintf("Failed to marshal JSON: %v", err))
		os.Exit(0)
	}
	return filteredJSON
}
