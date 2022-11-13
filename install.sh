#!/bin/bash

get_file() {
    dl_url="https://api.github.com/repos/HexmosTech/Lama2/releases/latest"
    api_resp=$(wget -nv -O - $dl_url)
}

get_platform() {
   architecture=""
   case $(uname -m) in
       i386)   architecture="386" ;;
       i686)   architecture="386" ;;
       x86_64) architecture="amd64" ;;
       arm)    dpkg --print-architecture | grep -q "arm64" && architecture="arm64" || architecture="arm" ;;
   esac 
}

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    echo "OS is Linux"
    get_file 
    get_platform
    # echo "$api_resp" | grep linux
    search="http.*linux-$architecture.tar.gz\"$"
    echo $search
    archive=$(echo "$api_resp" | grep "$search" | sed 's|[\"\,]*||g' | sed 's/browser_download_url://g' | xargs)
    wget -O /tmp/l2_latest.tar.gz $archive
    tar -xvzf /tmp/l2_latest.tar.gz -C /tmp
    sudo rm -f /usr/local/bin/l2 /usr/bin/l2
    sudo mv /tmp/l2 /usr/local/bin
elif [[ "$OSTYPE" == "darwin"* ]]; then
        # Mac OSX
    echo "OS is Mac OSX"
elif [[ "$OSTYPE" == "cygwin" ]]; then
        # POSIX compatibility layer and Linux environment emulation for Windows
    echo "OS is Cygwin"
elif [[ "$OSTYPE" == "msys" ]]; then
        # Lightweight shell and GNU utilities compiled for Windows (part of MinGW)
    echo "OS is msys"
elif [[ "$OSTYPE" == "win32" ]]; then
        # I'm not sure this can happen.
    echo "OS is win32"
elif [[ "$OSTYPE" == "freebsd"* ]]; then
        # ...
    echo "OS is freebsd"
else
        # Unknown.
    echo "Error: Unknown OS"
fi

