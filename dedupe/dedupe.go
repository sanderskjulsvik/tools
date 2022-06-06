package dedupe

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

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
	for md5, files := range d.GetDupes() {
		fmt.Printf("%s: ", md5)
		for _, file := range files {
			//fmt.Printf("    %d: %s%s", i, file.Path, os_spec.LineBreak)
			fmt.Printf("%s ", file.Path)
		}
		fmt.Println()
	}
}

func (d *Dupes) BuildDuplicates(files []File) {
	d.m = make(map[string][]*File)
	for _, file := range files {
		d.AddFile(file)
	}
}

func (d *Dupes) GetDupes() map[string][]*File {
	dupes := make(map[string][]*File)
	for md5, files := range d.m {
		if len(files) <= 1 {
			continue
		}
		dupes[md5] = files
	}
	return dupes
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
		log.Printf("(%d) Path: %s %s", i, path, os_spec.LineBreak)
		file := File{
			Name: info.Name(),
			Path: path,
			Md5:  HashFile(path),
		}
		files = append(files, file)
		return nil
	})

	if err != nil {
		log.Fatal("Failed to parse files!")
	}
	return files
}

func HashFile(path string) string {
	h := md5.New()
	buff := make([]byte, 1024*1024) // Buffer size 1 mb
	fmt.Printf("Opening file: %s%s", path, os_spec.LineBreak)
	fp, err := os.Open(path)
	if err != nil {
		log.Fatalf("Culd not read file: %s%s  err: %s%s", path, os_spec.LineBreak, err.Error(), os_spec.LineBreak)
	}
	defer func() {
		_ = fp.Close()
	}()

	for {
		nBytestRead, err := fp.Read(buff)
		if err != nil && err != io.EOF {
			log.Fatalf("Error reading file: %s%s err: %s%s", path, os_spec.LineBreak, err.Error(), os_spec.LineBreak)
		}
		h.Write(buff[:nBytestRead])
		if err == io.EOF {
			fmt.Println("  EOF")
			break
		}
		fmt.Printf("  Bytes read: %d%s", nBytestRead, os_spec.LineBreak)
	}
	return hex.EncodeToString(h.Sum(nil))
}

func DeDupe(path string) {

	files := GetFiles(path)

	for _, file := range files {
		fmt.Printf("File: %s, Md5: %s %s", file.Path, file.Md5, os_spec.LineBreak)
	}

}

func Run() {
	root := "tmp_test_dir/"
	BuildFolderTree(root)
	// h := HashFile("/home/sander/github.com/sander-skjulsvik/tools/tmp_test_dir/test_dir1/test_file1")
	// fmt.Println(h)
	dupes := Dupes{}
	files := GetFiles(root)
	dupes.BuildDuplicates(files)
	fmt.Println("Printing duplicates")
	dupes.PrintDuplicates()
}
