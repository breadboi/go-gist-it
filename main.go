package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	filePath := os.Args[1]

	// Run the GitHub CLI command to create a gist
	cmd := exec.Command("gh", "gist", "create", filePath)
	out, _ := cmd.CombinedOutput()

	// Extract the URL (assuming it's the last word in the output)
	outputLines := strings.Split(strings.TrimSpace(string(out)), "\n")
	gistURL := strings.TrimSpace(outputLines[len(outputLines)-1])

	// Copy the gist URL to the clipboard
	err := clipboard.WriteAll(gistURL)
	if err != nil {
		return
	}
}
