package main

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"testing"
)

func TestCLIWrapper(t *testing.T) {
	input := "<html><body><h1>Hello</h1><img src=\"cid:img1\"></body></html>"
	cmd := exec.Command("go", "run", "main.go", "--width", "80")
	cmd.Stdin = bytes.NewReader([]byte(input))
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v\nStderr: %s", err, stderr.String())
	}

	var output renderOutput
	if err := json.Unmarshal(stdout.Bytes(), &output); err != nil {
		t.Fatalf("Failed to parse JSON: %v\nOutput: %s", err, stdout.String())
	}

	if output.RenderedText == "" {
		t.Error("Expected non-empty rendered text")
	}
	if len(output.InlineImages) != 1 {
		t.Errorf("Expected 1 image, got %d", len(output.InlineImages))
	}
}
