package files

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	math "github.com/sander-skjulsvik/tools/libs/math"
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

func CreateFile(path string, content []byte) error {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return fmt.Errorf("unable to create dir: %w", err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to create file: %w", err)
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		return fmt.Errorf("unable to write to file: %w", err)
	}
	return nil
}

func TestCreateFileLarge(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	// Meta
	basePath := "testCreateFileLarge"
	filePath := filepath.Join(basePath, "test_large.txt")
	fileSize := int64(20*1e9 - 1) // B
	defer os.RemoveAll(basePath)

	// Create a large file

	err := CreateLargeFile(filePath, fileSize)
	if err != nil {
		t.Fatalf("Error creating large file: %v", err)
	}

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Fatalf("New file does not exist: %v", err)
	}
	// Check if the file has the correct size
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		t.Fatalf("Error getting file info: %v", err)
	}
	if math.Abs(fileInfo.Size()-fileSize) > WRITE_BUFFER_SIZE {
		t.Fatalf("File size is not correct: File size: %d, Expected file Size: %d, diff: %d,  err: %v", fileInfo.Size(), fileSize, fileInfo.Size()-fileSize, err)
	}
	t.Logf("correct: File size: %d, Expected file Size: %d, diff: %d", fileInfo.Size(), fileSize, fileInfo.Size()-fileSize)
}
