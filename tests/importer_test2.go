package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/HexmosTech/lama2/importer"
)

func TestPostman(_ *testing.T) {
	path, _ := os.Getwd()
	fmt.Println(path)
	pJSON := importer.ReadPostmanFile(filepath.Join(".", "data", "Backup.postman_dump.json"))
	// importer.PickEnvironmentID(pJSON)
	environID := "916fd217-d9fb-45a5-8fd2-bd6cb1de42a0"
	importer.PostmanConvert(pJSON, "l2output", environID)
}
