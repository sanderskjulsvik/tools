package common

import (
	"encoding/json"
	"fmt"
)

// Dupes is a struct for holding duplicate files
type Dupes struct {
	D map[string]*Dupe `json:"dupes"`
	// ProgressBar ProgressBar
}

// Dupe is a struct for holding duplicates of a file
type Dupe struct {
	Hash  string   `json:"hash"`
	Paths []string `json:"paths"`
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

func (dupes *Dupes) appendDupe(dupe *Dupe) {
	for _, path := range dupe.Paths {
		dupes.AppendHashedFile(path, dupe.Hash)
	}
}

func (dupes *Dupes) AppendDupes(other *Dupes) {
	for _, dupe := range other.D {
		dupes.appendDupe(dupe)
	}
}

// GetOnlyDupes returns a new dupes struct with only the files that have duplicates
func (dupes *Dupes) GetOnlyDupes() *Dupes {
	onlyDupes := NewDupes()
	for _, dupe := range dupes.D {
		if len(dupe.Paths) > 1 {
			onlyDupes.appendDupe(dupe)
		}
	}
	return &onlyDupes
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
	onlyDupes := dupes.GetOnlyDupes()
	for _, dupe := range onlyDupes.D {
		fmt.Printf("sha256:%s \n", dupe.Hash)
		for _, path := range dupe.Paths {
			fmt.Printf("    %s \n", path)
		}
		fmt.Println("")
	}
}

// GetJSON returns the dupes struct as a json byte array
func (dupes *Dupes) GetJSON() []byte {
	b, err := json.Marshal(dupes)
	if err != nil {
		panic(fmt.Errorf("unable to convert dupes to json, this is a bug: %w", err))
	}
	return b
}

// GetJSONOnlyDupes returns the dupes struct with only duplicates as a json byte array
func (dupes *Dupes) GetJSONOnlyDupes() []byte {
	onlyDupes := dupes.GetOnlyDupes()
	b, _ := json.Marshal(onlyDupes)
	return b
}

// Present prints all found files with hash and paths
func (dupes *Dupes) Present(onlyDupes bool) {
	if onlyDupes {
		dupes.PrintOnlyDupes()
	} else {
		dupes.print()
	}
}

// OnlyInBoth returns a new dupes struct with only the files that are in both dupes structs
func (dupes *Dupes) OnlyInBoth(other *Dupes) *Dupes {
	both := NewDupes()
	for hash, dupe := range dupes.D {
		if otherDupe, ok := other.D[hash]; ok {
			both.appendDupe(dupe)
			both.appendDupe(otherDupe)
		}
	}
	return &both
}

// OnlyInSelf returns a new dupes struct with only the files that are not in the other dupes struct
func (dupes *Dupes) OnlyInSelf(other *Dupes) *Dupes {
	onlyInSelf := NewDupes()
	for hash, dupe := range dupes.D {
		if _, ok := other.D[hash]; !ok {
			onlyInSelf.appendDupe(dupe)
		}
	}
	return &onlyInSelf
}

// HasSameFiles checks if two dupes structs have the same files,
// i.e. the same hashes, does not care about paths
func (dupes *Dupes) HasSameFiles(other *Dupes) bool {
	if len(dupes.D) != len(other.D) {
		return false
	}
	for hash := range dupes.D {
		if _, ok := other.D[hash]; !ok {
			return false
		}
	}
	for hash := range other.D {
		if _, ok := dupes.D[hash]; !ok {
			return false
		}
	}
	return true
}
