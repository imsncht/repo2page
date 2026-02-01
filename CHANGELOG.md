# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.6] - 2026-02-01

### Added
- **Optimized GitHub Fetching:** Replaced sequential file-by-file downloading with atomic tarball download and local extraction. This drastically reduces API calls (1 call per repo vs N calls) and improves speed by 10-100x.
- **Progress Visualization:** Added a real-time progress bar for file extractions and repository downloads.
- **Winget Distribution:** Fixed automated publishing to the Winget Community Repository to include both x64 and ARM64 architectures.
- **Local Repository Support:** Full file system traversal with `.gitignore` support.
- **Auto-Detection:** Smart detection of source type (Local path vs GitHub URL) and connectivity status.
- **Formatters:** Markdown, HTML, and Plain Text output formats.
- **Tree Visualization:** ASCII tree generation for repository structure.
- **Safety:** Binary file detection and file size limits (default 500KB).

### Fixed
- Fixed an issue where PAX global headers in GitHub tarballs could cause extraction errors.
- Fixed Winget publishing workflow to satisfy architecture consistency requirements.

## [0.1.0-alpha] - 2026-01-13
- Initial MVP release.
