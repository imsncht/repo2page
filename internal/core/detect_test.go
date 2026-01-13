package core

import (
	"testing"
)

func TestAutoDetectSource(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    SourceType
		wantErr bool
	}{
		{
			name:  "GitHub URL",
			input: "https://github.com/user/repo",
			want:  SourceGitHub,
		},
		{
			name:  "GitHub Shorthand",
			input: "user/repo",
			want:  SourceGitHub,
		},
		{
			name:  "Current Directory",
			input: ".",
			want:  SourceLocal,
		},
		{
			name:  "Valid Local Path",
			input: ".", // "." exists
			want:  SourceLocal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// For local paths, we need them to exist for AutoDetectSource to return SourceLocal
			// Otherwise it might return error or unknown.
			// Ideally AutoDetect checks file existence.
			
			got, err := AutoDetectSource(tt.input)
			if (err != nil) != tt.wantErr {
				// Special case: if we expect SourceLocal and it returns error because dir doesn't exist, ignore for specific tests 
				// or setup temp dirs? 
				// Let's assume AutoDetectSource logic:
				// 1. Is it GitHub URL/Short? -> GitHub
				// 2. Does file exist? -> Local
				// 3. Error
				
				// For the "Valid Local Path", if it doesn't exist on disk, we might want to skip or mock.
				// For simple unit testing without mocks, I'll rely on the logic I know.
				t.Errorf("AutoDetectSource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AutoDetectSource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseGitHubRepo(t *testing.T) {
	tests := []struct {
		input     string
		wantOwner string
		wantRepo  string
		wantErr   bool
	}{
		{"user/repo", "user", "repo", false},
		{"https://github.com/user/repo", "user", "repo", false},
		{"invalid", "", "", true},
		{"user/repo/extra", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			owner, repo, err := parseGitHubRepo(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseGitHubRepo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if owner != tt.wantOwner {
				t.Errorf("parseGitHubRepo() owner = %v, want %v", owner, tt.wantOwner)
			}
			if repo != tt.wantRepo {
				t.Errorf("parseGitHubRepo() repo = %v, want %v", repo, tt.wantRepo)
			}
		})
	}
}
