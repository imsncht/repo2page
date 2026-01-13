package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"repo2page/internal/core"
	"repo2page/internal/formatter"
	"repo2page/internal/ui"
	"repo2page/internal/version"
)

const (
	exitSuccess         = 0
	exitInvalidArgs     = 1
	exitSourceFailure   = 2
	exitConversionFail  = 3
	exitOutputFail      = 4
	exitNetworkFail     = 5
	exitUnexpectedError = 99
)

func main() {
	var (
		format       = flag.String("format", "md", "Output format: md, html, txt")
		output       = flag.String("output", "", "Output file path")
		branch       = flag.String("branch", "", "Git branch (auto-detected if empty)")
		commit       = flag.String("commit", "", "Specific commit hash")
		summary      = flag.Bool("summary", false, "Summary mode (no file contents)")
		noTree       = flag.Bool("no-tree", false, "Disable directory tree")
		maxFileSize  = flag.Int("max-file-size", 500, "Max file size to include (KB)")
		jsonOutput   = flag.Bool("json", false, "Emit machine-readable JSON summary")
		showVersion  = flag.Bool("version", false, "Show version information")
	)

	var excludes multiFlag
	flag.Var(&excludes, "exclude", "Exclude pattern (can be repeated)")

	flag.Parse()

	if *showVersion {
		fmt.Println(version.Full())
		os.Exit(exitSuccess)
	}

	args := flag.Args()
	if len(args) != 1 {
		printError("exactly one source argument is required")
		os.Exit(exitInvalidArgs)
	}

	source := args[0]

	outFormat := core.OutputFormat(strings.ToLower(*format))
	if outFormat != core.FormatMarkdown &&
		outFormat != core.FormatHTML &&
		outFormat != core.FormatText &&
		outFormat != core.FormatJSON {
		printError("invalid format: " + *format)
		os.Exit(exitInvalidArgs)
	}

	input := core.RepoInput{
		PathOrURL: source,
		Branch:    *branch,
		Commit:    *commit,
	}

	options := core.ConvertOptions{
		Format:             outFormat,
		IncludeTree:        !*noTree,
		IncludeReadmeFirst: true,
		ExcludePatterns:    excludes,
		MaxFileSizeKB:      *maxFileSize,
		SummaryMode:        *summary,
		OutputPath:         *output,
		ProgressCallback: func(total int64, r io.Reader) io.Reader {
			bar := ui.NewProgressBar(total, "Downloading repository")
			bar.Start()
			return ui.NewProxyReader(r, bar)
		},
	}

	result, err := core.Convert(input, options)
	if err != nil {
		handleError(err)
		os.Exit(exitConversionFail)
	}

	var content string

	switch options.Format {
	case core.FormatText:
		content = formatter.RenderText(
			result.RepoName,
			result.Source,
			result.Tree,
			result.Files,
			result.Stats,
			result.Warnings,
		)
	case core.FormatMarkdown:
		content = formatter.RenderMarkdown(
			result.RepoName,
			result.Source,
			result.Tree,
			result.Files,
			result.Stats,
			result.Warnings,
		)
	case core.FormatHTML:
		content = formatter.RenderHTML(
			result.RepoName,
			result.Source,
			result.Tree,
			result.Files,
			result.Stats,
			result.Warnings,
		)
	case core.FormatJSON:
		var err error
		content, err = formatter.RenderJSON(
			result.RepoName,
			result.Source,
			result.Tree,
			result.Files,
			result.Stats,
			result.Warnings,
		)
		if err != nil {
			printError("failed to flatten to JSON: " + err.Error())
			os.Exit(exitUnexpectedError)
		}
	default:
		printError("unsupported format")
		os.Exit(exitInvalidArgs)
	}

	// Determine output path
	outPath := *output
	if outPath == "" {
		base := filepath.Base(strings.TrimSuffix(source, filepath.Ext(source)))
		outPath = base + "." + string(outFormat)
	}

	// Write output
	if err := os.WriteFile(outPath, []byte(content), 0644); err != nil {
		printError("failed to write output file: " + err.Error())
		os.Exit(exitOutputFail)
	}

	// Machine-readable JSON output
	if *jsonOutput {
		j := map[string]any{
			"status":  "success",
			"files":   result.Stats.FileCount,
			"lines":   result.Stats.TotalLines,
			"size_kb": result.Stats.TotalSizeKB,
			"skipped": result.Stats.SkippedFiles,
			"output":  outPath,
			"warnings": result.Warnings,
		}

		data, _ := json.MarshalIndent(j, "", "  ")
		fmt.Println(string(data))
		os.Exit(exitSuccess)
	}

	// Human-readable output
	fmt.Println("✓ Conversion completed successfully")
	fmt.Printf("✓ Output: %s\n\n", outPath)
	fmt.Println("Statistics:")
	fmt.Printf("  Files:   %d\n", result.Stats.FileCount)
	fmt.Printf("  Lines:   %d\n", result.Stats.TotalLines)
	fmt.Printf("  Size:    %d KB\n", result.Stats.TotalSizeKB)
	fmt.Printf("  Skipped: %d\n", result.Stats.SkippedFiles)

	if len(result.Warnings) > 0 {
		fmt.Println("\nWarnings:")
		for _, w := range result.Warnings {
			fmt.Println("  -", w)
		}
	}

	os.Exit(exitSuccess)
}

// --------------------
// Helpers
// --------------------

type multiFlag []string

func (m *multiFlag) String() string {
	return strings.Join(*m, ", ")
}

func (m *multiFlag) Set(value string) error {
	*m = append(*m, value)
	return nil
}

func printError(msg string) {
	fmt.Fprintln(os.Stderr, "✗ Error:", msg)
	fmt.Fprintln(os.Stderr, "→ Run with --help for usage information")
}

func handleError(err error) {
	printError(err.Error())
}
