package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/HexmosTech/lama2/importer"
)

func TestPostman(t *testing.T) {
	path, _ := os.Getwd()
	fmt.Println(path)
	importer.PostmanConvert("./data/Backup.postman_dump.json", "blah")
}
