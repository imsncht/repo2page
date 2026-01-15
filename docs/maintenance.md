# Maintenance Guide (A-Z)

This documentation is intended for project maintainers. it covers everything required to keep `repo2page` up-to-date and functional across all distribution platforms.

---

## A. Versioning and Releases
We follow [Semantic Versioning (SemVer)](https://semver.org/). 
- **Major**: Breaking changes in CLI or output structure.
- **Minor**: New features (e.g., new output formats).
- **Patch**: Bug fixes and internal optimizations.

### The Release Checklist
1. Update `CHANGELOG.md` with the new version and changes.
2. Ensure internal version strings in `internal/version/version.go` are correct (if not automated).
3. Push all changes to `main`.
4. Create and push a new tag:
   ```bash
   git tag v1.x.x
   git push origin v1.x.x
   ```

---

## B. Automation (GitHub Actions)
The repository uses several workflows to automate maintenance:

- **Release (`release.yml`)**: Triggered by version tags (`v*`). It runs `goreleaser` to:
  - Build binaries for Windows, macOS, and Linux.
  - Create a GitHub Release with changelogs.
  - Update the [imsncht/homebrew-tap](https://github.com/imsncht/homebrew-tap).
  - Update the [imsncht/scoop-bucket](https://github.com/imsncht/scoop-bucket).
- **Winget (`winget.yml`)**: Triggered by new GitHub Releases. It automatically submits a Pull Request to the [Microsoft Winget Community Repository](https://github.com/microsoft/winget-pkgs).

---

## C. Package Managers

### 1. Scoop
The manifest is located at `repo2page.json` in your bucket repo.
- **Troubleshooting**: If a user reports a checksum error, verify the SHA256 in the manifest against the release asset.
- **Indexing**: To ensure visibility on `scoop.sh`, keep the `scoop-bucket` topic on the bucket repository.

### 2. Winget
Winget PRs usually take 12-48 hours for automated/moderator approval.
- **Manual Check**: Monitor PRs at [microsoft/winget-pkgs](https://github.com/microsoft/winget-pkgs/pulls?q=is%3Apr+repo2page).
- **Common Issues**: "Validation-Executable-Error" often occurs if the app returns a non-zero exit code when run without arguments during validation. Maintainers usually manually clear this for CLI tools.

### 3. Homebrew
The formula is maintained in your tap. GoReleaser handles the updates.

---

## D. Dependency Management
Run periodic updates for Go dependencies:
```bash
go get -u ./...
go mod tidy
```

---

## E. Documentation Upkeep
Whenever the CLI adds new flags or formats:
1. Update `docs/usage.md`.
2. Update the README's flag table.
3. Update `docs/formats.md` (if applicable).

---

## Z. Final Review before Release
- [ ] Tests pass locally (`go test ./...`)
- [ ] Linting is clean
- [ ] Documentation reflects all new changes
- [ ] No hardcoded paths in the code
