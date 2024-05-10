package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sander-skjulsvik/tools/dupes/lib/common"
	comparedirs "github.com/sander-skjulsvik/tools/dupesCompareDirs/lib"
)

func main() {
	// Define command-line flags
	mode := flag.String("mode", "all", "Mode to run in, modes: OnlyInboth, onlyInFirst, all")
	outputJson := flag.Bool("json", false, "If set to true Output as json")
	// Parse command-line flags
	flag.Parse()

	log.Printf("Number in argvs: %d\n", len(os.Args))
	if len(os.Args) < 3 {
		panic(fmt.Errorf("please provide to folders"))
	}
	dir1 := os.Args[len(os.Args)-2]
	dir2 := os.Args[len(os.Args)-1]

	// Check if directory paths are provided
	if dir1 == "" || dir2 == "" {
		fmt.Println("Please provide directory paths to compare")
		os.Exit(1)
	}
	log.Printf("Comparing directories: %s and %s\n", dir1, dir2)

	var newD *common.Dupes
	switch *mode {
	// Show dupes that is present in both directories
	case "OnlyInboth":
		newD = comparedirs.OnlyInAll(dir1, dir2)
	// Show dupes that is only present in first
	case "onlyInFirst":
		newD = comparedirs.OnlyInFirst(dir1, dir2)
		log.Println("Only in first")
		log.Printf("Number of dupes: %d\n", len(newD.D))
	case "all":
		newD = comparedirs.All(dir1, dir2)
	default:
		panic(fmt.Errorf("unknown mode: %s, supported modes: OnlyInboth, onlyInFirst, all ", *mode))
	}

	if *outputJson {
		fmt.Println(string(newD.GetJSON()))
	} else {
		newD.Present(false)
	}
}
