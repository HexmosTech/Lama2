#!/bin/bash
GOOS=js GOARCH=wasm go build -a -gcflags=all="-l -B -wb=false" -ldflags="-w -s" -o static/main.wasm
cp static/main.wasm /home/i3nux-mint/repos/liveapi/src/views/dist


