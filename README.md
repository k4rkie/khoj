# KHOJ

A lightweight, CLI tool to search for keywords in your codebase. Designed for speed, efficiency, and respect for your project's existing configuration.

"Khoj" (खोज) is a Nepali word for "Search."

> [!WARNING]
> This tool still under construction!

## Installation

If you have Go installed, you can install Khoj directly:

```bash
go install github.com/k4rkie/khoj@latest
```

Alternatively, clone the repo and build from source:

```bash
git clone https://github.com/k4rkie/khoj.git
cd khoj
go build -o khoj main.go
```

## Usage

Khoj is designed to be simple. By default, it searches for TODO in the current directory.

#### Basic Usage:

```bash
# Search for the default keyword "TODO" in the current directory
khoj
# Search for a specific keyword in a specific directory
khoj -d ./my-project -k "FIXME"
```

#### Piping (UNIX):

```bash
# Search for TODOs and pipe to grep to find only those in auth files
khoj -k "TODO" | grep "auth"
```

### Respecting Your Workspace

`khoj` is designed to keep your search results clean and relevant. It automatically detects and respects the `.gitignore` file in your project's root directory.

- **Automatic Filtering:** If you have ignored directories (like `node_modules/`, `venv/`, or `.git/`) in your `.gitignore`, `khoj` will skip them entirely.
- **Customization:** If you want `khoj` to search a directory that is currently ignored, simply remove the pattern from your `.gitignore` or use a different search directory.
