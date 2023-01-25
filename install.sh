#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

get_file() {
    dl_url="https://api.github.com/repos/HexmosTech/Lama2/releases/latest"
    api_resp=$(wget -nv -O - $dl_url)
}

get_platform() {
    architecture=""
    case $(uname -m) in
    i386) architecture="386" ;;
    i686) architecture="386" ;;
    x86_64) architecture="amd64" ;;
    arm) dpkg --print-architecture | grep -q "arm64" && architecture="arm64" || architecture="arm" ;;
    arm64) architecture="arm64" ;;
    *) exit 1
    esac
}

get_os() {
    the_os=""
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        echo "OS is Linux"
        the_os="linux"
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        # Mac OSX
        echo "OS is Mac OSX"
        the_os="darwin"
    elif [[ "$OSTYPE" == "cygwin" ]]; then
        # POSIX compatibility layer and Linux environment emulation for Windows
        echo "OS is Cygwin"
        echo "Installer not supported yet; please use release binary"
        exit
    elif [[ "$OSTYPE" == "msys" ]]; then
        # Lightweight shell and GNU utilities compiled for Windows (part of MinGW)
        echo "OS is msys"
        echo "Installer not supported yet; please use release binary"
        exit
    elif [[ "$OSTYPE" == "win32" ]]; then
        # I'm not sure this can happen.
        echo "OS is win32"
        echo "Installer not supported yet; please use release binary"
        exit
    elif [[ "$OSTYPE" == "freebsd"* ]]; then
        # ...
        echo "OS is freebsd"
        echo "Installer not supported yet; please use release binary"
        exit
    else
        # Unknown.
        echo "Error: Unknown OS"
        exit
    fi
}


get_file
get_platform
get_os
search="http.*${the_os}-${architecture}.tar.gz\"$"
echo "${search}"
archive=$(echo "${api_resp}" | grep "${search}" | sed 's|[\"\,]*||g' | sed 's/browser_download_url://g' | xargs)
wget -O /tmp/l2_latest.tar.gz "${archive}"
tar -xvzf /tmp/l2_latest.tar.gz -C /tmp
sudo rm -f /usr/local/bin/l2 /usr/bin/l2
sudo mv /tmp/l2 /usr/local/bin


if l2 > /dev/null 2>&1; then 
    echo -e $"${GREEN}Successfully installed Lama2; Type 'l2 <api_file>' to invoke Lama2${NC}"
else 
    echo -e $"${RED}Failure in installation; please report issue at github.com/HexmosTech/Lama2${NC}"
fi