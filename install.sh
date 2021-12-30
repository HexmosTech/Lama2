#!/bin/bash

cmd=http
[[ $(type -P "$cmd") ]] || 
    { echo "$cmd is NOT in PATH" 1>&2; sudo snap install http; }
cmd=go
[[ $(type -P "$cmd") ]] || 
    { echo "$cmd is NOT in PATH" 1>&2; sudo snap install go --classic; }
go get $(go list -f '{{ join .Imports "\n" }}')
go build -o elf flyweightgroup.com/httpelf
mv elf build
sudo rm /usr/bin/elf
sudo cp ./build/elf /usr/bin
echo "Built and installed latest version of httpelf"
