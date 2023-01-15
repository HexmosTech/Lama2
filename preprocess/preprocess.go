// Package preprocess provides facilities to expand
// environment variables in `.l2` API files and return
// the contents
package preprocess

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/HexmosTech/lama2/utils"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func LoadElfEnv(l2path string) {
	err := godotenv.Load(l2path)
	if err != nil {
		log.Info().Str("Type", "Preprocess").Msg("Didn't find l2.env in the API directory")
	}
}

func GetLamaFileAsString(path string) string {
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		log.Fatal().Str("Type", "Preprocess").Msg(fmt.Sprint("Couldn't read: ", path))
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
	LoadElfEnv("l2.env")
	content = os.ExpandEnv(content)
	utils.ChangeWorkingDir(oldDir)

	return content, dir
}
