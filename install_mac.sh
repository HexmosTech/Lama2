#!/bin/bash

cmd=http
[[ $(type -P "$cmd") ]] || 
    { echo "$cmd is NOT in PATH" 1>&2; brew install httpie; }
cmd=go
[[ $(type -P "$cmd") ]] || 
    { echo "$cmd is NOT in PATH" 1>&2; brew install go; }
go get $(go list -f '{{ join .Imports "\n" }}')
go build -o elf flyweightgroup.com/httpelf
mv elf build
sudo rm /usr/local/bin/elf
sudo cp ./build/elf /usr/local/bin
echo "Built and installed latest version of httpelf"