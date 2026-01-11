# repo2page – Product Requirements Document (PRD)
**Version 2.0 - Universal Distribution Edition**

**Status:** Final, implementation-ready  
**Last Updated:** January 2026

---

## Document Purpose
This document is the single authoritative specification for repo2page. It combines product vision, architecture, technical specifications, CLI UX, deployment strategy, acceptance criteria, and roadmap into one comprehensive specification.

---

## 1. Overview

**Product Name:** repo2page  
**Category:** Developer Tool / Documentation Utility  
**Primary Interfaces:** CLI Tool, Web Application  
**Core Capability:** Convert an entire GitHub or local repository into a single, simple, portable page while preserving structure and meaning.

**Value Proposition:**  
repo2page turns fragmented repositories into a unified artifact that can be shared, attached, archived, reviewed, or ingested by AI systems—accessible to everyone, everywhere, regardless of platform or connectivity.

---

## 2. Problem Statement

### Current Pain Points
Modern codebases are:
- Distributed across dozens or hundreds of files
- Dependent on GitHub or similar platforms for access
- Difficult to share with non-technical users
- Inconvenient to attach in emails, applications, audits, or submissions
- Inaccessible without internet connectivity

### Existing Solution Failures
Current approaches fail because they:
- Strip structure (README-only exports)
- Require GitHub access and internet connectivity (links)
- Are cumbersome and not universally accessible (ZIP archives)
- Are not readable or portable (raw dumps)
- Require runtime dependencies (Node.js, Python, etc.)

### Core Problem
**There is no fast, reliable, universally accessible way to convert a repository into a single, readable, self-contained document that works offline and requires zero dependencies.**

---

## 3. Goals & Non-Goals

### 3.1 Goals
✅ Convert a repository into one file  
✅ Preserve directory structure and file boundaries  
✅ Support both online and offline operation  
✅ Work with GitHub repositories (when online) and local repositories (always)  
✅ Produce deterministic, reproducible output  
✅ Distribute as zero-dependency standalone binaries for all major platforms  
✅ Keep the tool simple, fast, and predictable  
✅ Ensure universal accessibility across Windows, macOS, and Linux  
✅ Support multiple installation methods (package managers, direct download, scripts)

### 3.2 Non-Goals
❌ Code execution or analysis  
❌ IDE replacement  
❌ Real-time repo syncing  
❌ Deep semantic understanding of code (initially)  
❌ Web-based Git hosting (this is a converter, not a platform)  
❌ Dependency on runtime environments (Node.js, Python, JVM, etc.)

---

## 4. Target Users

### Primary Users
- **Developers** (job applications, code reviews, handoffs, portfolio sharing)
- **Open-source maintainers** (documentation, archival)
- **Students and researchers** (project submissions, academic work)
- **Indie hackers** (rapid prototyping, sharing MVPs)

### Secondary Users
- **Recruiters and technical reviewers** (evaluating candidates)
- **Legal/compliance teams** (code audits, IP review)
- **Technical writers** (documentation preparation)
- **AI practitioners** (preparing structured context for LLMs)

### User Environment Requirements
- Mixed connectivity (online/offline scenarios)
- Cross-platform environments (Windows, macOS, Linux)
- Varying technical expertise (CLI-comfortable to GUI-only)
- Enterprise and personal use cases

---

## 5. Core Use Cases

1. **Attach a full repository to an email or job application**
2. **Submit code where only PDFs or text files are accepted**
3. **Create an offline-readable snapshot of a project**
4. **Feed structured repository context into LLMs**
5. **Share code with non-GitHub users or users without internet access**
6. **Archive a repository at a specific commit or point in time**
7. **Work on repositories during flights, commutes, or in low-connectivity areas**
8. **Review code without requiring GitHub access or authentication**

---

## 6. Product Architecture

### 6.1 Three-Layer Architecture

```
┌─────────────────────────────────────────┐
│         CLI Interface (Primary)         │
│    - Cross-platform binary              │
│    - Zero external dependencies        │
│    - Flags and options handling         │
└─────────────────────────────────────────┘
                    │
                    ▼
┌─────────────────────────────────────────┐
│      Core Engine (Shared Library)       │
│    - Repository loading & parsing       │
│    - Tree building & traversal          │
│    - Format conversion                  │
│    - Ignore rules & filtering           │
└─────────────────────────────────────────┘
                    │
                    ▼
┌─────────────────────────────────────────┐
│    Web Wrapper (Optional/Future)        │
│    - Simple landing page                │
│    - Browser-based conversion           │
│    - Ephemeral processing               │
└─────────────────────────────────────────┘
```

### 6.2 Implementation Decisions (Finalized)

**Primary Language:** Go  
**Rationale:**
- Zero-dependency standalone binaries
- Excellent cross-compilation support
- Fast compilation and execution (<10ms startup)
- Small binary size (5-15MB)
- Robust standard library
- Excellent GitHub API libraries
- Strong CLI tooling ecosystem (cobra, viper)
- Native support for concurrent operations

**Alternative Considered:** TypeScript/Node.js  
**Rejected Because:**
- Requires Node.js runtime installation
- Larger bundle size (~50MB+)
- Slower startup time (~100-500ms)
- Additional user friction

**Runtime Requirements:** None (compiled binary)  
**Minimum Supported Versions:**
- Windows 10 or later
- macOS 10.13 (High Sierra) or later
- Linux kernel 3.2+ (glibc 2.17+)

**These decisions are binding for MVP and all future versions.**

---

## 7. Core Engine (repo2page-core)

### 7.1 Responsibilities

**The core engine MUST:**
- Accept a repository source (local path or GitHub URL)
- Auto-detect source type (local vs. GitHub)
- Auto-detect connectivity status (online vs. offline)
- Resolve directory structure deterministically
- Apply ignore and filter rules
- Read and normalize file contents
- Serialize the repository into a single-page format
- Return content and metadata
- Handle both online and offline scenarios gracefully

**The core engine MUST NOT:**
- Write files to disk without explicit user consent
- Perform UI or CLI-specific logic beyond the API
- Store user data persistently
- Execute or evaluate code
- Make assumptions about network availability

### 7.2 Public API (Binding)

```go
package repo2page

// Convert is the primary entry point for repository conversion
func Convert(input RepoInput, options ConvertOptions) (*ConvertResult, error)

// AutoDetectSource determines if input is local path or GitHub URL
func AutoDetectSource(input string) (SourceType, error)

// IsOnline checks if GitHub API is accessible
func IsOnline() bool
```

### 7.3 Type Definitions

```go
type SourceType int

const (
    SourceLocal SourceType = iota
    SourceGitHub
)

type RepoInput struct {
    Source     SourceType
    PathOrURL  string
    Branch     string  // optional, default: "main"
    Commit     string  // optional, specific commit hash
}

type ConvertOptions struct {
    Format              OutputFormat  // md, html, txt (default: md)
    IncludeTree         bool          // default: true
    IncludeReadmeFirst  bool          // default: true
    ExcludePatterns     []string      // custom ignore patterns
    MaxFileSizeKB       int           // default: 500
    SummaryMode         bool          // default: false
    OutputPath          string        // optional, auto-generate if empty
}

type OutputFormat string

const (
    FormatMarkdown OutputFormat = "md"
    FormatHTML     OutputFormat = "html"
  FormatText     OutputFormat = "txt"
  FormatJSON     OutputFormat = "json"
)

type ConvertResult struct {
    Content  string
    Format   OutputFormat
    Stats    Statistics
    Warnings []string
}

type Statistics struct {
    FileCount   int
    TotalLines  int
    TotalSizeKB int
    SkippedFiles int
    ProcessingTimeMs int64
}
```

### 7.4 Core Module Boundaries

#### Repository Loader
- **Local Loader:** Reads from filesystem using native file I/O
- **GitHub Loader:** Fetches via GitHub API (requires connectivity)
- **Authentication (optional):** Support using an environment variable `GITHUB_TOKEN` or `GITHUB_API_TOKEN` to increase API rate limits for public requests. Tokens MUST never be logged, written to disk, or stored persistently. Note: full private-repository access (authenticated private repo reads) is planned for v0.4; for MVP the token is used only to raise rate limits for public repo fetches.
- **Auto-detection:** Determines source type from input string
- **Connectivity Check:** Verifies online status before GitHub operations
- **No formatting or filtering beyond path resolution**

#### Source Detection Logic
```
Input Analysis:
├── Starts with http:// or https:// → GitHub
├── Matches pattern "user/repo" → GitHub (shorthand)
├── Path exists on filesystem → Local
└── Otherwise → Error: Invalid source
```

#### Connectivity Detection
```
Online Detection:
├── Attempt DNS lookup for github.com
├── Timeout: 2 seconds
├── Success → Online
└── Failure → Offline
```

#### Tree Resolver
- Builds in-memory tree structure
- Applies deterministic ordering (see §7.5)
- Maintains parent-child relationships
- Calculates depth and structure metadata

#### Ignore Engine
- Applies default ignore rules (see §11)
- Merges with user-provided patterns
- Respects .gitignore files when present
- Pattern matching using glob syntax

#### File Reader & Normalizer
- Reads file contents with encoding detection
- Normalizes line endings (CRLF → LF)
- Skips binary files automatically
- Enforces size limits
- Handles encoding errors gracefully

#### Language Detection
- Detects language by file extension
- Maps to common language identifiers
- Used for syntax highlighting hints
- Fallback to "text" for unknown types

#### Formatter
- **Markdown Formatter:** Code blocks with language tags, headings
- **HTML Formatter:** Self-contained with inline CSS, syntax highlighting
- **Plain Text Formatter:** Minimal formatting, clear separators

#### Page Assembler
- Combines metadata, tree, and file contents
- Enforces README-first rule (if present)
- Generates table of contents
- Adds headers and footers
- Inserts statistics and warnings

### 7.5 Deterministic Ordering Rule (Binding)

**Traversal Algorithm:**
1. **Depth-first traversal** of directory tree
2. **Alphabetical ordering** within each directory level
3. **Case-insensitive sorting** (cross-platform consistency)
4. **Directories before files** at same level
5. **Root-level README rendered first** if present (any case: README.md, readme.txt, etc.)

**Sorting Priority:**
```
Priority 1: README files (root level only)
Priority 2: Directories (alphabetically)
Priority 3: Files (alphabetically)
```

This ensures:
- Same repository always produces identical output
- Predictable structure across platforms
- Optimal reading order (documentation first)

### 7.6 Online/Offline Behavior

| Scenario | Source Type | Connectivity | Behavior |
|----------|-------------|--------------|----------|
| 1 | Local | Offline | ✅ Works - reads from filesystem |
| 2 | Local | Online | ✅ Works - reads from filesystem |
| 3 | GitHub | Offline | ❌ Error - requires internet connection |
| 4 | GitHub | Online | ✅ Works - fetches via GitHub API |

**Error Messages:**
- Offline + GitHub: "GitHub repositories require an internet connection. Try using a local clone instead."
- Invalid source: "Unable to determine source type. Provide a local path or GitHub URL."

### 7.7 Error Handling Rules (Core)

**Classification:**
- **Recoverable Issues** → Warnings (logged, processing continues)
- **Hard Failures** → Typed errors (processing stops, clear message)

**Recoverable Issues (Warnings):**
- File too large (skipped)
- Binary file encountered (skipped)
- Encoding error in specific file (skipped)
- Permission denied on specific file (skipped)

**Hard Failures (Errors):**
- Invalid input source
- Repository fetch/load failure
- No readable files found
- Unsupported output format
- Network timeout (GitHub sources)
- Invalid credentials (if authentication required)

**Error Message Requirements:**
- Human-readable (no stack traces to end users)
- Actionable (suggest next steps)
- Context-aware (explain what was attempted)
- Concise (1-3 lines maximum)

---

## 8. CLI Wrapper (repo2page)

### 8.1 Command Signature

```bash
repo2page <source> [options]
```

**Source Formats:**
- Local path: `./my-project`, `/home/user/code/app`, `C:\Projects\app`
- GitHub URL: `https://github.com/user/repo`
- GitHub shorthand: `user/repo` (auto-expands to full URL)

### 8.2 CLI Flags

| Flag | Description | Default | Type |
|------|-------------|---------|------|
| `--format` | Output format (md/html/txt) | `md` | string |
| `--output` | Output file name | auto-generated | string |
| `--branch` | Git branch | `main` | string |
| `--commit` | Specific commit hash | (none) | string |
| `--exclude` | Additional ignore patterns | (uses defaults) | []string |
| `--summary` | Summary mode (high-level only) | `false` | bool |
| `--no-tree` | Disable directory tree in output | `false` | bool |
| `--max-file-size` | Max file size to include (KB) | `500` | int |
| `--json` | Emit machine-readable JSON summary (for automation) | `false` | bool |
| `--help` | Show help message | - | - |
| `--version` | Show version information | - | - |

### 8.3 Usage Examples

```bash
# Basic: Convert local repository to Markdown
repo2page ./my-project

# Convert GitHub repository
repo2page https://github.com/torvalds/linux

# Convert GitHub repository (shorthand)
repo2page torvalds/linux

# Specific branch and format
repo2page user/repo --branch develop --format html

# Specific commit
repo2page ./project --commit abc123def --output snapshot.md

# Custom output with exclusions
repo2page . --output my-code.html --exclude "*.log" --exclude "tmp/*"

# Summary mode (structure only, no full file contents)
repo2page ./big-project --summary

# Plain text without tree
repo2page ./app --format txt --no-tree
```

### 8.4 Exit Codes

| Code | Meaning | Details |
|------|---------|---------|
| 0 | Success | Conversion completed successfully |
| 1 | Invalid arguments | Flag parsing error or invalid input |
| 2 | Source failure | Repository fetch/load failure |
| 3 | Conversion failure | Processing error during conversion |
| 4 | Output failure | Unable to write output file |
| 5 | Network error | GitHub API unreachable (online mode) |
| 99 | Unexpected error | Internal error (should be rare) |

### 8.5 CLI Output Format

**Standard Output (Success):**
```
Converting repository...
✓ Loaded 127 files
✓ Generated Markdown (543 KB)
✓ Saved to: my-project.md

Statistics:
  Files: 127
  Lines: 15,432
  Size: 543 KB
  Skipped: 3 (too large)
```

**Standard Error (Failure):**
```
✗ Error: Repository not found
  → Check that 'user/repo' exists and is public
  → Run with --help for usage information
```

**Machine-readable Output (when `--json` used):**
```json
{
  "status": "success",
  "files": 127,
  "lines": 15432,
  "size_kb": 543,
  "skipped": 3,
  "output": "my-project.md",
  "warnings": ["Skipped large file: assets/video.mp4 (2.3 MB > 500 KB)"]
}
```

### 8.6 Error Text Philosophy

All error messages must:
1. **Be short** (1-2 lines maximum)
2. **Explain the failure** clearly
3. **Suggest next action** or workaround
4. **Avoid technical jargon** when possible

**Examples:**

❌ Bad:
```
Error: syscall.Errno 0x2
```

✅ Good:
```
✗ Directory not found: ./my-project
  → Check the path and try again
```

---

## 9. Deployment & Distribution Strategy

### 9.1 Universal Distribution Philosophy

**Core Principle:** repo2page must be installable by anyone, on any platform, using their preferred method—with zero runtime dependencies.

**Target Platforms:**
- Windows (10, 11, Server 2016+)
- macOS (10.13+, Intel and Apple Silicon)
- Linux (major distributions: Ubuntu, Debian, Fedora, Arch, Alpine)

### 9.2 Distribution Channels (Priority Order)

#### Priority 1: Package Managers (Recommended)

**macOS:**
```bash
brew install repo2page
```

**Windows:**
```bash
# Chocolatey
choco install repo2page

# Scoop
scoop install repo2page

# Winget (Microsoft)
winget install repo2page
```

**Linux:**
```bash
# Homebrew (cross-platform)
brew install repo2page

# Snap (Ubuntu/universal)
snap install repo2page

# APT (Debian/Ubuntu)
sudo apt install repo2page

# DNF (Fedora/RHEL)
sudo dnf install repo2page

# Pacman (Arch)
sudo pacman -S repo2page
```

#### Priority 2: Direct Binary Downloads

**GitHub Releases:**
- `repo2page-v0.1.0-windows-amd64.exe`
- `repo2page-v0.1.0-darwin-amd64` (Intel Mac)
- `repo2page-v0.1.0-darwin-arm64` (Apple Silicon)
- `repo2page-v0.1.0-linux-amd64`
- `repo2page-v0.1.0-linux-arm64`

**Installation:**
1. Download appropriate binary for platform
2. Move to PATH location
3. Make executable (Unix: `chmod +x`)
4. Run `repo2page --version` to verify

#### Priority 3: One-Line Install Scripts

**Unix (macOS/Linux):**
```bash
curl -sSL https://install.repo2page.dev/install.sh | sh
```

**Windows (PowerShell):**
```powershell
irm https://install.repo2page.dev/install.ps1 | iex
```

**Script Behavior:**
- Detects OS and architecture automatically
- Downloads latest release from GitHub
- Installs to appropriate system location
- Updates PATH if necessary
- Verifies installation

#### Priority 4: Container Image (Optional)

```bash
# Docker
docker run --rm -v $(pwd):/repo repo2page/cli /repo

# Podman
podman run --rm -v $(pwd):/repo repo2page/cli /repo
```

### 9.3 Build & Release Automation

**Tool:** GoReleaser

**Configuration:** `.goreleaser.yml`
```yaml
project_name: repo2page

before:
  hooks:
    - go mod tidy
    - go test ./...

builds:
  - id: repo2page
    main: ./cmd/repo2page
    binary: repo2page
    env:
      - CGO_ENABLED=0
    goos:
      - linux
      - windows
      - darwin
    goarch:
      - amd64
      - arm64
    ignore:
      - goos: windows
        goarch: arm64
    ldflags:
      - -s -w
      - -X main.version={{.Version}}
      - -X main.commit={{.Commit}}
      - -X main.date={{.Date}}

archives:
  - id: repo2page-archive
    format: tar.gz
    name_template: >-
      {{ .ProjectName }}_
      {{- .Version }}_
      {{- .Os }}_
      {{- .Arch }}
    format_overrides:
      - goos: windows
        format: zip
    files:
      - README.md
      - LICENSE

checksum:
  name_template: 'checksums.txt'

changelog:
  sort: asc
  filters:
    exclude:
      - '^docs:'
      - '^test:'

brews:
  - repository:
      owner: yourusername
      name: homebrew-tap
    homepage: https://repo2page.dev
    description: Convert repositories into single, portable documents
    license: MIT
    install: |
      bin.install "repo2page"

chocolateys:
  - owners: Your Name
    title: repo2page
    project_url: https://github.com/yourusername/repo2page
    license_url: https://github.com/yourusername/repo2page/blob/main/LICENSE
    require_license_acceptance: false
    description: Convert repositories into single, portable documents

scoops:
  - repository:
      owner: yourusername
      name: scoop-bucket
    homepage: https://repo2page.dev
    description: Convert repositories into single, portable documents
    license: MIT

snapcrafts:
  - summary: Convert repositories into single, portable documents
    description: |
      repo2page converts entire repositories (GitHub or local) into single,
      portable documents in Markdown, HTML, or plain text format.
    grade: stable
    confinement: strict
```

**Release Process:**
```bash
# 1. Create and push tag
git tag -a v0.1.0 -m "Initial release"
git push origin v0.1.0

# 2. GoReleaser builds everything automatically
goreleaser release --clean

# 3. Artifacts generated:
#    - Binaries for all platforms
#    - Homebrew formula
#    - Chocolatey package
#    - Scoop manifest
#    - Snap package
#    - GitHub Release page
```

### 9.4 Installation Script Specifications

**install.sh (Unix)**
```bash
#!/bin/bash
set -e

REPO="yourusername/repo2page"
INSTALL_DIR="/usr/local/bin"

echo "Installing repo2page..."

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64) ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
    *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

# Get latest version
LATEST=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST" ]; then
    echo "Failed to fetch latest version"
    exit 1
fi

echo "Latest version: $LATEST"

# Download
FILENAME="repo2page_${LATEST}_${OS}_${ARCH}.tar.gz"
URL="https://github.com/$REPO/releases/download/$LATEST/$FILENAME"

echo "Downloading from: $URL"
curl -sL "$URL" | tar xz -C /tmp

# Install
if [ -w "$INSTALL_DIR" ]; then
    mv /tmp/repo2page "$INSTALL_DIR/repo2page"
    chmod +x "$INSTALL_DIR/repo2page"
else
    sudo mv /tmp/repo2page "$INSTALL_DIR/repo2page"
    sudo chmod +x "$INSTALL_DIR/repo2page"
fi

echo "✓ repo2page installed successfully!"
echo "Run: repo2page --help"
```

**install.ps1 (Windows)**
```powershell
$ErrorActionPreference = 'Stop'

$repo = "yourusername/repo2page"
$installDir = "$env:LOCALAPPDATA\Programs\repo2page"

Write-Host "Installing repo2page..." -ForegroundColor Cyan

# Get latest release
try {
    $release = Invoke-RestMethod "https://api.github.com/repos/$repo/releases/latest"
    $version = $release.tag_name
    $asset = $release.assets | Where-Object { $_.name -like "*windows_amd64.zip" }
    
    if (-not $asset) {
        throw "Windows binary not found in release"
    }
} catch {
    Write-Host "Failed to fetch latest release: $_" -ForegroundColor Red
    exit 1
}

Write-Host "Latest version: $version" -ForegroundColor Green

# Download
$zipPath = "$env:TEMP\repo2page.zip"
Write-Host "Downloading..."
Invoke-WebRequest $asset.browser_download_url -OutFile $zipPath

# Extract
Write-Host "Installing..."
New-Item -ItemType Directory -Force -Path $installDir | Out-Null
Expand-Archive -Path $zipPath -DestinationPath $installDir -Force

# Add to PATH
$userPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($userPath -notlike "*$installDir*") {
    Write-Host "Adding to PATH..."
    [Environment]::SetEnvironmentVariable("Path", "$userPath;$installDir", "User")
    $env:Path = "$env:Path;$installDir"
}

# Cleanup
Remove-Item $zipPath -Force

Write-Host "✓ repo2page installed successfully!" -ForegroundColor Green
Write-Host "Run: repo2page --help" -ForegroundColor Yellow
Write-Host "(You may need to restart your terminal)" -ForegroundColor Gray
```

### 9.5 Distribution Matrix

| Platform | Installation Method | Command | Auto-Updates | Recommended |
|----------|-------------------|---------|--------------|-------------|
| macOS | Homebrew | `brew install repo2page` | ✅ | ⭐ Primary |
| macOS | Direct Binary | Download + install manually | ❌ | Fallback |
| macOS | Install Script | `curl ... \| sh` | ❌ | Quick start |
| Windows | Chocolatey | `choco install repo2page` | ✅ | ⭐ Primary |
| Windows | Scoop | `scoop install repo2page` | ✅ | ⭐ Primary |
| Windows | Winget | `winget install repo2page` | ✅ | Recommended |
| Windows | PowerShell Script | `irm ... \| iex` | ❌ | Quick start |
| Windows | Direct Binary | Download `.exe` | ❌ | Fallback |
| Linux | Homebrew | `brew install repo2page` | ✅ | ⭐ Primary |
| Linux | Snap | `snap install repo2page` | ✅ | Ubuntu users |
| Linux | APT | `apt install repo2page` | ✅ | Debian/Ubuntu |
| Linux | DNF | `dnf install repo2page` | ✅ | Fedora/RHEL |
| Linux | Install Script | `curl ... \| sh` | ❌ | Quick start |
| Linux | Direct Binary | Download + install manually | ❌ | Fallback |
| All | Docker/Podman | `docker run ...` | N/A | CI/CD |

### 9.6 Binary Naming Convention

```
repo2page_<version>_<os>_<arch>.<extension>

Examples:
- repo2page_v0.1.0_linux_amd64.tar.gz
- repo2page_v0.1.0_darwin_arm64.tar.gz
- repo2page_v0.1.0_windows_amd64.zip
```

---

## 10. Web Wrapper (repo2page-web)

**Status:** Secondary Priority (Post-MVP)

**Purpose:** Provide browser-based access for users who prefer GUI or don't have CLI access.

### 10.1 Core Features
- Accepts GitHub URLs via web form
- Allows format selection (md/html/txt)
- Processes repositories ephemerally (no storage)
- Returns downloadable file immediately
- Shows processing status in real-time

### 10.2 Technical Requirements
- Stateless design (no user accounts)
- No persistent storage
- Rate limiting per IP
- Timeout: 30 seconds max per conversion
- CORS enabled for API access

### 10.3 Implementation Notes
- Can reuse core Go library via WASM or API endpoint
- Hosted on serverless platform (Vercel, Cloudflare Workers, etc.)
- Simple landing page with form
- Optional: Public API endpoint for programmatic access

---

## 11. Output Format Specifications

### 11.1 Common Structure (All Formats)

```
1. Document Header
   - Repository name and source
   - Conversion timestamp
   - Statistics summary

2. Directory Tree (optional, via --tree flag)
   - ASCII tree visualization
   - File counts per directory

3. File Contents (one section per file)
   - File path as heading/separator
   - File contents (formatted per output type)
   - Language identification (when applicable)

4. Document Footer
   - Generation info
   - Tool version
```

### 11.2 Markdown Format (Default)

**Structure:**
```markdown
# Repository: my-project

**Source:** /Users/dev/my-project  
**Generated:** 2026-01-03 14:23:45 UTC  
**Files:** 42 | **Lines:** 3,521 | **Size:** 156 KB

## Directory Structure

```
.
├── src/
│   ├── main.go
│   └── utils.go
├── tests/
│   └── main_test.go
└── README.md
```

## Files

### README.md

```markdown
# My Project
...
```

### src/main.go

```go
package main
...
```
```

**Features:**
- Headings for files (H3)
- Code blocks with language tags
- Proper Markdown escaping
- Clickable table of contents (via heading anchors)

### 11.3 HTML Format

**Structure:**
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Repository: my-project</title>
    <style>
        /* Inline CSS for self-contained file */
        body { font-family: -apple-system, sans-serif; max-width: 900px; margin: 40px auto; }
        .header { border-bottom: 2px solid #333; padding-bottom: 20px; }
        .file-section { margin: 40px 0; }
        .file-path { background: #f5f5f5; padding: 10px; font-family: monospace; }
        pre code { display: block; padding: 20px; background: #2d2d2d; color: #f8f8f2; overflow-x: auto; }
        /* Syntax highlighting via highlight.js or similar */
    </style>
</head>
<body>
    <div class="header">
        <h1>Repository: my-project</h1>
        <p><strong>Source:</strong> /Users/dev/my-project</p>
        <p><strong>Generated:</strong> 2026-01-03 14:23:45 UTC</p>
        <p><strong>Stats:</strong> 42 files | 3,521 lines | 156 KB</p>
    </div>
    
    <div class="tree">
        <h2>Directory Structure</h2>
        <pre>...</pre>
    </div>
    
    <div class="files">
        <h2>Files</h2>
        <div class="file-section">
            <div class="file-path">README.md</div>
            <pre><code class="language-markdown">...</code></pre>
        </div>
        <!-- More files... -->
    </div>
</body>
</html>
```

**Features:**
- Single self-contained file (no external dependencies)
- Inline CSS for styling
- Syntax highlighting via embedded library or pre-computed
- Responsive design
- Print-friendly styles

### 11.4 Plain Text Format

**Structure:**
```
================================================================================
REPOSITORY: my-project
================================================================================

Source: /Users/dev/my-project
Generated: 2026-01-03 14:23:45 UTC
Files: 42 | Lines: 3,521 | Size: 156 KB

================================================================================
DIRECTORY STRUCTURE
================================================================================

.
├── src/
│   ├── main.go
│   └── utils.go
├── tests/
│   └── main_test.go
└── README.md

================================================================================
FILES
================================================================================

--- README.md ---

# My Project
...

--- src/main.go ---

package main
...
```

**Features:**
- Minimal formatting
- Clear ASCII separators
- Fixed-width format
- Suitable for terminals and text editors
- Maximum compatibility (readable in any text viewer)
- No special characters beyond basic ASCII

---

## 12. Default Ignore Rules

### 12.1 Core Ignore Patterns (Always Applied)

**Version Control:**
```
.git/
.svn/
.hg/
.bzr/
```

**Dependencies:**
```
node_modules/
vendor/
.venv/
venv/
env/
__pycache__/
```

**Build Artifacts:**
```
dist/
build/
out/
target/
bin/
obj/
*.o
*.so
*.dylib
*.dll
*.exe
```

**IDE & Editor:**
```
.vscode/
.idea/
*.swp
*.swo
*~
.DS_Store
Thumbs.db
```

**Lock Files:**
```
package-lock.json
yarn.lock
Gemfile.lock
Cargo.lock
poetry.lock
pnpm-lock.yaml
```

**Log Files:**
```
*.log
logs/
npm-debug.log*
```

**Temporary Files:**
```
tmp/
temp/
*.tmp
*.cache
```

**Binary & Media (Large Files):**
```
*.jpg
*.jpeg
*.png
*.gif
*.pdf
*.zip
*.tar.gz
*.rar
*.7z
*.mp4
*.mp3
*.mov
*.avi
```

### 12.2 Ignore Priority Rules

1. **Built-in defaults** (lowest priority)
2. **Repository .gitignore** (medium priority)
3. **User --exclude flags** (highest priority)

### 12.3 .gitignore Integration

If a `.gitignore` file exists in the repository root:
- Parse and apply its patterns
- Merge with built-in defaults
- Respect gitignore syntax (negation with `!`, comments with `#`)

### 12.4 Custom Exclusions

Users can add patterns via CLI:
```bash
repo2page ./project --exclude "*.test.js" --exclude "docs/*"
```

Multiple `--exclude` flags are cumulative.

---

## 13. Acceptance Criteria (MVP)

### 13.1 Functional Requirements

✅ **FR-1:** Converts local repositories to single files  
✅ **FR-2:** Converts GitHub repositories to single files (when online)  
✅ **FR-3:** Auto-detects source type (local vs. GitHub)  
✅ **FR-4:** Auto-detects connectivity status  
✅ **FR-5:** Works offline for local repositories  
✅ **FR-6:** Fails gracefully when GitHub source accessed offline  
✅ **FR-7:** Produces deterministic output (same input → same output)  
✅ **FR-8:** Supports Markdown, HTML, and plain text formats  
✅ **FR-9:** All CLI flags function as specified  
✅ **FR-10:** Returns correct exit codes for all scenarios  
✅ **FR-11:** Applies ignore rules correctly  
✅ **FR-12:** Handles large files appropriately (skip with warning)  
✅ **FR-13:** Handles binary files appropriately (skip with warning)  
✅ **FR-14:** Generates accurate statistics  
✅ **FR-15:** README appears first when present  

### 13.2 Non-Functional Requirements

✅ **NFR-1:** Zero runtime dependencies (standalone binary)  
✅ **NFR-2:** Binary size under 20MB  
✅ **NFR-3:** Startup time under 50ms  
✅ **NFR-4:** Conversion completes within 30s for typical repos (<1000 files)  
✅ **NFR-5:** Memory usage under 500MB for typical repos  
✅ **NFR-6:** Cross-platform compatibility (Windows, macOS, Linux)  
✅ **NFR-7:** Installable via major package managers  
✅ **NFR-8:** Clear, actionable error messages  
✅ **NFR-9:** Comprehensive CLI help documentation  
✅ **NFR-10:** Follows semantic versioning  

### 13.3 Distribution Requirements

✅ **DR-1:** Available via Homebrew (macOS/Linux)  
✅ **DR-2:** Available via Chocolatey (Windows)  
✅ **DR-3:** Available via Scoop (Windows)  
✅ **DR-4:** Direct binary downloads for all platforms  
✅ **DR-5:** One-line install scripts (Unix + Windows)  
✅ **DR-6:** Automated build and release pipeline  
✅ **DR-7:** GitHub Releases with full changelogs  

### 13.4 Testing Requirements

✅ **TR-1:** Unit tests for all core modules (>80% coverage)  
✅ **TR-2:** Integration tests for CLI workflows  
✅ **TR-3:** Golden snapshot tests for output formats  
✅ **TR-4:** Cross-platform testing (CI/CD on all target platforms)  
✅ **TR-5:** Offline scenario testing  
✅ **TR-6:** Large repository stress testing (1000+ files)  
✅ **TR-7:** Error handling validation tests  

---

## 14. Testing Strategy

### 14.1 Unit Tests

**Scope:** Core engine modules in isolation

**Coverage Targets:**
- Repository loader: 90%
- Tree resolver: 95%
- Ignore engine: 90%
- File reader: 85%
- Formatters: 90%
- Page assembler: 85%

**Test Cases:**
```go
// Example test structure
package core_test

func TestTreeResolver_DeterministicOrdering(t *testing.T) {
    // Given: Repository with known structure
    // When: Build tree multiple times
    // Then: Order is always identical
}

func TestIgnoreEngine_AppliesDefaultRules(t *testing.T) {
    // Given: Repository with node_modules/
    // When: Apply ignore rules
    // Then: node_modules/ is excluded
}

func TestSourceDetection_LocalPath(t *testing.T) {
    // Given: Valid filesystem path
    // When: Detect source type
    // Then: Returns SourceLocal
}

func TestSourceDetection_GitHubURL(t *testing.T) {
    // Given: Valid GitHub URL
    // When: Detect source type
    // Then: Returns SourceGitHub
}

func TestConnectivity_Online(t *testing.T) {
    // Given: Network available
    // When: Check connectivity
    // Then: Returns true
}

func TestConnectivity_Offline(t *testing.T) {
    // Given: Network unavailable (mocked)
    // When: Check connectivity
    // Then: Returns false
}
```

### 14.2 Integration Tests

**Scope:** End-to-end CLI workflows

**Test Repositories:**
- Small repo (5-10 files)
- Medium repo (50-100 files)
- Large repo (1000+ files)
- Repository with various file types
- Repository with deep nesting
- Repository with special characters in filenames

**Test Scenarios:**
```bash
# Test 1: Basic local conversion
repo2page ./test-repo
assert_exit_code 0
assert_file_exists "test-repo.md"

# Test 2: Format selection
repo2page ./test-repo --format html
assert_file_exists "test-repo.html"

# Test 3: GitHub conversion (online)
repo2page octocat/Hello-World
assert_exit_code 0

# Test 4: GitHub conversion (offline - should fail)
simulate_offline
repo2page octocat/Hello-World
assert_exit_code 5
assert_error_contains "internet connection"

# Test 5: Custom exclusions
repo2page ./test-repo --exclude "*.test.js"
assert_file_not_contains "test-repo.md" "test.js"

# Test 6: Invalid source
repo2page /nonexistent/path
assert_exit_code 2
assert_error_contains "not found"
```

### 14.3 Golden Snapshot Tests

**Purpose:** Ensure output format consistency across versions

**Approach:**
1. Maintain reference "golden" outputs
2. Generate output from test repositories
3. Compare byte-for-byte with golden files
4. Flag any differences for review

**Golden Files:**
```
tests/golden/
├── small-repo.md
├── small-repo.html
├── small-repo.txt
├── medium-repo.md
└── with-readme.md
```

**Test Process:**
```go
func TestGoldenSnapshot_Markdown(t *testing.T) {
    result := Convert(testRepo, MarkdownOptions)
    golden := loadGoldenFile("small-repo.md")
    
    if result != golden {
        t.Errorf("Output differs from golden snapshot")
        // Optionally: save diff for inspection
    }
}
```

### 14.4 Cross-Platform Testing

**CI/CD Matrix:**
```yaml
# GitHub Actions example
strategy:
  matrix:
    os: [ubuntu-latest, macos-latest, windows-latest]
    go: ['1.21']

steps:
  - uses: actions/checkout@v3
  - uses: actions/setup-go@v4
    with:
      go-version: ${{ matrix.go }}
  - run: go test ./...
  - run: go build ./cmd/repo2page
  - run: ./repo2page --version
```

**Platform-Specific Tests:**
- Path handling (forward vs. backslash)
- Line ending normalization (CRLF vs. LF)
- File permission handling
- Case sensitivity (macOS insensitive vs. Linux sensitive)

### 14.5 Performance Benchmarks

**Benchmarks:**
```go
func BenchmarkConvert_SmallRepo(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Convert(smallTestRepo, defaultOptions)
    }
}

func BenchmarkConvert_LargeRepo(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Convert(largeTestRepo, defaultOptions)
    }
}
```

**Performance Targets:**
- Small repo (10 files): <100ms
- Medium repo (100 files): <1s
- Large repo (1000 files): <10s

### 14.6 Error Scenario Testing

**Error Cases:**
```
1. Invalid source path
2. Network timeout (GitHub)
3. Rate limiting (GitHub API)
4. Corrupted repository
5. Permission denied
6. Disk full (output write)
7. Invalid flags combination
8. Malformed .gitignore
9. Binary file overflow
10. Memory limit exceeded
```

Each error case must:
- Return appropriate exit code
- Display clear error message
- Not crash or panic
- Clean up partial state

---

## 15. Performance Constraints

### 15.1 Resource Limits

**File Size:**
- Default max: 500 KB per file
 - Configurable via `--max-file-size` flag (KB)
- Files exceeding limit: skipped with warning

**Repository Size:**
- Target: Handle repos up to 10,000 files
- Practical limit: 50,000 files (memory constraints)

**Memory:**
- Target: <100MB for typical repos
- Maximum: <500MB for large repos
- Stream processing where possible (avoid loading entire repo into memory)

**Time Complexity:**
- File reading: O(n) where n = number of files
- Tree building: O(n log n) for sorting
- Overall: Linear with file count

**Timeout Constraints:**
- CLI: No timeout (blocking operation)
- Web: 30 seconds maximum
- GitHub API: 10 seconds per request (with retry)

### 15.2 Optimization Strategies

**Concurrent Processing:**
```go
// Process files concurrently where safe
func (c *Converter) processFiles(files []File) {
    var wg sync.WaitGroup
    results := make(chan FileResult, len(files))
    
    for _, file := range files {
        wg.Add(1)
        go func(f File) {
            defer wg.Done()
            result := c.processFile(f)
            results <- result
        }(file)
    }
    
    wg.Wait()
    close(results)
}
```

**Lazy Loading:**
- Don't load file contents until needed
- Stream large files if possible
- Skip binary detection for known extensions

**Caching:**
- Cache GitHub API responses (short-term, in-memory only)
- Cache language detection results
- Cache ignore pattern compilation

### 15.3 Performance Monitoring

**Metrics to Track:**
```go
type PerformanceMetrics struct {
    LoadTimeMs      int64
    ProcessTimeMs   int64
    FormatTimeMs    int64
    WriteTimeMs     int64
    TotalTimeMs     int64
    PeakMemoryMB    int
    FilesProcessed  int
    FilesSkipped    int
}
```

**Logging (Debug Mode):**
```bash
repo2page ./project --debug

# Output:
# [DEBUG] Source detection: 2ms
# [DEBUG] Repository load: 45ms
# [DEBUG] Tree building: 12ms
# [DEBUG] File processing: 234ms
# [DEBUG] Format generation: 67ms
# [DEBUG] Write output: 8ms
# [DEBUG] Total: 368ms
```

---

## 16. Security Considerations

### 16.1 Threat Model

**Threats:**
1. **Malicious Repository Content**
   - Executable code in repository
   - Path traversal attempts
   - Symbolic link attacks
   - Extremely large files (DoS)

2. **GitHub API Abuse**
   - Rate limiting
   - Authentication token exposure
   - Private repository access

3. **Output File Risks**
   - XSS in HTML output
   - Script injection
   - Malformed content

4. **System Risks**
   - Arbitrary file write
   - Path traversal in output
   - Resource exhaustion

### 16.2 Security Controls

**Input Validation:**
```go
// Sanitize file paths
func sanitizePath(path string) (string, error) {
    // Remove path traversal attempts
    if strings.Contains(path, "..") {
        return "", errors.New("invalid path: contains ..")
    }
    
    // Resolve to absolute path
    absPath, err := filepath.Abs(path)
    if err != nil {
        return "", err
    }
    
    // Verify within allowed boundaries
    return absPath, nil
}
```

**No Code Execution:**
- Never execute or evaluate repository code
- Treat all files as untrusted data
- No dynamic imports or requires
- No template evaluation with user content

**Output Sanitization:**
```go
// HTML output must escape all user content
func escapeHTML(content string) string {
    return html.EscapeString(content)
}

// Markdown output must escape special characters
func escapeMarkdown(content string) string {
    // Escape backticks, brackets, etc.
    return mdEscape(content)
}
```

**Resource Limits:**
- Enforce max file size (500 KB default)
- Enforce max total repository size
- Timeout for network operations
- Limit concurrent file operations

**File System Safety:**
```go
// Validate output path
func validateOutputPath(path string) error {
    // Ensure not system directories
    forbidden := []string{"/", "/bin", "/usr", "C:\\Windows"}
    
    for _, dir := range forbidden {
        if strings.HasPrefix(path, dir) {
            return errors.New("cannot write to system directory")
        }
    }
    
    return nil
}
```

**GitHub API Security:**
- Never log or expose authentication tokens
- Use HTTPS only
- Validate GitHub URLs before fetching
- Handle rate limiting gracefully
- No persistent token storage

**Web Wrapper Security (Post-MVP):**
- No persistent storage (ephemeral processing)
- Rate limiting per IP
- Input validation on GitHub URLs
- CORS restrictions
- No user authentication (reduces attack surface)

### 16.3 Security Testing

**Test Cases:**
```
1. Attempt path traversal in repository
2. Provide extremely large files
3. Provide binary files with executable content
4. Attempt XSS in filenames
5. Attempt XSS in file contents
6. Test symbolic link handling
7. Test output path traversal
8. Test resource exhaustion (huge repo)
```

### 16.4 Responsible Disclosure

**Security Policy (SECURITY.md):**
```markdown
# Security Policy

## Reporting a Vulnerability

Please report security vulnerabilities to: security@repo2page.dev

Do not open public issues for security vulnerabilities.

We will respond within 48 hours and provide updates every 72 hours
until the issue is resolved.

## Supported Versions

Only the latest stable release receives security updates.
```

---

## 17. Failure Modes & Error Handling

### 17.1 Recoverable Failures (Warnings)

These generate warnings but don't stop processing:

| Scenario | Warning Message | Behavior |
|----------|----------------|----------|
| File too large | `Skipped large file: path/to/file.bin (2.3 MB > 500 KB limit)` | Skip file, continue |
| Binary file | `Skipped binary file: path/to/image.png` | Skip file, continue |
| Permission denied | `Cannot read file: path/to/file.txt (permission denied)` | Skip file, continue |
| Encoding error | `Cannot decode file: path/to/file.txt (invalid UTF-8)` | Skip file, continue |
| Empty file | `Skipped empty file: path/to/.gitkeep` | Skip file, continue |

**Output Example:**
```
Converting repository...
⚠ Skipped 3 files:
  - images/logo.png (binary)
  - data/large.json (2.1 MB, exceeds limit)
  - private/secret.txt (permission denied)

✓ Converted 124 files successfully
```

### 17.2 Hard Failures (Errors)

These stop processing and return non-zero exit code:

| Scenario | Exit Code | Error Message | Suggested Action |
|----------|-----------|---------------|------------------|
| Invalid source | 1 | `Invalid source: not found` | Check path/URL is correct |
| Repo load failure | 2 | `Failed to load repository` | Verify repository exists and is accessible |
| Network error | 5 | `GitHub API unreachable` | Check internet connection |
| GitHub not found | 2 | `Repository not found: user/repo` | Verify repository exists and is public |
| No readable files | 3 | `No files to convert` | Check repository has readable files |
| Output write failure | 4 | `Cannot write output file` | Check disk space and permissions |
| Invalid format | 1 | `Invalid format: xyz` | Use md, html, or txt |
| Offline + GitHub | 5 | `GitHub requires internet connection` | Use local clone or go online |

**Error Output Format:**
```bash
✗ Error: Repository not found: invalid/repo
  → Verify the repository exists and is public
  → Check your internet connection
  → Run with --help for usage information
```

### 17.3 Graceful Degradation

**Partial Success Philosophy:**
- Process as many files as possible
- Report warnings for skipped files
- Generate output even if some files failed
- Provide statistics on partial results

**Example:**
```
✓ Converted 97 of 100 files
⚠ Skipped 3 files (see warnings above)
✓ Output: my-project.md (324 KB)
```

### 17.4 Timeout Handling

**GitHub API Timeouts:**
```go
func (c *GitHubClient) fetchWithTimeout(url string) (*Response, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
    resp, err := c.client.Do(req)
    
    if err != nil {
        if errors.Is(err, context.DeadlineExceeded) {
            return nil, fmt.Errorf("GitHub API timeout (>10s)")
        }
        return nil, err
    }
    
    return resp, nil
}
```

### 17.5 Retry Logic

**GitHub API Retries:**
- Retry on 429 (rate limit) with exponential backoff and jitter (initial backoff: 1s, max backoff: 8s).
- Retry on 5xx (server errors) up to 3 total attempts (initial backoff: 1s, exponential multiplier 2).
- Do not retry on 4xx client errors (except 429). Use response headers (e.g., `Retry-After`) when present.
- Default retry parameters (configurable): `MaxRetries=3`, `InitialBackoff=1s`, `MaxBackoff=8s`, `Jitter=true`.

```go
func (c *GitHubClient) fetchWithRetry(url string, maxRetries int) (*Response, error) {
    for attempt := 0; attempt <= maxRetries; attempt++ {
        resp, err := c.fetch(url)
        
        if err == nil && resp.StatusCode < 500 {
            return resp, nil
        }
        
        if attempt < maxRetries {
            backoff := time.Duration(math.Pow(2, float64(attempt))) * time.Second
            time.Sleep(backoff)
        }
    }
    
    return nil, fmt.Errorf("failed after %d retries", maxRetries)
}
```

---

## 18. Versioning & Releases

### 18.1 Semantic Versioning

**Format:** `MAJOR.MINOR.PATCH`

**Version Increment Rules:**
- **MAJOR:** Breaking changes to CLI interface or output format
- **MINOR:** New features, non-breaking changes
- **PATCH:** Bug fixes, security patches

**Examples:**
- `v0.1.0` - Initial MVP release
- `v0.1.1` - Bug fix (file reading issue)
- `v0.2.0` - New feature (add --summary mode)
- `v1.0.0` - Stable release, API commitment

### 18.2 Pre-Release Versions

**Format:** `MAJOR.MINOR.PATCH-TAG.N`

**Tags:**
- `alpha` - Early development, unstable
- `beta` - Feature complete, testing phase
- `rc` - Release candidate, final testing

**Examples:**
- `v0.1.0-alpha.1`
- `v0.1.0-beta.2`
- `v1.0.0-rc.1`

### 18.3 Release Schedule

**MVP (v0.1.0):**
- Target: Q1 2026
- Core functionality complete
- CLI fully functional
- Basic distribution (Homebrew, direct download)

**v0.2.0:**
- Summary mode
- Enhanced formatters
- Full package manager coverage

**v1.0.0 (Stable):**
- Production-ready
- API stability commitment
- Comprehensive documentation
- Full test coverage

### 18.4 Changelog Format

**CHANGELOG.md:**
```markdown
# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/),
and this project adheres to [Semantic Versioning](https://semver.org/).

## [Unreleased]

### Added
- New features in development

### Changed
- Changes to existing functionality

### Fixed
- Bug fixes

## [0.1.0] - 2026-01-15

### Added
- Initial MVP release
- Convert local and GitHub repositories
- Support for Markdown, HTML, and plain text output
- CLI with comprehensive flags
- Deterministic output generation
- Automatic source detection
- Online/offline operation support

### Security
- Input validation for all file paths
- Output sanitization for HTML format
```

### 18.5 Release Process

**Step-by-Step:**
```bash
# 1. Update version in code
vim internal/version/version.go  # Update Version constant

# 2. Update CHANGELOG.md
vim CHANGELOG.md  # Move [Unreleased] to new version

# 3. Commit changes
git add .
git commit -m "chore: bump version to v0.1.0"

# 4. Create tag
git tag -a v0.1.0 -m "Release v0.1.0"

# 5. Push tag (triggers automated release)
git push origin v0.1.0

# 6. GoReleaser handles the rest:
#    - Build binaries for all platforms
#    - Create GitHub Release
#    - Upload artifacts
#    - Update package manager formulas
#    - Generate changelog
```

### 18.6 Version Embedding

**Build-time Variables:**
```go
// internal/version/version.go
package version

var (
    // Set via ldflags during build
    Version = "dev"
    Commit  = "unknown"
    Date    = "unknown"
)

func Full() string {
    return fmt.Sprintf("repo2page %s (commit %s, built %s)", 
        Version, Commit, Date)
}
```

**CLI Output:**
```bash
$ repo2page --version
repo2page v0.1.0 (commit abc123, built 2026-01-15T10:30:00Z)
```

---

## 19. Documentation Requirements

### 19.1 README.md

**Required Sections:**
1. **Project Overview** - What is repo2page?
2. **Key Features** - Bullet points of core capabilities
3. **Installation** - All installation methods
4. **Quick Start** - Basic usage examples
5. **Usage** - Comprehensive CLI documentation
6. **Examples** - Real-world use cases
7. **Output Formats** - Format descriptions and samples
8. **Configuration** - Advanced options
9. **Contributing** - How to contribute
10. **License** - License information

**Template:**
```markdown
# repo2page

Convert entire repositories into single, portable documents.

## Features

- ✅ Convert local and GitHub repositories
- ✅ Zero dependencies (standalone binary)
- ✅ Works offline for local repos
- ✅ Multiple output formats (Markdown, HTML, plain text)
- ✅ Deterministic, reproducible output
- ✅ Cross-platform (Windows, macOS, Linux)

## Installation

### macOS
```bash
brew install repo2page
```

### Windows
```bash
choco install repo2page
```

### Linux
```bash
curl -sSL https://install.repo2page.dev/install.sh | sh
```

## Quick Start

```bash
# Convert local repository
repo2page ./my-project

# Convert GitHub repository
repo2page github.com/user/repo

# Generate HTML
repo2page ./my-project --format html
```

## Usage

[Full documentation...]

## Examples

[Real-world examples...]

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md)

## License

MIT License - see [LICENSE](LICENSE)
```

### 19.2 CLI Help Text

**`repo2page --help`:**
```
repo2page - Convert repositories into single, portable documents

USAGE:
    repo2page <source> [options]

SOURCE:
    Local path:       ./my-project, /path/to/repo
    GitHub URL:       https://github.com/user/repo
    GitHub shorthand: user/repo

OPTIONS:
    --format <fmt>     Output format (md, html, txt) [default: md]
    --output <file>    Output filename [default: auto-generated]
    --branch <name>    Git branch [default: main]
    --commit <hash>    Specific commit hash
    --exclude <pat>    Ignore pattern (can be repeated)
    --summary          Summary mode (structure only)
    --no-tree          Disable directory tree
    --help             Show this help message
    --version          Show version information

EXAMPLES:
    repo2page ./my-project
    repo2page github.com/torvalds/linux --format html
    repo2page user/repo --branch develop --output snapshot.md

ONLINE/OFFLINE:
    Local repositories work offline.
    GitHub repositories require internet connection.

MORE INFO:
    Documentation: https://repo2page.dev/docs
    Report issues: https://github.com/yourusername/repo2page/issues
```

### 19.3 Contributing Guide (CONTRIBUTING.md)

**Required Sections:**
1. **How to Contribute**
2. **Development Setup**
3. **Building from Source**
4. **Running Tests**
5. **Code Style**
6. **Commit Messages**
7. **Pull Request Process**
8. **Reporting Bugs**
9. **Suggesting Features**

### 19.4 License

**Recommended:** MIT License

```
MIT License

Copyright (c) 2026 [Your Name/Organization]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

[Standard MIT License text...]
```

### 19.5 Online Documentation

**Documentation Site Structure:**
```
docs/
├── index.md (Getting Started)
├── installation.md (All installation methods)
├── usage.md (CLI reference)
├── formats.md (Output format specs)
├── configuration.md (Advanced options)
├── examples.md (Real-world examples)
├── faq.md (Frequently Asked Questions)
├── troubleshooting.md (Common issues)
├── api.md (Library API for developers)
└── contributing.md (Contribution guide)
```

**Hosting Options:**
- GitHub Pages
- Read the Docs
- Custom domain (repo2page.dev)

---

## 20. Open Source Strategy

### 20.1 License Selection

**Recommendation: MIT License**

**Rationale:**
- Permissive, business-friendly
- Allows commercial use
- Maximum adoption potential
- Standard for CLI tools

**Alternative Considered:** Apache 2.0
- More explicit patent protection
- More verbose
- Less common for CLI tools

**Decision:** MIT for MVP, can dual-license later if needed

### 20.2 Repository Structure

**GitHub Repository:**
```
repo2page/
├── .github/
│   ├── workflows/
│   │   ├── ci.yml (tests on push/PR)
│   │   ├── release.yml (automated releases)
│   │   └── codeql.yml (security scanning)
│   ├── ISSUE_TEMPLATE/
│   │   ├── bug_report.md
│   │   └── feature_request.md
│   └── PULL_REQUEST_TEMPLATE.md
├── cmd/
│   └── repo2page/
│       └── main.go
├── internal/
│   ├── core/
│   ├── github/
│   └── formatters/
├── pkg/
│   └── repo2page/ (public API)
├── tests/
│   ├── fixtures/
│   ├── golden/
│   └── integration/
├── docs/
├── .goreleaser.yml
├── go.mod
├── go.sum
├── Makefile
├── README.md
├── CHANGELOG.md
├── CONTRIBUTING.md
├── LICENSE
├── SECURITY.md
└── .gitignore
```

### 20.3 Community Guidelines

**Code of Conduct:**
- Adopt Contributor Covenant
- Enforce respectful, inclusive community
- Clear reporting mechanism

**Issue Labels:**
```
Type:
- bug
- feature
- documentation
- question

Priority:
- critical
- high
- medium
- low

Status:
- needs-triage
- in-progress
- blocked
- wontfix

Good First Issue:
- good-first-issue
- help-wanted
```

### 20.4 Contribution Workflow

**Pull Request Process:**
1. Fork repository
2. Create feature branch (`feat/new-feature`)
3. Make changes with tests
4. Run `make test` and `make lint`
5. Commit with conventional commits
6. Push and open PR
7. Address review feedback
8. Maintainer merges when approved

**Conventional Commits:**
```
feat: add summary mode flag
fix: resolve path traversal issue
docs: update installation guide
chore: bump dependencies
test: add integration tests for GitHub loader
refactor: simplify tree resolver
```

### 20.5 Maintenance Model

**Maintainer Responsibilities:**
- Review PRs within 72 hours
- Release patches for critical bugs within 48 hours
- Monthly minor releases (feature additions)
- Quarterly major releases (if needed)

**Community Support:**
- GitHub Discussions for questions
- Issues for bugs and features
- Discord/Slack (optional, post-v1.0)

---

## 21. Roadmap

### 21.1 MVP (v0.1.0) - Q1 2026

**Core Functionality:**
- ✅ Local repository conversion
- ✅ GitHub repository conversion
- ✅ Auto-detect source type and connectivity
- ✅ Markdown, HTML, plain text output
- ✅ CLI with all specified flags
- ✅ Deterministic output
- ✅ Default ignore rules
- ✅ Error handling and exit codes

**Distribution:**
- ✅ Standalone binaries (Windows, macOS, Linux)
- ✅ Homebrew formula
- ✅ Direct download from GitHub Releases
- ✅ Basic installation scripts

**Documentation:**
- ✅ README with installation and usage
- ✅ CLI help text
- ✅ LICENSE file

**Testing:**
- ✅ Unit tests (>80% coverage)
- ✅ Integration tests
- ✅ Cross-platform CI/CD

### 21.2 v0.2.0 - Q2 2026

**Features:**
- Summary mode (structure and metadata only, no full file contents)
- Enhanced HTML output with syntax highlighting
- Progress indicator for large repositories
- Configurable file size limits via CLI flag
- Support for custom output templates

**Distribution:**
- Chocolatey package (Windows)
- Scoop manifest (Windows)
- Snap package (Linux)
- Winget package (Windows)

**Improvements:**
- Performance optimization for large repos
- Better error messages with suggestions
- Memory usage reduction

**Documentation:**
- Online documentation site
- Video tutorials
- Real-world examples library

### 21.3 v0.3.0 - Q3 2026

**Features:**
- `.repo2page.yml` configuration file support
- Custom ignore rules per project
- Multiple output formats in one run
- Directory filtering (include specific paths only)
- Archive mode (include .git directory metadata)

**Advanced Options:**
- `--include-only <pattern>` - Process only matching files
- `--max-depth <n>` - Limit directory traversal depth
- `--split-by-dir` - Generate separate files per directory
- `--stats-only` - Output statistics JSON

**Quality of Life:**
- Auto-update notifications
- Shell completion (bash, zsh, fish)
- Color output for terminal
- Interactive mode (prompt for options)

### 21.4 v0.4.0 - Q4 2026

**Features:**
- Private GitHub repository support (token auth)
- GitLab support
- Bitbucket support
- Git history snapshot (specific date/tag)

**Web Wrapper:**
- Simple web interface launch
- Drag-and-drop local folder support
- Real-time preview before download
- API endpoint for programmatic access

**Integrations:**
- GitHub Action for automated documentation
- Pre-commit hook integration
- VS Code extension (optional)

### 21.5 v1.0.0 - Q1 2027 (Stable Release)

**Stability:**
- API freeze (breaking changes require v2.0)
- Production-ready for enterprise use
- Comprehensive documentation
- 95%+ test coverage
- Security audit completed

**Enterprise Features:**
- Batch processing (multiple repos)
- Custom branding in output
- Compliance mode (include licenses, copyrights)
- Diff mode (compare two versions)

**Performance:**
- Handle repos with 100,000+ files
- Streaming mode for extremely large repos
- Parallel processing optimization

**Ecosystem:**
- Plugin system for custom formatters
- Library packages for other languages (Python, JavaScript)
- CI/CD templates for major platforms

### 21.6 Post-1.0 (Future)

**Advanced Features (2.x):**
- Semantic code analysis
- Dependency graph visualization
- Code metrics and statistics
- Documentation generation from comments
- AI-powered summaries

**Platform Expansion:**
- Mobile apps (iOS/Android) for viewing
- Browser extension for quick GitHub exports
- Cloud service for hosted conversion
- Enterprise on-premises deployment

**Community:**
- Official forum/community site
- Plugin marketplace
- Professional support offerings
- Training and certification programs

---

## 22. Success Metrics

### 22.1 Technical Metrics

**Performance:**
- Conversion speed: <1s for typical repo (100 files)
- Binary size: <15MB
- Memory usage: <200MB for typical repo
- Test coverage: >85%
- Zero critical bugs in stable release

**Quality:**
- CLI exit code accuracy: 100%
- Deterministic output: 100% (same input = same output)
- Cross-platform compatibility: 100%
- Documentation coverage: All features documented

### 22.2 User Adoption Metrics

**Downloads (6 months post-launch):**
- GitHub releases: 10,000+ downloads
- Homebrew installs: 5,000+
- Chocolatey installs: 2,000+
- Total unique users: 15,000+

**Engagement:**
- GitHub stars: 1,000+
- GitHub issues: <50 open, >200 closed
- Average issue resolution time: <7 days
- Documentation page views: 50,000+

**Community:**
- Contributors: 20+
- Pull requests: 50+
- Forum posts/discussions: 500+

### 22.3 Business Metrics (If Applicable)

**Ecosystem Growth:**
- Integrations built by community: 10+
- Forks: 100+
- Dependent projects: 50+
- Mentions in blog posts/articles: 50+

**Enterprise Adoption:**
- Companies using in production: 20+
- Paid support contracts (if offered): 5+

---

## 23. Risk Management

### 23.1 Technical Risks

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| GitHub API rate limits | High | Medium | Implement caching, respect rate limits, provide local fallback |
| Large repo performance | Medium | High | Stream processing, configurable limits, progress indicators |
| Cross-platform bugs | Medium | Medium | Comprehensive CI/CD, early beta testing on all platforms |
| Binary size bloat | Low | Low | Regular profiling, minimize dependencies, compression |
| Security vulnerability | Low | High | Regular security audits, dependency scanning, responsible disclosure policy |

### 23.2 Product Risks

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Low user adoption | Medium | High | Strong marketing, solve real pain points, excellent documentation |
| Competing tools emerge | Medium | Medium | Focus on unique value (offline, universal access), build community |
| Feature creep | Medium | Medium | Strict PRD adherence, clear roadmap, say no to scope additions |
| Maintenance burden | Low | Medium | Clean architecture, comprehensive tests, community contributions |

### 23.3 Operational Risks

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Distribution pipeline breaks | Low | High | Automated testing of release process, backup distribution methods |
| GitHub API changes | Low | Medium | Version pinning, adapter pattern, monitoring for deprecations |
| Package manager issues | Medium | Low | Multiple distribution channels, direct download always available |
| Documentation outdated | Medium | Low | Automated doc generation where possible, regular reviews |

### 23.4 Contingency Plans

**If GitHub API becomes unreliable:**
- Prioritize local repository support
- Add support for other Git hosting services
- Implement local Git parsing (no API required)

**If binary size exceeds 20MB:**
- Investigate compression options
- Consider separate binaries for web vs. CLI
- Evaluate alternative languages/compilers

**If performance is insufficient:**
- Implement streaming for large files
- Add parallel processing
- Provide "quick mode" with reduced features

**If adoption is slow:**
- Increase marketing efforts
- Create video tutorials and demos
- Reach out to influencers and communities
- Add killer features that competitors lack

---

## 24. Product Philosophy

### 24.1 Core Principles

**Simplicity over Cleverness**
- Do one thing exceptionally well
- Avoid feature bloat
- Prioritize common use cases
- Default behavior should "just work"

**Determinism over Flexibility**
- Same input always produces same output
- Predictable behavior across platforms
- Clear, documented rules
- No magic or hidden behavior

**Portability over Platform Dependence**
- Works everywhere: Windows, macOS, Linux
- Zero runtime dependencies
- No platform-specific features unless critical
- Offline capability is first-class

**Trust over Novelty**
- Reliable, predictable tool
- Conservative approach to changes
- Respect user expectations
- Stability over bleeding-edge features

**User Respect**
- Fast startup, fast execution
- Clear error messages
- No telemetry or tracking
- Privacy-first design
- Open source, community-driven

### 24.2 Design Philosophy

**CLI First**
- CLI is the primary interface
- Web wrapper is secondary
- Power users are primary audience
- Simple for beginners, powerful for experts

**Convention over Configuration**
- Smart defaults work for 90% of use cases
- Minimal required configuration
- Flags for edge cases
- Configuration file for advanced users (post-MVP)

**Progressive Enhancement**
- Basic functionality works everywhere
- Advanced features when available
- Graceful degradation
- No all-or-nothing dependencies

**Developer Experience**
- Fast compilation
- Easy to contribute
- Clear architecture
- Well-tested code
- Comprehensive documentation

### 24.3 The "One Repo, One Page" Philosophy

This tool exists because:

1. **Repositories are fragmented** - Code is split across hundreds of files
2. **Context is lost** - Structure gets lost in file browsers
3. **Sharing is hard** - Can't attach a repo to an email
4. **Access is limited** - GitHub requires accounts and internet
5. **AI needs context** - LLMs work better with unified documents

**repo2page solves this by:**
- Converting any repository into a single, readable document
- Preserving structure and relationships
- Making code portable and shareable
- Working everywhere, online and offline
- Requiring zero dependencies

**The promise:**
> "Any repository, any format, anywhere, anytime."

---

## 25. Implementation Guidelines for Coding Agents

### 25.1 Agent Instructions

**You are implementing repo2page as specified in this PRD.**

**Hard Constraints:**
- Language: Go (version 1.21+)
- Architecture: Shared core + CLI wrapper
- Must follow deterministic ordering rules
- Must implement all error handling as specified
- Must support online/offline operation
- Must auto-detect source type and connectivity

**Implementation Order:**
1. Core type definitions (`internal/core/types.go`)
2. Source detection (`internal/core/detect.go`)
3. Local repository loader (`internal/loader/local.go`)
4. GitHub repository loader (`internal/loader/github.go`)
5. Tree resolver (`internal/core/tree.go`)
6. Ignore engine (`internal/ignore/engine.go`)
7. File reader and normalizer (`internal/core/reader.go`)
8. Formatters (Markdown, HTML, Plain Text)
9. Page assembler (`internal/core/assembler.go`)
10. CLI wrapper (`cmd/repo2page/main.go`)
11. Tests (unit, integration, golden snapshots)

**Do NOT add features outside this document.**

**If you encounter ambiguity:**
1. Refer to the PRD section on that topic
2. Choose the simplest, most predictable behavior
3. Follow Go idioms and conventions
4. Ask for clarification if critically ambiguous

### 25.2 Code Style Guidelines

**Go Standards:**
- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use `gofmt` for formatting
- Use `golint` and `go vet` for linting
- Idiomatic error handling (no exceptions)

**Project Conventions:**
```go
// Package names: lowercase, single word
package core

// Exported types: PascalCase
type RepoInput struct {}

// Unexported types: camelCase
type treeNode struct {}

// Constants: PascalCase or ALL_CAPS for clarity
const MaxFileSizeKB = 500

// Functions: PascalCase (exported), camelCase (unexported)
func Convert() {}
func buildTree() {}

// Error handling: explicit, informative
if err != nil {
    return fmt.Errorf("failed to load repository: %w", err)
}

// Comments: Complete sentences, explain why not what
// detectSource determines if input is a local path or GitHub URL.
// It prioritizes GitHub URLs to avoid false positives with local paths.
```

**Testing Conventions:**
```go
// Test function naming
func TestConvert_LocalRepository_Success(t *testing.T) {}

// Table-driven tests for multiple cases
func TestDetectSource(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        want     SourceType
        wantErr  bool
    }{
        {"local path", "./project", SourceLocal, false},
        {"GitHub URL", "https://github.com/user/repo", SourceGitHub, false},
        {"GitHub shorthand", "user/repo", SourceGitHub, false},
        {"invalid", "", SourceLocal, true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := DetectSource(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("want error %v, got %v", tt.wantErr, err)
            }
            if got != tt.want {
                t.Errorf("want %v, got %v", tt.want, got)
            }
        })
    }
}
```

### 25.3 Project Structure Template

```
repo2page/
├── cmd/
│   └── repo2page/
│       └── main.go                 # CLI entry point
│
├── internal/
│   ├── core/
│   │   ├── types.go               # Core type definitions
│   │   ├── detect.go              # Source detection
│   │   ├── converter.go           # Main conversion logic
│   │   ├── tree.go                # Tree resolver
│   │   ├── reader.go              # File reader
│   │   └── assembler.go           # Page assembler
│   │
│   ├── loader/
│   │   ├── local.go               # Local filesystem loader
│   │   └── github.go              # GitHub API loader
│   │
│   ├── ignore/
│   │   ├── engine.go              # Ignore pattern engine
│   │   └── patterns.go            # Default patterns
│   │
│   ├── formatter/
│   │   ├── markdown.go            # Markdown formatter
│   │   ├── html.go                # HTML formatter
│   │   └── text.go                # Plain text formatter
│   │
│   └── version/
│       └── version.go             # Version info
│
├── pkg/
│   └── repo2page/
│       └── api.go                 # Public API (for library use)
│
├── tests/
│   ├── fixtures/
│   │   ├── small-repo/            # Test repository
│   │   └── with-readme/           # Test repository
│   │
│   ├── golden/
│   │   ├── small-repo.md          # Expected output
│   │   └── small-repo.html        # Expected output
│   │
│   └── integration/
│       └── cli_test.go            # CLI integration tests
│
├── .github/
│   └── workflows/
│       ├── ci.yml                 # CI pipeline
│       └── release.yml            # Release automation
│
├── .goreleaser.yml                # Release configuration
├── go.mod                         # Go module definition
├── go.sum                         # Dependency checksums
├── Makefile                       # Build automation
├── README.md                      # User documentation
├── CHANGELOG.md                   # Version history
├── CONTRIBUTING.md                # Contribution guide
├── LICENSE                        # MIT License
├── SECURITY.md                    # Security policy
└── .gitignore                     # Git ignore rules
```

### 25.4 Makefile Template

```makefile
.PHONY: build test lint clean install

# Build configuration
BINARY_NAME=repo2page
VERSION=$(shell git describe --tags --always --dirty)
COMMIT=$(shell git rev-parse --short HEAD)
DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS=-ldflags "-s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)"

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	go build $(LDFLAGS) -o bin/$(BINARY_NAME) ./cmd/repo2page

# Build for all platforms
build-all:
	@echo "Building for all platforms..."
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/$(BINARY_NAME)-linux-amd64 ./cmd/repo2page
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o bin/$(BINARY_NAME)-darwin-amd64 ./cmd/repo2page
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o bin/$(BINARY_NAME)-darwin-arm64 ./cmd/repo2page
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o bin/$(BINARY_NAME)-windows-amd64.exe ./cmd/repo2page

# Run tests
test:
	@echo "Running tests..."
	go test -v -race -cover ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Run linters
lint:
	@echo "Running linters..."
	golangci-lint run ./...

# Format code
fmt:
	@echo "Formatting code..."
	gofmt -s -w .
	go mod tidy

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -f coverage.out coverage.html

# Install locally
install: build
	@echo "Installing $(BINARY_NAME)..."
	cp bin/$(BINARY_NAME) /usr/local/bin/

# Run integration tests
test-integration:
	@echo "Running integration tests..."
	go test -v -tags=integration ./tests/integration/...

# Update golden snapshots
test-update-golden:
	@echo "Updating golden snapshots..."
	go test -v ./... -update-golden

# Development: run with local changes
dev:
	go run ./cmd/repo2page $(ARGS)

# Show help
help:
	@echo "Available targets:"
	@echo "  build              - Build the binary"
	@echo "  build-all          - Build for all platforms"
	@echo "  test               - Run tests"
	@echo "  test-coverage      - Run tests with coverage report"
	@echo "  test-integration   - Run integration tests"
	@echo "  test-update-golden - Update golden snapshot tests"
	@echo "  lint               - Run linters"
	@echo "  fmt                - Format code"
	@echo "  clean              - Clean build artifacts"
	@echo "  install            - Install locally"
	@echo "  dev                - Run with local changes (use ARGS=...)"
	@echo "  help               - Show this help"
```

### 25.5 CI/CD Pipeline Template

**`.github/workflows/ci.yml`:**
```yaml
name: CI

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    name: Test
    runs-on: ${{ matrix.os }}
    strategy:
      matrix:
        os: [ubuntu-latest, macos-latest, windows-latest]
        go: ['1.21']
    
    steps:
      - name: Checkout code
        uses: actions/checkout@v4
      
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: ${{ matrix.go }}
      
      - name: Cache Go modules
        uses: actions/cache@v4
        with:
          path: ~/go/pkg/mod
          key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
      
      - name: Run tests
        run: make test
      
      - name: Run integration tests
        run: make test-integration
      
      - name: Upload coverage
        if: matrix.os == 'ubuntu-latest'
        uses: codecov/codecov-action@v4
        with:
          file: ./coverage.out

  lint:
    name: Lint
    runs-on: ubuntu-latest
    
    steps:
      - name: Checkout code
        uses: actions/checkout@v4
      
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      
      - name: Run golangci-lint
        uses: golangci/golangci-lint-action@v4
        with:
          version: latest

  build:
    name: Build
    runs-on: ubuntu-latest
    
    steps:
      - name: Checkout code
        uses: actions/checkout@v4
      
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      
      - name: Build
        run: make build
      
      - name: Test CLI
        run: |
          ./bin/repo2page --version
          ./bin/repo2page --help
```

**`.github/workflows/release.yml`:**
```yaml
name: Release

on:
  push:
    tags:
      - 'v*'

permissions:
  contents: write

jobs:
  goreleaser:
    runs-on: ubuntu-latest
    
    steps:
      - name: Checkout code
        uses: actions/checkout@v4
        with:
          fetch-depth: 0
      
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      
      - name: Run GoReleaser
        uses: goreleaser/goreleaser-action@v5
        with:
          version: latest
          args: release --clean
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

---

## 26. Quality Assurance Checklist

### 26.1 Pre-Release Checklist

**Functionality:**
- [ ] Local repository conversion works
- [ ] GitHub repository conversion works
- [ ] Auto-detection identifies source correctly
- [ ] Offline mode works for local repos
- [ ] Online check works for GitHub repos
- [ ] All output formats generate correctly
- [ ] All CLI flags function as specified
- [ ] Exit codes are correct for all scenarios
- [ ] Ignore rules apply correctly
- [ ] README appears first when present
- [ ] Large files are skipped with warnings
- [ ] Binary files are skipped with warnings

**Cross-Platform:**
- [ ] Builds successfully on Linux
- [ ] Builds successfully on macOS (Intel)
- [ ] Builds successfully on macOS (ARM)
- [ ] Builds successfully on Windows
- [ ] CLI works on Linux
- [ ] CLI works on macOS
- [ ] CLI works on Windows
- [ ] Path handling works on all platforms
- [ ] Line endings normalize correctly

**Performance:**
- [ ] Small repos (<10 files) convert in <100ms
- [ ] Medium repos (<100 files) convert in <1s
- [ ] Large repos (<1000 files) convert in <10s
- [ ] Binary size is under 20MB
- [ ] Memory usage is under 200MB for typical repos

**Testing:**
- [ ] All unit tests pass
- [ ] All integration tests pass
- [ ] Golden snapshot tests pass
- [ ] Test coverage is >85%
- [ ] No flaky tests

**Documentation:**
- [ ] README is complete and accurate
- [ ] CLI help text is correct
- [ ] CHANGELOG is updated
- [ ] Examples are tested and work
- [ ] Installation instructions are verified

**Distribution:**
- [ ] Binaries build for all platforms
- [ ] Homebrew formula works
- [ ] Install scripts work
- [ ] GitHub Release created successfully
- [ ] All artifacts uploaded

**Security:**
- [ ] No hardcoded credentials
- [ ] Input validation works
- [ ] Path traversal prevented
- [ ] Output sanitization works
- [ ] No code execution vulnerabilities

### 26.2 Post-Release Checklist

**Verification:**
- [ ] Download binaries from GitHub Releases
- [ ] Test installation on fresh machines (all platforms)
- [ ] Verify Homebrew formula works
- [ ] Verify install scripts work
- [ ] Test with real-world repositories
- [ ] Check for any critical bugs

**Communication:**
- [ ] Announcement on GitHub Releases
- [ ] Update README with latest version
- [ ] Post on relevant forums/communities
- [ ] Update documentation site (if applicable)
- [ ] Social media announcement (if applicable)

**Monitoring:**
- [ ] Watch for new issues
- [ ] Monitor download statistics
- [ ] Check for negative feedback
- [ ] Monitor performance metrics
- [ ] Track error reports (if telemetry added)

---

## 27. Glossary

**Terms and Definitions:**

- **Artifact:** The generated output file (Markdown, HTML, or plain text)
- **Binary:** Non-text file (images, executables, compiled files)
- **CLI:** Command-Line Interface
- **Deterministic:** Same input always produces same output
- **Ephemeral:** Temporary, not stored persistently
- **Golden Snapshot:** Reference output used for testing
- **Hard Failure:** Error that stops processing
- **Ignore Pattern:** Rule specifying files/directories to exclude
- **MVP:** Minimum Viable Product
- **Recoverable Failure:** Issue that generates warning but continues processing
- **Repository/Repo:** Code repository (local directory or GitHub project)
- **Source Detection:** Process of determining if input is local or GitHub
- **Standalone Binary:** Executable with no external dependencies
- **Tree:** Directory structure representation
- **Wrapper:** Interface layer (CLI or Web) around core engine

---

## 28. Appendices

### Appendix A: GitHub API Reference

**Endpoints Used:**
```
GET /repos/{owner}/{repo}
GET /repos/{owner}/{repo}/contents/{path}
GET /repos/{owner}/{repo}/git/trees/{tree_sha}?recursive=1
```

**Authentication:**
- Public repos: No auth required
- Private repos (post-MVP): Personal access token

**Rate Limits:**
- Unauthenticated: 60 requests/hour
- Authenticated: 5000 requests/hour

**Error Codes:**
- 200: Success
- 404: Not found
- 403: Rate limit or permission denied
- 5xx: Server errors (retry)

### Appendix B: File Extension to Language Mapping

```go
var languageMap = map[string]string{
    ".go":     "go",
    ".js":     "javascript",
    ".ts":     "typescript",
    ".py":     "python",
    ".rb":     "ruby",
    ".java":   "java",
    ".c":      "c",
    ".cpp":    "cpp",
    ".rs":     "rust",
    ".sh":     "bash",
    ".yaml":   "yaml",
    ".yml":    "yaml",
    ".json":   "json",
    ".xml":    "xml",
    ".html":   "html",
    ".css":    "css",
    ".md":     "markdown",
    ".sql":    "sql",
    ".php":    "php",
    ".swift":  "swift",
    ".kt":     "kotlin",
    ".r":      "r",
    ".m":      "objective-c",
    ".scala":  "scala",
    ".dart":   "dart",
    ".lua":    "lua",
    ".vim":    "vim",
    ".pl":     "perl",
}
```

### Appendix C: Binary File Detection

**Extensions (always binary):**
```
.exe, .dll, .so, .dylib, .o, .a, .lib
.jpg, .jpeg, .png, .gif, .bmp, .ico, .svg
.pdf, .doc, .docx, .xls, .xlsx, .ppt, .pptx
.zip, .tar, .gz, .bz2, .7z, .rar
.mp3, .mp4, .avi, .mov, .wmv, .flv
.ttf, .otf, .woff, .woff2
.class, .jar, .war, .ear
.dmg, .iso, .img
```

**Magic Number Detection (if extension unknown):**
- Check first 512 bytes for null bytes
- If contains null bytes → likely binary
- Otherwise → treat as text

### Appendix D: Output Size Estimates

**Typical Repository Sizes:**

| Repository Size | File Count | Output Size (MD) | Conversion Time |
|-----------------|------------|------------------|-----------------|
| Small | 10-50 | 50-500 KB | <0.5s |
| Medium | 50-500 | 500 KB - 5 MB | 0.5-3s |
| Large | 500-2000 | 5-20 MB | 3-15s |
| Very Large | 2000-10000 | 20-100 MB | 15-60s |

**Note:** HTML output is typically 20-30% larger than Markdown due to inline CSS and formatting.

---

## 29. Conclusion

This PRD serves as the complete specification for repo2page from concept through implementation, distribution, and beyond. It is designed to:

1. **Guide implementation** with clear, actionable requirements
2. **Ensure consistency** across all features and platforms
3. **Enable autonomous development** by AI coding agents or human developers
4. **Provide a reference** for all product decisions
5. **Document philosophy** and long-term vision

**Key Takeaways:**

- **Mission:** Convert any repository into a single, portable document
- **Technology:** Go-based standalone binaries with zero dependencies
- **Distribution:** Universal accessibility via package managers and direct downloads
- **Philosophy:** Simplicity, determinism, portability, and user respect
- **Timeline:** MVP in Q1 2026, stable v1.0 by Q1 2027

**Success Definition:**

repo2page succeeds when any developer, anywhere, can instantly convert any repository into a shareable document—online or offline, on any platform, with a single command.

---

**Document Status:** Final  
**Version:** 2.0  
**Last Updated:** January 2026  
**Maintained By:** Project Lead  
**Next Review:** Upon MVP completion

---

*End of Product Requirements Document*