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
	"strings"

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
	IsMultipart  string
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
	folderMap  map[string]*Folder
	requestMap map[string]Request
	environMap map[string]Environ
)

func init() {
	folderMap = make(map[string]*Folder)
	requestMap = make(map[string]Request)
	environMap = make(map[string]Environ)
}

func createFolderTree(folderIDOrder []string, parentPath string) {
	for _, fID := range folderIDOrder {
		folder := folderMap[fID]

		targetPath := filepath.Join(parentPath, folder.Name)
		folder.FullPath = targetPath
		fmt.Println("@@P1", folder.FullPath, targetPath)
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
		folderMap[fID] = &Folder{fName, fID, fOrder, ""}
		fmt.Println("===")
	}
	fmt.Println("*****")
	fmt.Println(folderMap)
}

func createRequestFiles(collectionPath string) {
	for rID, rObj := range requestMap {
		op := make([]string, 0)
		if rObj.RequestType == "GET" {
			fmt.Println(">>rObj", rID)
			op = append(op, "GET")
			op = append(op, rObj.TheURL)
			res := strings.Join(op, "\n")
			fmt.Println(res)
			targetPath := ""
			if len(rObj.ParentFolder) == 0 {
				targetPath = filepath.Join(collectionPath, rObj.Name+".l2")
				fmt.Println("TPATH1", targetPath)
			} else {
				fmt.Println("@TP", folderMap[rObj.ParentFolder])
				targetPath = filepath.Join(folderMap[rObj.ParentFolder].FullPath, rObj.Name+".l2")
				fmt.Println("TPATH2", targetPath)
			}
			err := os.WriteFile(targetPath, []byte(res), 0644)
			if err != nil {
				log.Fatal().Msg(err.Error())
			}
		}
	}
}

func generateRequestMap(requestsList *gabs.Container) {
	fmt.Println("@@The requests", requestsList)
	for _, request := range requestsList.Children() {
		fmt.Println(request)
		rName := request.S("name").Data().(string)
		rID := request.S("id").Data().(string)
		rURL := request.S("url").Data().(string)
		rRequestType := request.S("method").Data().(string)
		rRawModeData := ""
		if rRequestType == "POST" {
			rRawModeDataObj := request.S("rawModeData").Data()
			if rRawModeDataObj != nil {
				rRawModeData = rRawModeDataObj.(string)
			}
		}
		rMultiData := make(map[string][]string)
		rFolderData := request.S("folder").Data()
		rFolder := ""
		if rFolderData != nil {
			rFolder = rFolderData.(string)
		}
		rHeaderData := make(map[string]string)
		requestMap[rID] = Request{rURL, rName, "", "", rRequestType, rRawModeData, rMultiData, rFolder, rHeaderData, rID}
		fmt.Println("===")
	}
	fmt.Println("*****")
	fmt.Println(requestMap)
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
		collectionRequests := collection.S("requests")
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
		generateRequestMap(collectionRequests)
		createRequestFiles(collectionPath)
	}
}
