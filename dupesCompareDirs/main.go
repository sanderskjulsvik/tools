package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sander-skjulsvik/tools/dupes/lib/common"
	comparedirs "github.com/sander-skjulsvik/tools/dupesCompareDirs/lib"
)

func main() {
	// Define command-line flags
	mode := flag.String("mode", "all", "Mode to run in")
	dir1 := flag.String("dir1", "", "Path to the first directory")
	dir2 := flag.String("dir2", "", "Path to the second directory")

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
		panic(fmt.Errorf("unknown mode: %s", *mode))
	}
	newD.Present(false)
}
