package preprocess

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/HexmosTech/lama2/utils"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func loadElfEnv() {
	err := godotenv.Load("l2.env")
	if err != nil {
		log.Info().Str("Type", "Preprocess").Msg("Didn't find l2.env in the API directory")
	}
}

func getLamaFileAsString(path string) string {
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		log.Fatal().Str("Type", "Preprocess").Msg(fmt.Sprint("Couldn't read: ", path))
	}
	return string(b)
}

func PreprocessLamaFile(input_f string) (string, string) {
	content := getLamaFileAsString(input_f)
	_, dir, _ := utils.GetFilePathComponents(input_f)
	oldDir, _ := os.Getwd()

	utils.ChangeWorkingDir(dir)
	loadElfEnv()
	content = os.ExpandEnv(content)
	utils.ChangeWorkingDir(oldDir)

	return content, dir
}
