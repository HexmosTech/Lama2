package main

import (
	contoller "github.com/HexmosTech/lama2/controller"
)

var version string

func main() {
	if len(version) == 0 {
		version = "vUnset"
	}
	// filePath := "/home/lince/hexmos/apihub/apihub/fb_backend_v3/search.l2"
	// filePath := "/home/lince/hexmos/apihub/apihub/parse_cloud/oneLogin/externalProviderSignup/login.l2"
	// execute.ProcessForLSP(filePath)
	contoller.Process(version)
}
