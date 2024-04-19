package common

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// Function to run to find dupes.
type Run func(string, bool) *Dupes

type Dupe struct {
	Hash  string
	Paths []*string
}

type File struct {
	Path string
	Hash string
}

type Dupes struct {
	D map[string]*Dupe
}

func (dupes Dupes) New() Dupes {
	dupes = Dupes{}
	dupes.D = make(map[string]*Dupe)
	return dupes
}

func (dupes *Dupes) Append(path string) (*Dupes, error) {
	hash, err := HashFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable append file: %w", err)
	}
	dupes.AppendHashedFile(path, hash)
	return dupes, nil
}

// AppendHashedFile
func (dupes *Dupes) AppendHashedFile(path string, hash string) {
	if _, ok := dupes.D[hash]; !ok {
		// If file hash has not been found yet
		dupes.D[hash] = &Dupe{
			Hash:  hash,
			Paths: []*string{&path},
		}
	} else {
		dupes.D[hash].Paths = append(dupes.D[hash].Paths, &path)
	}
}

func (dupes *Dupes) Print() {
	for _, dupe := range dupes.D {
		fmt.Printf("sha256:%s \n", dupe.Hash)
		for _, path := range dupe.Paths {
			fmt.Printf("    %s \n", *path)
		}
		fmt.Println("")
	}
}

func (dupes *Dupes) PrintOnlyDupes() {
	for _, dupe := range dupes.D {
		if len(dupe.Paths) > 1 {
			fmt.Printf("sha256:%s \n", dupe.Hash)
			for _, path := range dupe.Paths {
				fmt.Printf("    %s \n", *path)
			}
			fmt.Println("")
		}
	}
}

func (dupes *Dupes) Present(onlyDupes bool) {
	if onlyDupes {
		dupes.PrintOnlyDupes()
	} else {
		dupes.Print()
	}
}

func HashString(b []byte) string {
	return hex.EncodeToString(b)
}

func IsFile(f os.FileInfo) bool {
	return f.Mode().IsRegular()
}

func HashFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open: %s: %w", path, err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", fmt.Errorf("failed to hash: %s: %w", path, err)
	}

	return HashString(h.Sum(nil)), nil
}
