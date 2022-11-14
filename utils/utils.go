package utils

import (
	"encoding/json"
	"strings"
	"unicode"

	"github.com/HexmosTech/gabs/v2"
	"github.com/rs/zerolog"
)

func SetJson(parentObj *gabs.Container, childObj *gabs.Container, key string) *gabs.Container {
	// parentObj.Object(key)
	temp := gabs.New()
	temp.Set(childObj, key)
	parentObj.Merge(temp)
	return parentObj
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func ContainsRune(s []rune, e rune) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsStringPartial(s []string, e string) bool {
	for _, a := range s {
		if strings.Contains(e, a) {
			return true
		}
	}
	return false
}

// UnicodeCategory returns the Unicode Character Category of the given rune.
func UnicodeCategory(r rune) string {
	for name, table := range unicode.Categories {
		if len(name) == 2 && unicode.Is(table, r) {
			return name
		}
	}
	return "Cn"
}

func ConfigureZeroLog(level string) {
	log_level_map := make(map[string]zerolog.Level)
	log_level_map["ERROR"] = zerolog.ErrorLevel // less information
	log_level_map["INFO"] = zerolog.InfoLevel
	log_level_map["DEBUG"] = zerolog.DebugLevel
	log_level_map["TRACE"] = zerolog.TraceLevel // maximum information
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(log_level_map[level])
}
