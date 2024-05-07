package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define command-line flags
	dir1 := flag.String("dir1", "", "First directory path")
	dir2 := flag.String("dir2", "", "Second directory path")

	// Parse command-line flags
	flag.Parse()

	// Check if directory paths are provided
	if *dir1 == "" || *dir2 == "" {
		fmt.Println("Please provide directory paths using -dir1 and -dir2 flags")
		os.Exit(1)
	}

	CompareDirs(*dir1, *dir2)
}
