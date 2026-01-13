# Release Gates & Checklist

Before tagging a new release (e.g., `git tag v0.1.0`), ensure the following gates are passed:

## 1. Functional Testing
- [ ] **Unit Tests:** Run `make test` (or `go test ./...`) locally. Must pass.
- [ ] **Integration Tests:** `tests/integration/cli_test.go` must pass.
- [ ] **Manual Verification:**
    - [ ] Run `repo2page octocat/Hello-World` -> Verify `Hello-World.md` created.
    - [ ] Check `--format json` output.
    - [ ] Check `--summary` flag.

## 2. Code Quality
- [ ] **Linting:** Run `go vet ./...`.
- [ ] **Formatting:** Run `go fmt ./...`.
- [ ] **Security:** Check GitHub Actions CodeQL status (if pushed).

## 3. Documentation
- [ ] **CHANGELOG:** Update `CHANGELOG.md` with new features/fixes under the version header.
- [ ] **README:** Ensure usage instructions match current flags and behavior.
- [ ] **Version:** Check `internal/version/version.go` logic (automated by GoReleaser, but good to verify).

## 4. Release Process
1.  **Commit:** Ensure all changes are committed.
2.  **Tag:** `git tag -a v0.1.0 -m "Release v0.1.0"`
3.  **Push:** `git push origin v0.1.0`
4.  **Monitor:** Watch the "Release" GitHub Action.
5.  **Verify:** Check the "Releases" tab on GitHub for the generated artifacts.
