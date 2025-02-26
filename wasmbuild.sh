#!/bin/bash
export AWS_ACCESS_KEY_ID=""
export AWS_SECRET_ACCESS_KEY=""
GOOS=js GOARCH=wasm go build -a -gcflags=all="-l -B -wb=false" -ldflags="-w -s" -o static/main.wasm
# cp static/main.wasm /home/i3nux-mint/repos/liveapi/src/views
aws s3 cp static/main.wasm s3://temp-2d/temp/main.wasm --region ap-south-1


