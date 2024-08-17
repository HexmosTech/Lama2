#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

REPO="HexmosTech/Lama2"
ARTIFACT_NAME="l2-binary"

# Function to get the latest workflow run ID
get_latest_workflow_run_id() {
    curl -s "https://api.github.com/repos/$REPO/actions/runs?status=success" | 
    jq '.workflow_runs[0].id'
}

# Function to get the artifact download URL
get_artifact_url() {
    local run_id=$1
    curl -s "https://api.github.com/repos/$REPO/actions/runs/$run_id/artifacts" |
    jq -r ".artifacts[] | select(.name == \"$ARTIFACT_NAME\") | .archive_download_url"
}

# Main installation process
main() {
    echo "Fetching the latest successful workflow run..."
    run_id=$(get_latest_workflow_run_id)
    echo "Latest workflow run ID: $run_id"

    if [ -z "$run_id" ]; then
        echo -e "${RED}Failed to fetch the latest workflow run.${NC}"
        exit 1
    fi

    echo "Getting artifact download URL..."
    artifact_url=$(get_artifact_url $run_id)

    if [ -z "$artifact_url" ]; then
        echo -e "${RED}Failed to get artifact download URL.${NC}"
        exit 1
    fi

    echo "Downloading l2 binary..."
    curl -L -o /tmp/l2.zip "$artifact_url"

    echo "Extracting l2 binary..."
    unzip -o /tmp/l2.zip -d /tmp

    echo "Installing l2 binary..."
    sudo mv /tmp/l2 /usr/local/bin/l2
    sudo chmod +x /usr/local/bin/l2

    if l2 > /dev/null 2>&1; then 
        echo -e "${GREEN}Successfully installed Lama2; Type 'l2 <api_file>' to invoke Lama2${NC}"
    else 
        echo -e "${RED}Failure in installation; please report issue at github.com/HexmosTech/Lama2${NC}"
    fi

    # Clean up
    rm /tmp/l2.zip
}

# Run the main function
main