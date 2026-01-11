repo2page/
├── cmd/
│   └── repo2page/
│       └── main.go                  # CLI entry point (Agent A10)
│
├── internal/
│   ├── core/
│   │   ├── types.go                 # Core types & interfaces (Agent A1)
│   │   ├── detect.go                # Source + connectivity detection (Agent A2)
│   │   ├── converter.go             # Orchestrates conversion pipeline (Agent A9/A10)
│   │   ├── tree.go                  # Deterministic tree resolver (Agent A5)
│   │   ├── reader.go                # File reader & normalizer (Agent A7)
│   │   └── assembler.go             # Page assembly logic (Agent A9)
│   │
│   ├── loader/
│   │   ├── local.go                 # Local filesystem loader (Agent A3)
│   │   └── github.go                # GitHub API loader (Agent A4)
│   │
│   ├── ignore/
│   │   ├── engine.go                # Ignore rules engine (Agent A6)
│   │   └── patterns.go              # Default ignore patterns (Agent A6)
│   │
│   ├── formatter/
│   │   ├── markdown.go              # Markdown formatter (Agent A8)
│   │   ├── html.go                  # HTML formatter (Agent A8)
│   │   └── text.go                  # Plain text formatter (Agent A8)
│   │
│   └── version/
│       └── version.go               # Version, commit, build date
│
├── pkg/
│   └── repo2page/
│       └── api.go                   # Public API (library usage)
│
├── tests/
│   ├── fixtures/
│   │   ├── small-repo/
│   │   │   ├── README.md
│   │   │   └── main.go
│   │   │
│   │   ├── medium-repo/
│   │   │   ├── src/
│   │   │   │   └── app.go
│   │   │   ├── tests/
│   │   │   │   └── app_test.go
│   │   │   └── README.md
│   │   │
│   │   └── with-binary/
│   │       ├── image.png
│   │       └── README.md
│   │
│   ├── golden/
│   │   ├── small-repo.md
│   │   ├── small-repo.html
│   │   ├── small-repo.txt
│   │   ├── medium-repo.md
│   │   └── with-readme.md
│   │
│   └── integration/
│       └── cli_test.go              # End-to-end CLI tests (Agent A11)
│
├── docs/
│   ├── index.md                     # Getting started
│   ├── installation.md             # All install methods
│   ├── usage.md                    # CLI reference
│   ├── formats.md                  # Output format specs
│   ├── configuration.md            # Advanced config
│   ├── examples.md                 # Real-world examples
│   ├── faq.md                      # FAQs
│   ├── troubleshooting.md          # Common issues
│   ├── api.md                      # Library API docs
│   └── contributing.md             # Contribution guide
│
├── .github/
│   ├── workflows/
│   │   ├── ci.yml                   # CI (tests, lint)
│   │   ├── release.yml              # GoReleaser pipeline
│   │   └── codeql.yml               # Security scanning
│   │
│   ├── ISSUE_TEMPLATE/
│   │   ├── bug_report.md
│   │   └── feature_request.md
│   │
│   └── PULL_REQUEST_TEMPLATE.md
│
├── scripts/
│   ├── install.sh                  # Unix install script
│   └── install.ps1                 # Windows install script
│
├── bin/                             # Local build artifacts (gitignored)
│
├── .goreleaser.yml                  # Release config
├── go.mod                           # Go module definition
├── go.sum                           # Dependency checksums
├── Makefile                        # Build/test automation
├── README.md                       # Project overview
├── CHANGELOG.md                    # Release history
├── CONTRIBUTING.md                 # Contribution guide
├── LICENSE                         # MIT license
├── SECURITY.md                     # Security policy
├── .gitignore                      # Git ignore rules
└── .editorconfig                   # Formatting consistency
