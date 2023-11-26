package l2env

import (
	"os"

	"github.com/HexmosTech/lama2/preprocess"
	trie "github.com/Vivino/go-autocomplete-trie"
	// "github.com/rs/zerolog/log"
)

func ProcessEnvironmentVariables(searchQuery, directory string) interface{} {
	envMap, err := preprocess.GetL2EnvVariables(directory)
	if err != nil {
		// Potential Errors:
		// - JSON marshalling errors.
		// log.Error().Str("Type", "Preprocess").Msg(err.Error())
		os.Exit(0)
	}
	if searchQuery == "" { // Handle empty searchQuery: ""
		return envMap
	}
	// Handle non-empty searchQuery: "query"
	return GetRelevantEnvs(envMap, searchQuery)
}

func GetRelevantEnvs(envMap map[string]map[string]interface{}, searchQuery string) map[string]interface{} {
	envTrie := trie.New()
	for key := range envMap {
		envTrie.Insert(key)
	}

	suggestions := envTrie.SearchAll(searchQuery)
	filteredEnvs := make(map[string]interface{})
	for _, suggestion := range suggestions {
		if env, found := envMap[suggestion]; found {
			filteredEnvs[suggestion] = env
		}
	}
	return filteredEnvs
}
