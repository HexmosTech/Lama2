#!/bin/bash

cmd=http
[[ $(type -P "$cmd") ]] || 
    { echo "$cmd is NOT in PATH" 1>&2; sudo snap install http; }
cmd=go
[[ $(type -P "$cmd") ]] || 
    { echo "$cmd is NOT in PATH" 1>&2; sudo snap install go --classic; }
go mod tidy
go build -o l2 github.com/HexmosTech/lama2
mv l2 build
sudo rm /usr/bin/l2
sudo cp ./build/l2 /usr/bin
echo "Built and installed latest version of Lama2; run command l2 to get help."
