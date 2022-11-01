package utils

import (
	"encoding/json"
	"unicode"
)

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

// UnicodeCategory returns the Unicode Character Category of the given rune.
func UnicodeCategory(r rune) string {
	for name, table := range unicode.Categories {
		if len(name) == 2 && unicode.Is(table, r) {
			return name
		}
	}
	return "Cn"
}
