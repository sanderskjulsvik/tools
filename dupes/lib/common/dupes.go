package common

import "fmt"

// Dupes is a struct for holding duplicate files
type Dupes struct {
	D map[string]*Dupe
	// ProgressBar ProgressBar
}

// Dupe is a struct for holding duplicates of a file
type Dupe struct {
	Hash  string
	Paths []string
}

// NewDupes creates a new dupes struct
func NewDupes() Dupes {
	dupes := Dupes{}
	dupes.D = make(map[string]*Dupe)
	return dupes
}

// Append appends a file to the dupes struct
func (dupes *Dupes) Append(path string) (*Dupes, error) {
	hash, err := HashFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable append file: %w", err)
	}
	dupes.AppendHashedFile(path, hash)
	return dupes, nil
}

// AppendHashedFile appends a file to the dupes struct
func (dupes *Dupes) AppendHashedFile(path string, hash string) {
	if _, ok := dupes.D[hash]; !ok {
		// If file hash has not been found yet
		dupes.D[hash] = &Dupe{
			Hash:  hash,
			Paths: []string{path},
		}
	} else {
		dupes.D[hash].Paths = append(dupes.D[hash].Paths, path)
	}
}

func (dupes *Dupes) print() {
	for _, dupe := range dupes.D {
		fmt.Printf("sha256:%s \n", dupe.Hash)
		for _, path := range dupe.Paths {
			fmt.Printf("    %s \n", path)
		}
		fmt.Println("")
	}
}

// PrintOnlyDupes prints only files that have duplicates
func (dupes *Dupes) PrintOnlyDupes() {
	for _, dupe := range dupes.D {
		if len(dupe.Paths) > 1 {
			fmt.Printf("sha256:%s \n", dupe.Hash)
			for _, path := range dupe.Paths {
				fmt.Printf("    %s \n", path)
			}
			fmt.Println("")
		}
	}
}

// Present prints all found files with hash and paths
func (dupes *Dupes) Present(onlyDupes bool) {
	if onlyDupes {
		dupes.PrintOnlyDupes()
	} else {
		dupes.print()
	}
}
