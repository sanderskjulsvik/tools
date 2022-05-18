package dedupe

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/sander-skjulsvik/tools/os_spec"
)

type File struct {
	name string
	path string
	md5  string
}

func GetMD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetFileContent(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Printf("GetFileContent: Cloud not read file: %s", path)
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
			name: info.Name(),
			path: path,
			md5:  GetMD5(content),
		}
		files = append(files, file)
		return nil
	})

	if err != nil {
		log.Fatal("Failed to parse files!")
	}
	return files
}

func FindDuplicates() {

}

func DeDupe(path string) {

	files := GetFiles(path)

	for _, file := range files {
		fmt.Printf("File: %s, md5: %s%s", file.path, file.md5, os_spec.LineBreak)
	}

}
