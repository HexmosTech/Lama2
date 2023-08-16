package env

import (
	"os"

	"github.com/HexmosTech/lama2/lama2cmd"
	"github.com/HexmosTech/lama2/preprocess"
	"github.com/HexmosTech/lama2/utils"
	trie "github.com/Vivino/go-autocomplete-trie"
	"github.com/rs/zerolog/log"
)

func ProcessEnvironmentVariables(o *lama2cmd.Opts, directory string) {
	envMap, err := preprocess.GetL2EnvVariables(directory)
	if err != nil {
		// Potential Errors:
		// - JSON marshalling errors.
		log.Error().Str("Type", "Preprocess").Msg(err.Error())
		os.Exit(0)
	}
	// Check if it's an -e invocation
	if o.Env != "UNSET_VU5TRVQ" {
		if o.Env == "" { // Handle empty -e=''
			utils.MarshalAndPrintJSON(envMap)
		} else { // Handle non-empty -e
			relevantEnvs := GetRelevantEnvs(envMap, o)
			utils.MarshalAndPrintJSON(relevantEnvs)
		}
	}
	// If not an -e invocation, the function just continues.
}

func GetRelevantEnvs(envMap map[string]map[string]interface{}, o *lama2cmd.Opts) map[string]interface{} {
	envTrie := trie.New()
	for key := range envMap {
		envTrie.Insert(key)
	}

	searchQuery := o.Env
	suggestions := envTrie.SearchAll(searchQuery)
	filteredEnvs := make(map[string]interface{})
	for _, suggestion := range suggestions {
		if env, found := envMap[suggestion]; found {
			filteredEnvs[suggestion] = env
		}
	}
	return filteredEnvs
}
