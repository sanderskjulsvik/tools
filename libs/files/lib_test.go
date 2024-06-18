package files

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopy(t *testing.T) {
	// Meta
	basePath := "testCopy"
	srcPath := filepath.Join(basePath, "test_src.txt")
	srcConent := []byte("Hello, World!")
	dstPath := filepath.Join(basePath, "test_dst.txt")
	defer os.RemoveAll(basePath)
	// Create a test file
	err := CreateFile(srcPath, srcConent)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	// Defer the deletion of the test file
	// Copy the test file to a new location
	err = Copy(srcPath, dstPath)
	if err != nil {
		t.Fatalf("Error copying file: %v", err)
	}
	// Check if the new file exists
	if _, err := os.Stat(dstPath); os.IsNotExist(err) {
		t.Fatalf("New file does not exist: %v", err)
	}
	// check if the new file has the same content as the original file
	dstContent, err := os.ReadFile(dstPath)
	if err != nil {
		t.Fatalf("Error hashing destination file: %v", err)
	}
	if string(srcConent) != string(dstContent) {
		t.Fatalf("Source and destination file has different content.")
	}
}
