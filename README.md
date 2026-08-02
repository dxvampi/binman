# Binman

A simple binary version manager. Originally built because I wanted to manage multiple Java versions on the same Linux system for Minecraft servers, and was kind of tedious.

## Installation

### Linux / macOS

```bash
curl -fsSL https://raw.githubusercontent.com/dxvampi/binman/main/scripts/install/install.sh | bash
```

### Windows

```powershell
irm https://raw.githubusercontent.com/dxvampi/binman/main/scripts/install/install.ps1 | iex
```

> **Note (Windows):** since the binary isn't signed, SmartScreen may show a warning the first time you run it. Click "More info" and then "Run anyway".

## Usage

### Add an alias

```bash
binman config
```

Prompts for an alias and a path to a binary. You can add more than one at once

### List aliases

```bash
binman list
```

### Run a binary by alias

```bash
binman -b <alias> [args...]
```

### Remove a binary

```bash
binman remove
```

### Get the path of a binary

```bash
binman which <alias>
```

Prints the path for the given alias. Useful for scripting, e.g. `$(binman which java17)`.

### Check for updates

\`\`\`bash
binman update
\`\`\`

Checks GitHub for a newer release and prompts to install it. BinmanX also checks for updates automatically in the background (at most once every 24 hours) and will prompt you after any command finishes if a new version is found.

## Building from source

Requires Go 1.24 or newer.

\`\`\`bash
git clone https://github.com/dxvampi/binman.git
cd binman
go build -o binman
\`\`\`

To cross-compile for other platforms/architectures, see `scripts/build/linuxonly-build-all.sh`.

## License
BinmanX is licensed under the [GNU AGPLv3](LICENSE)