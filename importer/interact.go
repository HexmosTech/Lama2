package importer

func PostmanImporter(postmanFile string, outDir string) {
	pJSON := ReadPostmanFile(postmanFile)
	// importer.PickEnvironmentID(pJSON)
	environID := PickEnvironmentID(pJSON)
	PostmanConvert(pJSON, outDir, environID)
}
