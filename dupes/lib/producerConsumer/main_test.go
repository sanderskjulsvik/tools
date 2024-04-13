package producerConsumer

import (
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/sander-skjulsvik/tools/dupes/lib/common"
)

func TestMain(m *testing.T) {
	common.TestRun("test_main_producer_consumer", Run)
}

func TestGetFiles(t *testing.T) {
	baseDir := "test_get_files/"
	defer os.RemoveAll(filepath.Clean(baseDir))

	// Nesting
	{
		workDir := baseDir + "test_nesting/"
		os.MkdirAll(workDir+filepath.Clean("/folder/folder/"), 0755)
		expectedFilePaths := []string{
			workDir + "nesting_file_name",
			workDir + "folder/" + "nesting_file_name",
			workDir + "folder/" + "folder/" + "nesting_file_name",
		}
		for _, file := range expectedFilePaths {
			common.CreateFile(file, "nesting_file_content")
		}
		calculatedFilePaths := make(chan string)
		go getFiles(workDir, calculatedFilePaths)
		ind := 0
		for calculatedPath := range calculatedFilePaths {
			if !slices.Contains(expectedFilePaths, filepath.ToSlash(calculatedPath)) {
				t.Errorf("calculatedPath: %s, not in expectedPaths: %v", calculatedPath, expectedFilePaths)
			}

			ind++
			if ind > 3 {
				t.Errorf(
					"Expected to find one file, fund %d", ind,
				)
			}
		}
	}

	// Empty file
	{
		workDir := baseDir + "test_emtpy_file/"
		os.MkdirAll(filepath.Clean(workDir), 0755)
		common.CreateEmptyFile(workDir + "empty_file")
		common.CreateFile(workDir+"not_empty_file", "not_empty_file")

		calculatedFilePaths := make(chan string)
		go getFiles(workDir, calculatedFilePaths)
		ind := 0
		for range calculatedFilePaths {
			ind++
		}
		if ind != 2 {
			t.Errorf("Expected 2 files, but got: %d", ind)
		}
	}

	// Symlink
	{
		workDir := baseDir + "test_symlink/"
		os.MkdirAll(filepath.Clean(workDir), 0755)
		common.CreateEmptyFile(workDir + "source_file")
		os.Symlink(workDir+"source_file", workDir+"destination_file")

		calculatedFilePaths := make(chan string)
		go getFiles(workDir, calculatedFilePaths)
		ind := 0
		for calculatedPath := range calculatedFilePaths {
			if filepath.ToSlash(calculatedPath) == workDir+"destination_file" {
				t.Errorf(
					"Calculated path is a symlink: %s", calculatedPath,
				)
			}
			if filepath.ToSlash(calculatedPath) != workDir+"source_file" {
				t.Errorf(
					"Fund file is not source the source file: %s", calculatedPath,
				)
			}
			ind++
			if ind != 1 {
				t.Errorf(
					"Expected to find one file, fund %d", ind,
				)
			}
		}
	}

}
