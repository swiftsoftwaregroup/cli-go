package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestCliE2E(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "cli-e2e-test")
	if err != nil {
		t.Fatalf("Could not create temp dir: %s", err)
	}
	defer os.RemoveAll(tempDir)

	// Test cases
	testCases := []struct {
		name     string
		args     []string
		input    string
		expected string
	}{
		{
			name:     "Default greeting",
			args:     []string{"greet"},
			input:    "Chris",
			expected: "Hello, Chris!",
		},
		{
			name:     "Spanish greeting",
			args:     []string{"greet", "--language", "es"},
			input:    "Velina",
			expected: "Hola, Velina!",
		},
		{
			name:     "Bulgarian greeting",
			args:     []string{"greet", "-l", "bg"},
			input:    "Ема",
			expected: "Здравей, Ема!",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create input file
			inputFile := filepath.Join(tempDir, "input.txt")
			if err := os.WriteFile(inputFile, []byte(tc.input), 0644); err != nil {
				t.Fatalf("Could not write input file: %s", err)
			}

			// Prepare command
			args := append(tc.args, inputFile)
			cmd := exec.Command("./cli-go", args...)

			// Run command and capture output
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("Command failed: %s", err)
			}

			// Check output
			actual := strings.TrimSpace(string(output))
			if actual != tc.expected {
				t.Errorf("Expected output %q, got %q", tc.expected, actual)
			}
		})
	}
}
