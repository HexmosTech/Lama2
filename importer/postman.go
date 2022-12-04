// Package `import` helps developers switch
// to Lama2 from various popular tools. The
// conversion may not be perfect, but it should
// help teams get started with minimal manual
// effort
package importer

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/HexmosTech/gabs/v2"
	"github.com/rs/zerolog/log"
)

// create 3 structures:
// Folder
// Request
// Environ

// folderMap
// requestMap
// environMap

type Folder struct {
	Name     string
	Ident    string
	Order    []string
	FullPath string
}

type Request struct {
	TheURL       string
	Name         string
	Method       string
	Auth         string
	RequestType  string
	RawModeData  string
	MultiData    map[string][]string
	ParentFolder string
	HeaderData   map[string]string
	Ident        string
}

type Environ struct {
	Name   string
	Values map[string]string
	Ident  string
}

var (
	folderMap  map[string]Folder
	requestMap map[string]Request
	environMap map[string]Environ
)

func init() {
	folderMap = make(map[string]Folder)
	requestMap = make(map[string]Request)
	environMap = make(map[string]Environ)
}

func createFolderTree(folderIDOrder []string, parentPath string) {
	for _, fID := range folderIDOrder {
		folder := folderMap[fID]

		targetPath := filepath.Join(parentPath, folder.Name)
		folder.FullPath = targetPath
		err := os.MkdirAll(targetPath, os.ModePerm)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}

		createFolderTree(folder.Order, targetPath)
	}
}

func generateFolderMap(foldersList *gabs.Container) {
	fmt.Println("@@The folders", foldersList)
	for _, folder := range foldersList.Children() {
		fmt.Println(folder)
		fName := folder.S("name").Data().(string)
		fID := folder.S("id").Data().(string)
		fOrder := make([]string, 0)
		for _, e := range folder.S("folders_order").Children() {
			fOrder = append(fOrder, e.Data().(string))
		}
		folderMap[fID] = Folder{fName, fID, fOrder, ""}
		fmt.Println("===")
	}
	fmt.Println("*****")
	fmt.Println(folderMap)
}

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
	collections := pJSON.S("collections")
	for _, collection := range collections.Children() {
		collectionFolders := collection.S("folders")
		collectionName := collection.S("name")
		collectionPath := filepath.Join(targetFolder, collectionName.Data().(string))
		err := os.MkdirAll(collectionPath, os.ModePerm)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}
		generateFolderMap(collectionFolders)
		cOrder := make([]string, 0)
		for _, e := range collection.S("folders_order").Children() {
			cOrder = append(cOrder, e.Data().(string))
		}
		fmt.Println("cOrder = ", cOrder)
		createFolderTree(cOrder, collectionPath)
	}
}
