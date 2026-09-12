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

$ErrorActionPreference = "Stop"

$repoUrl = "https://codeberg.org/dxvampi/binman.git"
$cloneDir = "binman-install-tmp"
$minVersion = [version]"1.24"

Write-Host "STEP 0 -> CHECKING DEPENDENCIES"

if (-not (Get-Command git -ErrorAction SilentlyContinue)) {
    Write-Host "Git is not installed on your system or is not in PATH"
    exit 1
}

if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    Write-Host "Go is not installed on your system or is not in PATH"
    exit 1
}

$goVersionRaw = (go env GOVERSION) -replace '^go', ''
$goVersion = [version]$goVersionRaw

if ($goVersion -lt $minVersion) {
    Write-Host "Go version $goVersionRaw is installed, but $minVersion or higher is required"
    exit 1
}

Write-Host "STEP 1 -> CLONING REPO"

try {
    Set-Location $cloneDir

    Write-Host "STEP 2 -> BUILDING"
    go build -ldflags="-s -w" -o binman.exe .
    go install -ldflags="-s -w" .

    Write-Host "binman installed successfully"
}
finally {
    Write-Host "Cleaning up..."
    Set-Location ..
    Remove-Item -Recurse -Force $cloneDir
}