package main

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
)

type Dupe struct {
	hash  string
	Paths []*string
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
		return "", nil
	}

	hash, err := hashFile(path)
	if err != nil {
		return "", fmt.Errorf("Could not process file: %s", err.Error())
	}
	return hash, nil
}

func main() {
	flag.Parse()
	src := flag.Arg(0)
	dupes := map[string]*Dupe{}

	err := filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		hash, err := visit(path, info, err)
		if err != nil {
			return err
		}
		if hash == "" {
			return nil
		}
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

		return nil
	})
	if err != nil {
		log.Fatalf("Failed to walk src: %s, with err: %s", src, err.Error())
	}
	for _, dupe := range dupes {
		fmt.Printf("sha256:%s \n", dupe.hash)
		for _, path := range dupe.Paths {
			fmt.Printf("    %s \n", *path)
		}
		fmt.Println("")
	}
}
