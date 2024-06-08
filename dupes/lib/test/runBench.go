package test

import (
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
func generateFiles(path string, numberOfFiles, sizeOfFilesB, groupSizes int) {
}

func createFileRandom(path string, size int) {
	fp := os.Open(path)
	defer fp.Close()
	fp.Write([]byte("This is the common content for all files in the folder."))
}
