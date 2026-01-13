package ui

import (
	"fmt"
	"io"
	"strings"
	"sync"
	"time"
)

// Progress indicates the state of an operation
type Progress struct {
	w       io.Writer
	mu      sync.Mutex
	done    chan bool
	ticker  *time.Ticker
	message string
	
	// For bar
	total   int64
	current int64
	isBar   bool
}

// NewProgressBar creates a new progress bar
func NewProgressBar(total int64, message string) *Progress {
	return &Progress{
		w:       io.Writer(nil), // default to stdout if needed, acts as "std" usually
		message: message,
		total:   total,
		isBar:   true,
		done:    make(chan bool),
	}
}

// NewSpinner creates a new indeterminate spinner
func NewSpinner(message string) *Progress {
	return &Progress{
		message: message,
		isBar:   false,
		done:    make(chan bool),
	}
}

// Start begins the progress animation
func (p *Progress) Start() {
	p.ticker = time.NewTicker(100 * time.Millisecond)
	go p.run()
}

// Stop ends the progress animation
func (p *Progress) Stop() {
	if p.ticker != nil {
		p.ticker.Stop()
		p.done <- true
		close(p.done)
		fmt.Print("\r\033[K") // Clear line
	}
}

// Add updates the progress
func (p *Progress) Add(n int64) {
	p.mu.Lock()
	p.current += n
	p.mu.Unlock()
}

func (p *Progress) run() {
	chars := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	i := 0
	
	for {
		select {
		case <-p.done:
			return
		case <-p.ticker.C:
			p.mu.Lock()
			current := p.current
			total := p.total
			msg := p.message
			p.mu.Unlock()

			if p.isBar && total > 0 {
				percent := float64(current) / float64(total) * 100
				width := 30
				completed := int(float64(width) * percent / 100)
				bar := strings.Repeat("█", completed) + strings.Repeat("░", width-completed)
				fmt.Printf("\r%s [%s] %.1f%%", msg, bar, percent)
			} else {
				// Spinner
				fmt.Printf("\r%s %s", chars[i%len(chars)], msg)
				i++
			}
		}
	}
}

// Reader wraps an io.Reader to update progress
type Reader struct {
	r   io.Reader
	bar *Progress
}

func (r *Reader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	r.bar.Add(int64(n))
	if err == io.EOF {
		r.bar.Stop()
	}
	return n, err
}

func NewProxyReader(r io.Reader, bar *Progress) io.Reader {
	return &Reader{r: r, bar: bar}
}
