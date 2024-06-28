package test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	sfiles "github.com/sander-skjulsvik/tools/libs/files"
)

func TestGenerateFiles(t *testing.T) {
	const (
		basePath      = "test_generate_files"
		numberOffiles = 10
		uniqueFiles   = 9
		fileSize      = 1e3
	)
	defer os.RemoveAll(basePath)

	err := generateFiles(basePath, numberOffiles, int(fileSize), uniqueFiles)
	if err != nil {
		t.Fatalf("Failed to generate files: %v", err)
	}

	// Test if 1.txt and 10.txt is equal
	b, err := sfiles.FilesEqual(
		filepath.Join(basePath, "1.txt"),
		filepath.Join(basePath, "10.txt"),
	)
	if err != nil {
		t.Fatal("Failed to compare files 1.txt and 10.txt")
	}
	if !b {
		t.Fatal("1 and 10 is not equal")
	}

	// Test if 1 is equal to any other
	for i := 2; i < numberOffiles; i++ {
		b, err := sfiles.FilesEqual(
			filepath.Join(basePath, "1.txt"),
			filepath.Join(basePath, fmt.Sprintf("%d.txt", i)),
		)
		if err != nil {
			t.Fatalf("Failed to compare files 1.txt and %d.txt", i)
		}
		if b {
			t.Fatalf("Unequal files are equal: 1.txt and %d.txt", i)
		}
	}

}
