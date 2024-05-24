package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sander-skjulsvik/tools/dupes/lib/common"
	comparedirs "github.com/sander-skjulsvik/tools/dupesCompareDirs/lib"
	"github.com/sander-skjulsvik/tools/libs/progressbar"
)

func main() {
	// Define command-line flags
	mode := flag.String("mode", "all", "Mode to run in, modes: OnlyInboth, onlyInFirst, all")

	genericCliInput := comparedirs.HandleCliInput()

	log.Printf("Comparing directories: %s and %s\n", genericCliInput.Dir1, genericCliInput.Dir2)

	// Progress bar
	pbs := progressbar.NewUiPCollection()

	var newD *common.Dupes
	switch *mode {
	// Show dupes that is present in both directories
	case "OnlyInboth":
		newD = comparedirs.OnlyInAll(pbs, genericCliInput.Dir1, genericCliInput.Dir2)
	// Show dupes that is only present in first
	case "onlyInFirst":
		newD = comparedirs.OnlyInFirst(pbs, genericCliInput.Dir1, genericCliInput.Dir2)
		log.Println("Only in first")
		log.Printf("Number of dupes: %d\n", len(newD.D))
	case "all":
		newD = comparedirs.All(pbs, genericCliInput.Dir1, genericCliInput.Dir2)
	default:
		panic(fmt.Errorf("unknown mode: %s, supported modes: OnlyInboth, onlyInFirst, all ", *mode))
	}

	if genericCliInput.OutputJson {
		fmt.Println(string(newD.GetJSON()))
	} else {
		newD.Present(false)
	}
}
