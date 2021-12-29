#!/bin/bash

cmd=http
[[ $(type -P "$cmd") ]] || 
    { echo "$cmd is NOT in PATH" 1>&2; sudo snap install http; }
cmd=go
[[ $(type -P "$cmd") ]] || 
    { echo "$cmd is NOT in PATH" 1>&2; sudo snap install go --classic; }
go build
mv httpelf build
sudo cp ./build/httpelf /usr/bin
echo "Built and installed latest version of httpelf"
