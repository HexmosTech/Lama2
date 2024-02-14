#!/bin/bash
GOOS=js GOARCH=wasm go build -a -gcflags=all="-l -B -wb=false" -ldflags="-w -s" -o static/main.wasm
cp /home/sreedeep/Downloads/Lama2/static/main.wasm /home/sreedeep/js-widget/dist/main.wasm


resp, err := httpie.Lama2Entry(cmdSlice, strings.NewReader(stdinBody), proxyURL, proxyUserName, proxyUserPassword, allowRedirects)