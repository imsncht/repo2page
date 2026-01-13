# Contributing to repo2page

Thank you for your interest in contributing to **repo2page**! We welcome contributions from developers of all skill levels.

## Getting Started

1.  **Fork the repository** on GitHub.
2.  **Clone your fork** locally:
    ```bash
    git clone https://github.com/your-username/repo2page.git
    cd repo2page
    ```
3.  **Install Go** (version 1.25 or later).

## Development Workflow

### Building
Use the provided Makefile or standard Go commands:
```bash
# Build binary
make build

# Run directly
go run ./cmd/repo2page [args]
```

### Testing
Run the test suite to ensure everything is working:
```bash
make test
```

### Code Style
We follow standard Go idioms. Please ensure your code is formatted:
```bash
go fmt ./...
```

## Submitting a Pull Request

1.  Create a new branch for your feature or fix (`git checkout -b feature/amazing-feature`).
2.  Commit your changes with clear, descriptive messages.
3.  Push to your fork and submit a Pull Request.
4.  Ensure all tests pass and describe your changes in the PR description.

## Reporting Issues

If you find a bug or have a feature request, please open an issue on the GitHub repository. Provide as much detail as possible, including steps to reproduce the issue.

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
