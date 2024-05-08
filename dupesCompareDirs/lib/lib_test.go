package dupescomparedirs_test

import (
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/sander-skjulsvik/tools/dupes/lib/test"
	comparedirs "github.com/sander-skjulsvik/tools/dupesCompareDirs/lib"
)

func getFileMap() map[string]test.File {
	files := []test.File{
		{
			Name:    "Group1",
			Content: "Group1Content",
		},
		{
			Name:    "Group2",
			Content: "Group2Content",
		},
		{
			Name:    "Group3",
			Content: "Group3Content",
		},
		{
			Name:    "Group4",
			Content: "Group4Content",
		},
		{
			Name:    "Group5",
			Content: "Group3Content",
		},
		{
			Name:    "Group5",
			Content: "Group3Content",
		},
	}

	fileMap := make(map[string]test.File)
	for _, file := range files {
		fileMap[file.Name] = file
	}
	return fileMap
}

func setupD1(prefix string) test.Folder {
	testFiles := getFileMap()
	files := []test.File{
		testFiles["Group1"],
		testFiles["Group2"],
		testFiles["Group3"],
	}
	folder := test.Folder{
		Name:  prefix,
		Files: files,
		Folders: []test.Folder{
			{
				Name:  "Folder1",
				Files: files,
				Folders: []test.Folder{
					{
						Name:  "Folder1",
						Files: files,
						Folders: []test.Folder{
							{
								Name:    "Folder1",
								Files:   files,
								Folders: []test.Folder{},
							},
						},
					},
				},
			},
			{
				Name:  "Folder2",
				Files: files,
				Folders: []test.Folder{
					{
						Name:  "Folder1",
						Files: files,
						Folders: []test.Folder{
							{
								Name:    "Folder2",
								Files:   files,
								Folders: []test.Folder{},
							},
							{
								Name:  "Folder2",
								Files: files,
								Folders: []test.Folder{
									{
										Name:    "Folder2",
										Files:   files,
										Folders: []test.Folder{},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	folder.Generate()
	return folder
}

func setupD2(prefix string) test.Folder {
	testFiles := getFileMap()
	files := []test.File{
		testFiles["Group1"],
		testFiles["Group2"],
		testFiles["Group4"],
		testFiles["Group5"],
	}

	folder := test.Folder{
		Name:  prefix,
		Files: files,
		Folders: []test.Folder{
			{
				Name:  "Folder1",
				Files: files,
				Folders: []test.Folder{
					{
						Name:  "Folder1",
						Files: files,
						Folders: []test.Folder{
							{
								Name:    "Folder1",
								Files:   files,
								Folders: []test.Folder{},
							},
						},
					},
				},
			},
			{
				Name:  "Folder2",
				Files: files,
				Folders: []test.Folder{
					{
						Name:  "Folder1",
						Files: files,
						Folders: []test.Folder{
							{
								Name:    "Folder2",
								Files:   files,
								Folders: []test.Folder{},
							},
							{
								Name:  "Folder2",
								Files: files,
								Folders: []test.Folder{
									{
										Name:    "Folder2",
										Files:   files,
										Folders: []test.Folder{},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	folder.Generate()
	return folder
}

func setup(rootPath string) (test.Folder, test.Folder) {
	p1 := filepath.Join(rootPath, "d1")
	p2 := filepath.Join(rootPath, "d2")
	d1 := setupD1(p1)
	d2 := setupD2(p2)
	return d1, d2
}

func cleanUp(rootPath string) {
	os.RemoveAll(rootPath)
}

// OnlyInboth returns dupes that is present in both directories
func TestOnlyInboth(t *testing.T) {
	rootPath := "test_only_in_both"
	d1, d2 := setup(rootPath)
	defer cleanUp(rootPath)

	calcDupes := comparedirs.OnlyInboth(
		filepath.Join(rootPath, "d1"),
		filepath.Join(rootPath, "d2"),
	)
	if len(calcDupes.D) != 2 {
		t.Errorf("Expected 2 dupes, got %d", len(calcDupes.D))
	}

	allPaths := slices.Concat(d1.GetFullFilePaths(), d2.GetFullFilePaths())
	expectedPaths := []string{}
	for _, path := range allPaths {
		if filepath.Base(path) == "Group3" {
			expectedPaths = append(expectedPaths, path)
		}
	}

	calcPaths := []string{}
	for _, dupe := range calcDupes.D {
		calcPaths = slices.Concat(calcPaths, dupe.Paths)
	}
	if !slices.Equal(calcPaths, expectedPaths) {
		t.Errorf("Expected %v, got %v", expectedPaths, calcPaths)
	}
}

// OnlyInFirst returns dupes that is only present in first directory
func TestOnlyInFirst(t *testing.T) {
	rootPath := "test_only_in_first"
	d1, d2 := setup(rootPath)
	defer cleanUp(rootPath)

	calcDupes := comparedirs.OnlyInFirst(
		filepath.Join(rootPath, "d1"),
		filepath.Join(rootPath, "d2"),
	)

	if len(calcDupes.D) != 1 {
		t.Errorf("Expected 1 dupes, got %d", len(calcDupes.D))
	}

	allPaths := slices.Concat(d1.GetFullFilePaths(), d2.GetFullFilePaths())
	expectedPaths := []string{}
	for _, path := range allPaths {
		if filepath.Base(path) == "Group1" || filepath.Base(path) == "Group2" {
			expectedPaths = append(expectedPaths, path)
		}
	}
	calcPaths := []string{}
	for _, dupe := range calcDupes.D {
		calcPaths = slices.Concat(calcPaths, dupe.Paths)
	}
	if !slices.Equal(calcPaths, expectedPaths) {
		t.Errorf("Expected %v, got %v", expectedPaths, calcPaths)
	}
}

// All returns all dupes in both directories
func TestAll(t *testing.T) {
	rootPath := "test_all"
	d1, d2 := setup(rootPath)
	defer cleanUp(rootPath)

	calcDupes := comparedirs.OnlyInboth(
		filepath.Join(rootPath, "d1"),
		filepath.Join(rootPath, "d2"),
	)
	if len(calcDupes.D) != 2 {
		t.Errorf("Expected 2 dupes, got %d", len(calcDupes.D))
	}

	allPaths := slices.Concat(d1.GetFullFilePaths(), d2.GetFullFilePaths())
	calcPaths := []string{}
	for _, dupe := range calcDupes.D {
		calcPaths = slices.Concat(calcPaths, dupe.Paths)
	}
	if !slices.Equal(calcPaths, allPaths) {
		t.Errorf("Expected %v, got %v", allPaths, calcPaths)
	}
}
