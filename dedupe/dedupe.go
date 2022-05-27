package dedupe

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"crypto"

	"github.com/sander-skjulsvik/tools/os_spec"
)

type File struct {
	Name string
	Path string
	Md5  string
}

type Dupes struct {
	m map[string][]*File
}

func (d *Dupes) AddFile(file File) {
	// If Md5 does not exist, make list
	if _, ok := d.m[file.Md5]; !ok {
		d.m[file.Md5] = []*File{}
	}
	// Add file to list
	d.m[file.Md5] = append(d.m[file.Md5], &file)
}

func (d *Dupes) PrintDuplicates() {
	for Md5, files := range d.m {
		if len(files) > 1 {
			fmt.Printf("md5: %s%s", Md5, os_spec.LineBreak)
			for i, file := range files {
				fmt.Printf("    %d: %s", i, file.Path)
			}
		}
	}
}

func (d *Dupes) BuildDuplicates(files []File) {
	d.m = make(map[string][]*File)
	for _, file := range files {
		d.AddFile(file)
	}
}

func GetMD5(text string) string {
	hash := crypto.MD5.New().Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetFileContent(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Printf("GetFileContent: Cloud not read file: %s, err: %s", path, err.Error())
	}
	return string(content), err
}

func GetFiles(path string) []File {
	i := 0
	var files []File
	// https://siongui.github.io/2016/02/04/go-walk-all-files-in-directory/
	err := filepath.Walk(path, func(path string, info os.FileInfo, e error) error {
		i++
		if e != nil {
			return e
		}
		if info.IsDir() {
			return nil
		}
		content, err := GetFileContent(path)
		log.Printf("(%d) Path: %s %s", i, path, os_spec.LineBreak)
		if err != nil {
			return nil
		}
		file := File{
			Name: info.Name(),
			Path: path,
			Md5:  GetMD5(content),
		}
		files = append(files, file)
		return nil
	})

	if err != nil {
		log.Fatal("Failed to parse files!")
	}
	return files
}

func DeDupe(path string) {

	files := GetFiles(path)

	for _, file := range files {
		fmt.Printf("File: %s, Md5: %s%s", file.Path, file.Md5, os_spec.LineBreak)
	}

}

func Run() {
	root := "tmp_test_dir/"
	BuildFolderTree(root)
	files := GetFiles(root)
	dupes := Dupes{}
	dupes.BuildDuplicates(files)
	dupes.PrintDuplicates()
}
