package dedupe

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"
)

func copy(src string, dst string) {
	fin, err := os.Open(src)
	if err != nil {
		log.Fatalf("Could not open file: %s, err: %s", src, err.Error())
	}
	fout, err := os.Create(dst)
	if err != nil {
		log.Fatalf("Could not create file: %s, err: %s", src, err.Error())
	}
	_, err = io.Copy(fout, fin)
	if err != nil {
		log.Fatalf("Could not copy to file: %s, err: %s", src, err.Error())
	}

}

func BuildFolderTree(root string) {
	os.MkdirAll(root, 0740)
	// Setup dirs
	dirs := []string{
		"test1/",
		"test2/test1/",
		"test2/test2/",
	}
	for _, dir := range dirs {
		os.MkdirAll(root+dir, 0740)
	}
	// Build unique files
	for _, dir := range dirs {
		for i := 1; i <= 3; i++ {
			path := root + dir + "test" + fmt.Sprint(i)
			f, err := os.Create(path)
			if err != nil {
				log.Fatalf("Could not create file:%s", path)
			}
			w := bufio.NewWriter(f)
			io.CopyN(w, rand.Reader, 10*1024)

		}
	}
	// Create copies
	// - Copy: same dir different name
	copy(root+"test1/test1", root+"test1/test11")
	copy(root+"test2/test2", root+"test1/test11")
	// - Copy: different dir
	copy(root+"test1/test2", root+"test2/test111")
	copy(root+"test2/test2", root+"test2/test112")
}
