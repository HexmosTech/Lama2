#!/bin/bash
if [[ ${1} == '-h' ]] || [[ $# -eq 0 ]]; then
    echo "HELP"
    echo "${0} ver1...ver2"
    exit
fi
git log --pretty=oneline ${1}...${2}
