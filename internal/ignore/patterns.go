package ignore

// DefaultIgnorePatterns are always applied.
// Order does not matter; priority is handled in the engine.
var DefaultIgnorePatterns = []string{
	// Version control
	".git/",
	".svn/",
	".hg/",
	".bzr/",

	// Dependencies
	"node_modules/",
	"vendor/",
	".venv/",
	"venv/",
	"env/",
	"__pycache__/",

	// Build artifacts
	"dist/",
	"build/",
	"out/",
	"target/",
	"bin/",
	"obj/",
	"*.o",
	"*.so",
	"*.dylib",
	"*.dll",
	"*.exe",

	// IDE & editor
	".vscode/",
	".idea/",
	"*.swp",
	"*.swo",
	"*~",
	".DS_Store",
	"Thumbs.db",

	// Lock files
	"package-lock.json",
	"yarn.lock",
	"pnpm-lock.yaml",
	"Gemfile.lock",
	"Cargo.lock",
	"poetry.lock",

	// Logs & temp
	"*.log",
	"logs/",
	"tmp/",
	"temp/",
	"*.tmp",
	"*.cache",

	// Binary & media
	"*.jpg",
	"*.jpeg",
	"*.png",
	"*.gif",
	"*.pdf",
	"*.zip",
	"*.tar.gz",
	"*.rar",
	"*.7z",
	"*.mp4",
	"*.mp3",
	"*.mov",
	"*.avi",
}
