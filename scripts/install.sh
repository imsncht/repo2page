#!/bin/bash
# repo2page installer script
# Usage: curl -sfL https://raw.githubusercontent.com/imsncht/repo2page/main/scripts/install.sh | sh

set -e

OWNER="imsncht"
REPO="repo2page"
BINARY="repo2page"
FORMAT="tar.gz"
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [ "$ARCH" == "x86_64" ]; then
    ARCH="amd64"
elif [ "$ARCH" == "aarch64" ] || [ "$ARCH" == "arm64" ]; then
    ARCH="arm64"
else
    echo "Unsupported architecture: $ARCH"
    exit 1
fi

LATEST_URL="https://api.github.com/repos/$OWNER/$REPO/releases/latest"

echo "Finding latest release for $OS/$ARCH..."

# Get download URL for specific OS/Arch and format
DOWNLOAD_URL=$(curl -s $LATEST_URL | grep "browser_download_url" | grep "$OS" | grep "$ARCH" | grep "$FORMAT" | cut -d '"' -f 4)

if [ -z "$DOWNLOAD_URL" ]; then
    echo "Error: Could not find release for $OS $ARCH"
    exit 1
fi

echo "Downloading $DOWNLOAD_URL..."
curl -sfL "$DOWNLOAD_URL" -o "/tmp/$BINARY.$FORMAT"

echo "Installing to /usr/local/bin..."
tar -xzf "/tmp/$BINARY.$FORMAT" -C /tmp
sudo mv "/tmp/$BINARY" /usr/local/bin/
chmod +x "/usr/local/bin/$BINARY"

echo "Successfully installed repo2page!"
repo2page --version
