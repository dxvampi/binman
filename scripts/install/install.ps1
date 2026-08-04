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

Write-Host "STEP 1 -> CLONING REPO"
git clone $repoUrl $cloneDir
Set-Location $cloneDir

Write-Host "STEP 2 -> BUILDING"
go build -ldflags="-s -w" -o binman.exe .
go install -ldflags="-s -w" .

echo "STEP 3 -> DELETING TEMPORAL FILES"
Set-Location ..
Remove-Item -Recurse -Force $cloneDir

Write-Host "binman installed successfully"