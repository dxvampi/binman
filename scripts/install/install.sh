#!/bin/bash
set -e

REPO="dxvampi/binman"
INSTALL_DIR="$HOME/.local/bin"
os=$(uname -s | tr '[:upper:]' '[:lower:]')
arch=$(uname -m)

case $arch in
x86_64) arch="amd64" ;;
aarch64|arm64) arch="arm64" ;;
*) echo "Unsupported architecture: $arch"; exit 1 ;;
esac

binary_name="binman-$os-$arch"
url="https://github.com/$REPO/releases/latest/download/$binary_name"

mkdir -p "$INSTALL_DIR"
tmp_file="$(mktemp "$INSTALL_DIR/binman.XXXXXX")"
curl -L -o "$tmp_file" "$url"
chmod +x "$tmp_file"
mv "$tmp_file" "$INSTALL_DIR/binman"

case "$SHELL" in
*/zsh) shell_rc="$HOME/.zshrc" ;;
*) shell_rc="$HOME/.bashrc" ;;
esac

if ! echo "$PATH" | grep -q "$INSTALL_DIR"; then
echo "export PATH=\"\$PATH:$INSTALL_DIR\"" >> "$shell_rc"
echo "Restart your terminal or run: source $shell_rc"
fi

if [ "$os" == "darwin" ]; then
xattr -d com.apple.quarantine "$INSTALL_DIR/binman" 2>/dev/null || true
fi