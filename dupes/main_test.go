package main_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	main "github.com/sander-skjulsvik/tools/dupes"
)

func TestMain(m *testing.M) {
	var (
		baseDir            = "./testing_nested_structure"
		numLevels          = 5 // Number of levels of nested folders
		numFoldersPerLevel = 5 // Number of folders per level
		numFilesPerFolder  = 5 // Number of files per folder
		content            = "This is the common content for all files in the folder."
	)

	// Create the base directory if it doesn't exist
	if err := os.MkdirAll(baseDir, os.ModePerm); err != nil {
		fmt.Println("Error creating base directory:", err)
		return
	}

	// Generate the nested folder structure
	generateNestedStructure(baseDir, numLevels, numFoldersPerLevel, numFilesPerFolder, content)

	fmt.Println("Nested folder structure generated successfully.")
	main.Run(baseDir, "single")
}

func generateNestedStructure(dirPath string, levels, foldersPerLevel, filesPerFolder int, content string) {
	if levels <= 0 {
		return
	}

	for i := 1; i <= foldersPerLevel; i++ {
		folderName := fmt.Sprintf("Folder%d", i)
		folderPath := filepath.Join(dirPath, folderName)

		// Create the folder if it doesn't exist
		if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}

		// Create files in the folder
		for j := 1; j <= filesPerFolder; j++ {
			fileName := fmt.Sprintf("File%d.txt", j)
			filePath := filepath.Join(folderPath, fileName)

			// Check if this file should have unique content
			if i%2 == 0 {
				fileContent := fmt.Sprintf("%s\n%s", content, strings.Repeat(fmt.Sprintf("UniqueContent%d", j), 5))
				err := os.WriteFile(filePath, []byte(fileContent), os.ModePerm)
				if err != nil {
					fmt.Println("Error creating file:", err)
					return
				}
			} else {
				// Files in odd-numbered folders have the same content
				err := os.WriteFile(filePath, []byte(content), os.ModePerm)
				if err != nil {
					fmt.Println("Error creating file:", err)
					return
				}
			}
		}

		// Recursively generate the nested structure for the next level
		generateNestedStructure(folderPath, levels-1, foldersPerLevel, filesPerFolder, content)
	}
}
