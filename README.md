# repo2page

**repo2page** is a developer tool that converts an entire GitHub or local repository into a single, portable, structure-preserving document.

It is designed to be **simple**, **fast**, and **universally accessible**, distributed as a single binary with **zero runtime dependencies**.

## Features

- **Local & GitHub Support:** Seamlessly convert local directories or remote GitHub repositories.
- **Improved Performance:** Optimized GitHub fetching engine (via tarball download) for 10-100x faster processing.
- **Multiple Formats:** Export to Markdown (`.md`), HTML (`.html`), JSON (`.json`), or Plain Text (`.txt`).
- **Structure Preservation:** Maintains the directory tree structure in the output.
- **Smart Requires:** Automatically respects `.gitignore` and comes with sensible default ignore patterns (e.g., node_modules, .git).
- **Zero Dependencies:** Compiles to a single binary. No Node.js, Python, or JVM required.
- **Offline Capable:** Works offline for local repositories.

## Installation

### One-Line Script (macOS/Linux)
```bash
curl -sfL https://raw.githubusercontent.com/imsncht/repo2page/main/scripts/install.sh | sh
```

### Homebrew (macOS/Linux)
```bash
brew tap imsncht/homebrew-tap
brew install repo2page
```
OR
```bash
brew install imsncht/homebrew-tap/repo2page
```

### Scoop (Windows)
```bash
scoop bucket add imsncht https://github.com/imsncht/scoop-bucket.git
scoop install repo2page
```

### Winget (Windows)
```bash
winget install repo2page
```

### From Source

Ensure you have Go 1.23+ installed.

```bash
go install github.com/imsncht/repo2page/cmd/repo2page@latest
```

*(Note: Official binaries are available on the Releases page)*

## Usage

### Basic Usage

**Convert a GitHub repository:**
```bash
repo2page user/repo
# Output: repo.md
```

**Convert a local directory:**
```bash
repo2page ./my-project
# Output: my-project.md
```

### CLI Options

```bash
repo2page [flags] <source>
```

| Flag | Description | Default |
|------|-------------|---------|
| `--format`, `-f` | Output format (`md`, `html`, `txt`, `json`) | `md` |
| `--output`, `-o` | Output file path (defaults to repo name + ext) | Auto |
| `--branch`, `-b` | GitHub branch to fetch | `main` |
| `--summary`, `-s` | Summary mode: only output tree structure | `false` |
| `--exclude`, `-e` | Additional exclude patterns (glob) | `[]` |
| `--max-file-size` | Max size per file in KB | `500` |
| `--no-tree` | Disable directory tree visualization | `false` |
| `--json` | Output stats/structure in JSON format | `false` |
| `--version`, `-v` | Show version information | `false` |

### Examples

**Convert specific branch to HTML:**
```bash
repo2page --format html --branch develop --output dev-docs.html user/repo
```

**Exclude specific directories:**
```bash
repo2page --exclude "dist/**" --exclude "test/**" ./my-local-project
```

**Generate a summary of the repo structure only:**
```bash
repo2page --summary user/repo
```

## Configuration

`repo2page` respects your `.gitignore` file automatically. It also includes comprehensive default ignore patterns for common build artifacts, lock files, and system files (e.g., `.DS_Store`, `node_modules/`, `vendor/`).

## Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on how to get started.

## License

[MIT License](LICENSE)
