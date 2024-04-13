package main

import (
	"flag"

	"github.com/sander-skjulsvik/tools/dedupe/dedupe"
)

func main() {
	var flagPath string
	flag.StringVar(&flagPath, "path", ".", "Directory to find duplicates")
	// var flagThreads int
	// flag.IntVar(&flagThreads, "threads", 1, "Number of threads used to find dupes")
	flag.Parse()
	dedupe.Run(flagPath)
}
