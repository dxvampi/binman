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