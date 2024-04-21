package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func TestRun(path string, run Run) {
	var (
		baseDir            = path
		numLevels          = 1 // Number of levels of nested folders
		numFoldersPerLevel = 2 // Number of folders per level
		numFilesPerFolder  = 2 // Number of files per folder
		content            = "This is the common content for all files in the folder."
	)

	// Create the base directory if it doesn't exist
	defer os.RemoveAll(baseDir)
	if err := os.MkdirAll(baseDir, os.ModePerm); err != nil {
		fmt.Println("Error creating base directory:", err)
		return
	}

	// Generate the nested folder structure
	GenerateNestedStructure(baseDir, numLevels, numFoldersPerLevel, numFilesPerFolder, content)

	fmt.Println("Nested folder structure generated successfully.")
	run(baseDir, false)
	fmt.Printf("Done running! \n")
}

func GenerateNestedStructure(dirPath string, levels, foldersPerLevel, filesPerFolder int, content string) {
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
		GenerateNestedStructure(folderPath, levels-1, foldersPerLevel, filesPerFolder, content)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CreateEmptyFile(path string) {
	d := []byte("")
	check(os.WriteFile(filepath.Clean(path), d, 0644))
}

func CreateFile(path, content string) {
	check(os.WriteFile(filepath.Clean(path), []byte(content), 0644))
}
