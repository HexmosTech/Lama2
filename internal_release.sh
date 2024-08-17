#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

REPO="HexmosTech/Lama2"
ARTIFACT_NAME="l2"

get_latest_run_id() {
    curl -s "https://api.github.com/repos/$REPO/actions/runs?status=success" | 
    jq '.workflow_runs[0].id'
}

download_artifact() {
    run_id=$(get_latest_run_id)
    if [ -z "$run_id" ]; then
        echo -e "${RED}Failed to get the latest successful run ID${NC}"
        exit 1
    fi

    artifact_url=$(curl -s "https://api.github.com/repos/$REPO/actions/runs/$run_id/artifacts" | 
                   jq -r ".artifacts[] | select(.name == \"$ARTIFACT_NAME\") | .archive_download_url")

    if [ -z "$artifact_url" ]; then
        echo -e "${RED}Failed to find the artifact download URL${NC}"
        exit 1
    fi

    echo "Downloading artifact from $artifact_url"
    curl -L -o /tmp/l2.zip "$artifact_url"
    
    unzip /tmp/l2.zip -d /tmp
    rm /tmp/l2.zip
}

install_binary() {
    sudo rm -f /usr/local/bin/l2 /usr/bin/l2
    sudo mv /tmp/l2 /usr/local/bin
    sudo chmod +x /usr/local/bin/l2
}

# Main execution
download_artifact
install_binary

if l2 --version > /dev/null 2>&1; then 
    echo -e "${GREEN}Successfully installed Lama2; Type 'l2 <api_file>' to invoke Lama2${NC}"
else 
    echo -e "${RED}Failure in installation; please report issue at github.com/HexmosTech/Lama2${NC}"
fi