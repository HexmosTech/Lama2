// Package `import` helps developers switch
// to Lama2 from various popular tools. The
// conversion may not be perfect, but it should
// help teams get started with minimal manual
// effort
package importer

import (
	"fmt"
	"os"

	"github.com/HexmosTech/gabs/v2"
	"github.com/rs/zerolog/log"
)

// PostmanConvert takes in a Postman data file
// and generates a roughly equivalent Lama2 repository.
// Collections and subcollections become folders.
// Requests become files, environments get stored in
// `l2.env` while file attachments get copied relative
// to the API file
func PostmanConvert(postmanFile string, targetFolder string) {
	fmt.Println(postmanFile, targetFolder)
	contents, e := os.ReadFile(postmanFile)
	if e != nil {
		log.Fatal().Msg(e.Error())
	}
	pJSON, e2 := gabs.ParseJSON(contents)
	if e2 != nil {
		log.Fatal().Msg(e.Error())
	}
	coll := pJSON.S("collections")
	for _, child := range coll.Children() {
		for _, folder := range child.S("folders_order").Children() {
			fmt.Println(folder)
		}
		fmt.Println("===")
	}
}
