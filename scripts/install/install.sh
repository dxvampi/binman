# Copyright (C) 2026 dxvampi
# 
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as
# published by the Free Software Foundation, either version 3 of the
# License, or (at your option) any later version.
# 
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU Affero General Public License for more details.
# 
# You should have received a copy of the GNU Affero General Public License
# along with this program.  If not, see <https://www.gnu.org/licenses/>.

#!/usr/bin/env bash

set -e

echo "STEP 0 -> CHECKING DEPENDENCIES"
if ! command -v git &> /dev/null; then
    echo "Git is not installed on your system or is not in PATH"
    exit 1
fi

if ! command -v go &> /dev/null; then
    echo "Go is not installed on your system or is not in PATH"
    exit 1
fi

GO_VERSION="$(go env GOVERSION | sed 's/go//')"
MIN_VERSION="1.24"

if [ "$(printf '%s\n%s' "$MIN_VERSION" "$GO_VERSION" | sort -V | head -n1)" != "$MIN_VERSION" ]; then
    echo "Go version $GO_VERSION is installed, but $MIN_VERSION or higher is required"
    exit 1
fi

echo "STEP 1 -> CLONING REPO"

git clone -b main https://codeberg.org/dxvampi/binman.git binman-tmp
trap 'echo "Cleaning up..."; rm -rf binman-tmp' EXIT
cd binman-tmp

echo "STEP 2 -> BUILDING"

go build -ldflags="-s -w" -o binman .
go install -ldflags="-s -w" .

echo "STEP 3 -> CONFIGURING PATH"

GOPATH_BIN="$(go env GOPATH)/bin"

case ":$PATH:" in
    *":$GOPATH_BIN:"*)
        echo "Go bin already in PATH."
        ;;
    *)
        echo "Adding Go bin to PATH..."
        for rc in "$HOME/.zshrc" "$HOME/.bashrc"; do
            if [ -f "$rc" ]; then
                if ! grep -q "$GOPATH_BIN" "$rc"; then
                    echo "export PATH=\$PATH:$GOPATH_BIN" >> "$rc"
                    echo "Configured in: $rc"
                fi
            fi
        done
        ;;
esac

echo "binman successfully installed"