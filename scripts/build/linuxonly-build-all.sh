#!/bin/bash
# scripts/build/build-all.sh
set -e

VERSION=$1
mkdir -p dist

targets=(
    "linux amd64"
    "linux arm64"
    "windows amd64"
    "windows arm64"
    "darwin amd64"
    "darwin arm64"
)

for target in "${targets[@]}"; do
  os=$(echo $target | cut -d' ' -f1)
  arch=$(echo $target | cut -d' ' -f2)
  output="dist/binman-$os-$arch"
  if [ "$os" == "windows" ]; then
    output="$output.exe"
  fi
  echo "Building $output..."
  GOOS=$os GOARCH=$arch go build -ldflags "-X binman/internal/version.Version=$VERSION" -o $output
done