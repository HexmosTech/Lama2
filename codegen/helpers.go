package codegen

import "strings"

func SplitLangLib(targetLangLib string) (string, string) {
	res := strings.Split(targetLangLib, ".")
	if len(res) == 2 {
		return res[0], res[1]
	}
	return targetLangLib, ""
}
