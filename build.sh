#!/bin/bash

cmd=http
[[ $(type -P "$cmd") ]] || 
    { echo "$cmd is NOT in PATH" 1>&2; sudo snap install http; }
go build
mv httpelf build
echo "Find httpelf binary in ./build/ directory"
