package dupes

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type Dupe struct {
	hash  string
	Paths []*string
}

type File struct {
	path string
}

func hashString(b []byte) string {
	return hex.EncodeToString(b)
}

func hashFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("Failed to open: %s: %w", path, err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", fmt.Errorf("Failed to hash: %s: %w", path, err)
	}

	return hashString(h.Sum(nil)), nil
}

func visit(path string, f os.FileInfo, err error) (string, error) {
	if !f.Mode().IsRegular() {
		return "nil", nil
	}
	return path, nil
}

func walker(root string, files chan<- string) {
	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		_, err = visit(path, info, err)
		if err != nil {
			return err
		}
		files <- path
		return nil
	})
}

func processer(dupes sync.Map, files <-chan string) {

}

func main() {
	flag.Parse()
	src := flag.Arg(0)
	var files chan string
	walker(src, files)
	processer(files)

	if d, ok := dupes[hash]; !ok {
		// If file hash has not been found yet
		dupe := Dupe{
			hash:  hash,
			Paths: []*string{&path},
		}

		dupes[hash] = &dupe
	} else {
		d.Paths = append(d.Paths, &path)
	}
	if err != nil {
		log.Fatalf("Failed to walk src: %s, with err: %s", src, err.Error())
	}
	dupes := map[string]*Dupe{}

	for _, dupe := range dupes {
		fmt.Printf("sha256:%s \n", dupe.hash)
		for _, path := range dupe.Paths {
			fmt.Printf("    %s \n", *path)
		}
		fmt.Println("")
	}
}
