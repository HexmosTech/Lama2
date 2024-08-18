#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color

REPO="HexmosTech/Lama2"
BETA_VERSION="v1.6.2-beta"

echo -e "${YELLOW}Installing beta version $BETA_VERSION${NC}"
RELEASE_URL="https://api.github.com/repos/$REPO/releases/tags/$BETA_VERSION"

download_and_install() {
    release_info=$(curl -s $RELEASE_URL)
    download_url=$(echo "$release_info" | grep "browser_download_url.*l2\"" | cut -d '"' -f 4)

    if [ -z "$download_url" ]; then
        echo -e "${RED}Failed to find the download URL for $BETA_VERSION${NC}"
        exit 1
    fi

    echo -e "${GREEN}Downloading l2 from $download_url${NC}"
    curl -L -o /tmp/l2 "$download_url"
    
    sudo mv /tmp/l2 /usr/local/bin/l2
    sudo chmod +x /usr/local/bin/l2
}

# Main execution
download_and_install

if l2 --version > /dev/null 2>&1; then 
    echo -e "${GREEN}Successfully installed Lama2 beta version; Type 'l2 <api_file>' to invoke Lama2${NC}"
    echo -e "${YELLOW}Installed version:${NC}"
    l2 --version
else 
    echo -e "${RED}Failure in installation; please report issue at github.com/HexmosTech/Lama2${NC}"
fi