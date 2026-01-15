# Contributing to repo2page

First off, thank you for considering contributing to `repo2page`! It's people like you that make the project great.

## Development Setup

1. **Fork and Clone**:
   ```bash
   git clone https://github.com/YOUR_USERNAME/repo2page.git
   cd repo2page
   ```

2. **Install Go**: Ensure you have Go 1.23+ installed.

3. **Install Dependencies**:
   ```bash
   go mod download
   ```

4. **Run Tests**:
   ```bash
   go test ./...
   ```

## Development Workflow

- **Branching**: Create a feature branch for your changes (`git checkout -b feat/my-new-feature`).
- **Coding Style**: Follow standard Go conventions. Use `go fmt` before committing.
- **Testing**: Add tests for new features or bug fixes.
- **Commits**: Use descriptive commit messages.

## Submitting Pull Requests

1. Push your branch to your fork.
2. Open a Pull Request against the `main` branch.
3. Provide a clear description of the changes and link any related issues.
4. Ensure CI tests pass.

## Licensing
By contributing, you agree that your contributions will be licensed under the project's [MIT License](LICENSE).
