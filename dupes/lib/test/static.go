package test

import (
	"fmt"
	"os"
	"strings"

	"github.com/sander-skjulsvik/tools/dupes/lib/common"
)

type Folder struct {
	Name    string
	Files   []File
	Folders []Folder
}

type File struct {
	Name    string
	Content string
}

func SetupExpectedDupes(path string) {
	// Create the base directory if it doesn't exist
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		fmt.Println("Error creating base directory:", err)
		return
	}
	// Create the nested folder structure from ExectedDupes
	for hash, dupe := range ExpectedDupes.D {
		for _, filePath := range dupe.Paths {
			// Create the parent folders if they don't exist
			CrateParentFolders(filePath)
			// Create the file with content from ExpectedDupesHashMap
			CreateFile(filePath, ExpectedDupesHashMap[hash])

		}
	}
}

func CrateParentFolders(path string) {
	// Remove  entry in path
	parentPath := strings.Join(strings.Split(path, "/")[:len(strings.Split(path, "/"))-1], "/")

	// Create the base directory if it doesn't exist
	if err := os.MkdirAll(parentPath, os.ModePerm); err != nil {
		fmt.Println("Error creating base directory:", err)
		return
	}
}

var RootPath string = "TestStructureRoot/"

var ExpectedDupesHashMap map[string]string = map[string]string{
	"1ff0d1ef84204e0fd88c39dd6efb1ba449c6d4a4841f2906425515412cf6178b": "Group1Content",
	"5f8009bbd085f744953a9c8a7983f19e0f16edd275505b15af35953240fb5502": "Group2Content",
	"1ee71ec5183be6c5b38bfe53319dcd950309947e20e21a228ffe5abc269b186e": "Group3Content",
	"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855": "",
}

var ExpectedDupes common.Dupes = common.Dupes{
	D: map[string]*common.Dupe{
		"1ff0d1ef84204e0fd88c39dd6efb1ba449c6d4a4841f2906425515412cf6178b": {
			Hash: "1ff0d1ef84204e0fd88c39dd6efb1ba449c6d4a4841f2906425515412cf6178b",
			Paths: []string{
				RootPath + "Folder1/Group1",
				RootPath + "Folder1/Folder1/Group1",
				RootPath + "Folder2/Group1",
				RootPath + "Folder2/Folder1/Folder2/Folder2/Group1",
				RootPath + "Group1_2",
			},
		}, "5f8009bbd085f744953a9c8a7983f19e0f16edd275505b15af35953240fb5502": {
			Hash: "5f8009bbd085f744953a9c8a7983f19e0f16edd275505b15af35953240fb5502",
			Paths: []string{
				RootPath + "Folder1/Group2",
				RootPath + "Folder1/Folder1/Group2",
				RootPath + "Folder2/Group2",
				RootPath + "Folder2/Folder1/Folder2/Folder2/Group2",
				RootPath + "Folder2/Folder1/Folder2/Folder2/Group2_2",
				// Same name as files in group1
				RootPath + "Folder2/Folder1/Folder2/Group1",
				RootPath + "Group2_2",
			},
		}, "1ee71ec5183be6c5b38bfe53319dcd950309947e20e21a228ffe5abc269b186e": {
			Hash: "1ee71ec5183be6c5b38bfe53319dcd950309947e20e21a228ffe5abc269b186e",
			Paths: []string{
				// Only one file
				RootPath + "Folder2/Group3",
			},
		}, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855": {
			Hash: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			Paths: []string{
				RootPath + "Folder2/emtpy1",
				RootPath + "Folder2/Folder1/emtpy1",
				RootPath + "Folder2/Folder1/Folder2/emtpy1",
			},
		},
	},
}
