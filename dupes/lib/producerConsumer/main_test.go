package producerConsumer

import (
	"os"
	"slices"
	"testing"

	"github.com/sander-skjulsvik/tools/dupes/lib/common"
)

func TestMain(m *testing.T) {
	common.TestRun(common.DEFAULT_TEST_DIR+"test_main_producer_consumer", Run)
}

func TestGetFiles(t *testing.T) {
	baseDir := common.DEFAULT_TEST_DIR + "test_get_files/"
	defer os.RemoveAll(baseDir)
	// Setup

	// Nesting
	{
		workDir := baseDir + "nesting/"
		defer os.RemoveAll(workDir)
		os.MkdirAll(workDir+"folder/folder/", 0755)
		expectedFilePaths := []string{
			"nesting_file_name",
			"folder/" + "nesting_file_name",
			"folder/" + "folder/" + "nesting_file_name",
		}
		for _, file := range expectedFilePaths {
			common.CreateFile(workDir+file, "nesting_file_content")
		}
		common.CreateFile(workDir+"nesting_file_name", "nesting_file_content")
		common.CreateFile(workDir+"folder/"+"nesting_file_name", "nesting_file_content")
		common.CreateFile(workDir+"folder/"+"folder/"+"nesting_file_name", "nesting_file_content")
		calculatedFilePaths := make(chan string)
		getFiles(workDir, calculatedFilePaths)
		ind := 0
		for calculatedPath := range calculatedFilePaths {
			if !slices.Contains(expectedFilePaths, calculatedPath) {
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
		workDir := baseDir + "empty_file/"
		defer os.RemoveAll(workDir)
		os.MkdirAll(workDir, 0755)
		common.CreateEmptyFile(workDir + "empty_file")
		common.CreateFile(workDir+"not_empty_file", "not_empty_file")

		calculatedFilePaths := make(chan string)
		getFiles(workDir, calculatedFilePaths)
		ind := 0
		for calculatedPath := range calculatedFilePaths {
			if calculatedPath != workDir+"not_empty_file" {
				t.Errorf(
					"Calculated path not as expected. Expected: %s, got: %s",
					workDir+"not_empty_file", calculatedPath,
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

	// Symlink
	{
		workDir := baseDir + "empty_file/"
		defer os.RemoveAll(workDir)
		os.MkdirAll(workDir, 0755)
		common.CreateEmptyFile(workDir + "source_file")
		os.Symlink(workDir+"source_file", workDir+"destination_file")

		calculatedFilePaths := make(chan string)
		getFiles(workDir, calculatedFilePaths)
		ind := 0
		for calculatedPath := range calculatedFilePaths {
			if calculatedPath == workDir+"destination_file" {
				t.Errorf(
					"Calculated path is a symlink: %s", calculatedPath,
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
