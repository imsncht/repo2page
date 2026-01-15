# Usage Guide

Learn how to use `repo2page` to convert repositories into documents efficiently.

## Basic Usage

### Convert a GitHub Repository
Provide the `user/repo` handle. The tool will automatically fetch the default branch.
```bash
repo2page imsncht/repo2page
```

### Convert a Local Directory
Point to any local folder on your system.
```bash
repo2page ./my-project
```

---

## CLI Options

| Flag | Description | Default |
|------|-------------|---------|
| `--format`, `-f` | Output format (`md`, `html`, `txt`, `json`) | `md` |
| `--output`, `-o` | Specific output file path | Auto |
| `--branch`, `-b` | GitHub branch to fetch | `main` (auto) |
| `--commit`, `-c` | Specific commit hash (GitHub only) | Latest |
| `--summary`, `-s` | Output only the directory tree structure | `false` |
| `--exclude`, `-e` | Glob patterns to ignore (repeatable) | `[]` |
| `--max-file-size` | Max size per file in KB | `500` |
| `--no-tree` | Disable the visual tree in the output | `false` |
| `--json` | Output machine-readable JSON summary | `false` |

---

## Advanced Examples

### Exporting to HTML with custom excludes
```bash
repo2page -f html -e "tests/**" -e "**.log" imsncht/repo2page
```

### Generating a quick structure summary
```bash
repo2page --summary ./large-project
```

### Fetching a specific branch
```bash
repo2page --branch develop imsncht/repo2page
```

---

## Smart Filtering
`repo2page` automatically respects your `.gitignore` file. It also includes sensible defaults to skip:
- `.git/` directories
- `node_modules/`
- Binary files and images
- Hidden system files (e.g., `.DS_Store`)
