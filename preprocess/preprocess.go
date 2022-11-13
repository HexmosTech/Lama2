package preprocess

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func loadElfEnv() {
	err := godotenv.Load("l2.env")
	if err != nil {
		log.Info().Str("Type", "Preprocess").Msg("Error loading elf.env file")
	}
}

func getLamaFileAsString(path string) string {
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		log.Fatal().Str("Type", "Preprocess").Msg(fmt.Sprint("Couldn't read: ", path))
	}
	return string(b)
}

func getFilePathComponents(name string) (string, string, string) {
	fullpath, _ := filepath.Abs(name)
	dir, fname := path.Split(fullpath)
	return fullpath, dir, fname
}

func changeWorkingDir(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		log.Fatal().
			Str("Type", "Preprocess").
			Str("dir", dir).
			Msg(fmt.Sprint("Moving into dir failed"))
	}
}

func PreprocessLamaFile(input_f string) string {
	content := getLamaFileAsString(input_f)
	_, dir, _ := getFilePathComponents(input_f)
	changeWorkingDir(dir)
	loadElfEnv()
	content = os.ExpandEnv(content)
	return content
}
