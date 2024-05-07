package main

import (
	"flag"
	"fmt"
	"os"

	comparedirs "github.com/sander-skjulsvik/tools/compareDirs/lib"
	"github.com/sander-skjulsvik/tools/dupes/lib/common"
)

func main() {
	// Define command-line flags
	dir1 := flag.String("dir1", "", "First directory path")
	dir2 := flag.String("dir2", "", "Second directory path")
	mode := flag.String("mode", "All", "Mode to run in")

	// Parse command-line flags
	flag.Parse()

	// Check if directory paths are provided
	if *dir1 == "" || *dir2 == "" {
		fmt.Println("Please provide directory paths using -dir1 and -dir2 flags")
		os.Exit(1)
	}

	var newD *common.Dupes
	switch *mode {
	// Show dupes that is present in both directories
	case "OnlyInboth":
		newD = comparedirs.OnlyInboth(*dir1, *dir2)
	// Show dupes that is only present in first
	case "onlyInFirst":
		newD = comparedirs.OnlyInFirst(*dir1, *dir2)
	case "all":
		newD = comparedirs.All(*dir1, *dir2)
	default:
		panic(fmt.Errorf("unknown mode: %s", mode))
	}
	newD.Present(false)
}
