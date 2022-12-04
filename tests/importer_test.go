package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/HexmosTech/lama2/importer"
)

func TestPostman(t *testing.T) {
	path, _ := os.Getwd()
	fmt.Println(path)
	importer.PostmanConvert(filepath.Join(".", "data", "Backup.postman_dump.json"), "l2output")
}
