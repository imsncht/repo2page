# Installation Guide

`repo2page` is designed to be easy to install across all major platforms.

## 1. Windows (Recommended)

### Scoop
The easiest way to install on Windows:
```powershell
scoop bucket add imsncht https://github.com/imsncht/scoop-bucket.git
scoop install repo2page
```

### Winget
```powershell
winget install repo2page
```

---

## 2. macOS & Linux

### Homebrew
```bash
brew tap imsncht/homebrew-tap
brew install repo2page
```

### One-Line Script
```bash
curl -sfL https://raw.githubusercontent.com/imsncht/repo2page/main/scripts/install.sh | sh
```

---

## 3. From Source
If you have Go 1.23 or higher installed:
```bash
go install github.com/imsncht/repo2page/cmd/repo2page@latest
```

---

## 4. Manual Download
You can also download the latest binaries directly from our [Releases Page](https://github.com/imsncht/repo2page/releases).
