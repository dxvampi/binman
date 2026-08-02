$ErrorActionPreference = "Stop"

$repo = "dxvampi/binman"
$installDir = "$env:LOCALAPPDATA\binman"

$arch = if ([Environment]::Is64BitOperatingSystem) { "amd64" } else { "386" }

New-Item -ItemType Directory -Force -Path $installDir | Out-Null
$url = "https://github.com/$repo/releases/latest/download/binman-windows-$arch.exe"
$dest = "$installDir\binman.exe"

Write-Host "Downloading binman..."
Invoke-WebRequest -Uri $url -OutFile $dest

$currentPath = [Environment]::GetEnvironmentVariable("PATH", "User")
if ($currentPath -notlike "*$installDir*") {
    [Environment]::SetEnvironmentVariable("PATH", "$currentPath;$installDir", "User")
    Write-Host "Added $installDir to PATH. Restart your terminal."
}
Write-Host "binman installed successfully!"