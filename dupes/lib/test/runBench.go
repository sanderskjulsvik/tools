package test

import (
	"crypto/rand"
	"fmt"
	"os"
	"testing"

	"github.com/sander-skjulsvik/tools/dupes/lib/common"
)

func TestBench(path string, run common.Run, t *testing.T) {
	// Clean up the test files after the test is done
	defer os.RemoveAll(path)
	// Setup the expected dupes
	SetupBench(path)
	// Run the run function to find the dupes
}

func SetupBench(path string) {
}

func TestBenchLargeFiles() {}

func TestBenchMixed() {}

// doing single lvl, assuming deep files does not make a difference
func generateFiles(path string, numberOfFiles, sizeOfFilesB, groupSize int) {
	for i := 0; i < numberOfFiles; i++ {
		createFileRandom(fmt.Sprint(i), sizeOfFilesB)
		for range groupSize - 1 {
			file.copy
			i++
		}
	}
}

func createFileRandom(path string, size int) {
	chunkSize := size
	if size < 1e4 {
		chunkSize = 1e4
	}
	fp, err := os.Open(path)
	check(err)
	defer fp.Close()
	written := 0
	for written < size {
		chunk := make([]byte, chunkSize)
		rand.Read(chunk)
		w, writeErr := fp.Write(chunk)
		written += w
		if writeErr != nil {
			check(fmt.Errorf("Error writing to file: %v", writeErr))
		}
	}
}
