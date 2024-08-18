#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color

REPO="HexmosTech/Lama2"
BETA_VERSION="v1.6.2-beta"
VSIX_URL="https://github.com/HexmosTech/Lama2Code/raw/feature-ui/lama2-3.7.3.vsix"

echo -e "${YELLOW}Installing beta version $BETA_VERSION${NC}"
RELEASE_URL="https://api.github.com/repos/$REPO/releases/tags/$BETA_VERSION"

download_and_install_l2() {
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

download_and_install_vsix() {
    echo -e "${GREEN}Downloading VSIX file from $VSIX_URL${NC}"
    wget -O /tmp/lama2-3.7.3.vsix "$VSIX_URL"

    if [ $? -ne 0 ]; then
        echo -e "${RED}Failed to download VSIX file${NC}"
        exit 1
    fi

    echo -e "${GREEN}Installing VSIX in VS Code${NC}"
    code --install-extension /tmp/lama2-3.7.3.vsix

    if [ $? -ne 0 ]; then
        echo -e "${RED}Failed to install VSIX in VS Code${NC}"
        exit 1
    fi

    rm /tmp/lama2-3.7.3.vsix
}

# Main execution
download_and_install_l2

if l2 --version > /dev/null 2>&1; then 
    echo -e "${GREEN}Successfully installed Lama2 beta version; Type 'l2 <api_file>' to invoke Lama2${NC}"
    echo -e "${YELLOW}Installed version:${NC}"
    l2 --version

    # Download and install VSIX after successful l2 installation
    download_and_install_vsix
else 
    echo -e "${RED}Failure in installation; please report issue at github.com/HexmosTech/Lama2${NC}"
fi