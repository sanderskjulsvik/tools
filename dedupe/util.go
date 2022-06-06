package dedupe

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"
)

func Copy(src string, dst string) {
	fin, err := os.Open(src)
	if err != nil {
		log.Fatalf("Copy: Could not open file: %s, err: %s", src, err.Error())
	}
	fout, err := os.Create(dst)
	if err != nil {
		log.Fatalf("Copy: Could not create file: %s, err: %s", src, err.Error())
	}
	_, err = io.Copy(fout, fin)
	if err != nil {
		log.Fatalf("Copy: Could not Copy to file: %s, err: %s", src, err.Error())
	}

}

func BuildFolderTree(root string) {
	os.MkdirAll(root, 0740)
	// Setup dirs
	dirs := []string{
		"test_dir1/",
		"test_dir2/",
		"test_dir2/test_dir21/",
		"test_dir3/",
		"test_dir3/test_dir31/",
		"test_dir3/test_dir32/",
		"test_dir3/test_dir31/test_dir321/",
		"test_dir3/test_dir31/test_dir322/",
		"test_dir3/test_dir32/test_dir321/",
		"test_dir3/test_dir32/test_dir322/",
	}
	for _, dir := range dirs {
		os.MkdirAll(root+dir, 0740)
	}
	// Build unique files
	for _, dir := range dirs {
		for i := 1; i <= 3; i++ {
			path := root + dir + "test_file" + fmt.Sprint(i)
			f, err := os.Create(path)
			if err != nil {
				log.Fatalf("Could not create file: %s", path)
			}
			w := bufio.NewWriter(f)
			io.CopyN(w, rand.Reader, 1024*1024+1)

		}
	}
	// Create copies
	Copy(root+"test_dir1/test_file1", root+"test_dir1/test_file_copy1")
	Copy(root+"test_dir2/test_file1", root+"test_dir1/test_file_copy2")
	Copy(root+"test_dir3/test_dir32/test_dir321/test_file1", root+"test_dir1/test_file_copy3")
	Copy(root+"test_dir3/test_dir32/test_dir321/test_file1", root+"test_dir3/test_dir32/test_dir322/test_file_copy1")
}
