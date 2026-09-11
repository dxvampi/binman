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

echo "STEP 1 -> CLONING REPO"

git clone -b main https://codeberg.org/dxvampi/binman.git binman-tmp
cd binman-tmp

echo "STEP 2 -> BUILDING"

go build -ldflags="-s -w" -o binman .
go install -ldflags="-s -w" .

echo "STEP 3 -> DELETING TEMPORAL FILES"

cd ..
rm -rf binman-tmp

echo "STEP 4 -> CONFIGURING PATH"

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