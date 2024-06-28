package test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/sander-skjulsvik/tools/dupes/lib/common"
	sfiles "github.com/sander-skjulsvik/tools/libs/files"
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
func generateFiles(path string, numberOfFiles, size, uniqueFiles int) error {
	i := 1
	for ; i <= uniqueFiles; i++ {
		err := sfiles.CreateLargeFile(
			filepath.Join(path, fmt.Sprintf("%d.txt", i)),
			int64(size),
			int64(i+1),
		)
		if err != nil {
			return err
		}
	}
	for ; i <= numberOfFiles; i++ {
		err := sfiles.Copy(
			filepath.Join(path, fmt.Sprintf("%d.txt", i%uniqueFiles)),
			filepath.Join(path, fmt.Sprintf("%d.txt", i)),
		)
		if err != nil {
			return err
		}
	}
	return nil
}
