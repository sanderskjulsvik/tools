package test

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
)

type Folder struct {
	Name    string
	Files   []File
	Folders []Folder
}

type File struct {
	Name    string
	Content string
	Hash    string
}

func (folder *Folder) Generate() error {
	// Create the folder if it doesn't exist
	os.MkdirAll(folder.Name, os.ModePerm)
	err := os.Chdir(folder.Name)
	if err != nil {
		return fmt.Errorf("Error changing directory: %e", err)
	}

	// Create files in the folder
	for _, file := range folder.Files {
		filePath := filepath.Join(folder.Name, file.Name)
		CreateFile(filePath, file.Content)
	}

	// Create child folders
	for _, childFolder := range folder.Folders {
		childFolder.Generate()
	}
	return nil
}

func (folder *Folder) Clean() {
	os.RemoveAll(folder.Name)
}

func (folder *Folder) GetFullFilePaths() []string {
	paths := []string{}
	for _, file := range folder.Files {
		paths = append(paths, filepath.Join(folder.Name, file.Name))
	}
	for _, childFolder := range folder.Folders {
		paths = slices.Concat(paths, childFolder.GetFullFilePaths())
	}
	return paths
}
