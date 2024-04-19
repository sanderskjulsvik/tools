package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	producerConsumer "github.com/sander-skjulsvik/tools/dupes/lib/producerConsumer"
	singleThread "github.com/sander-skjulsvik/tools/dupes/lib/singleThread"
)

func main() {

	method := strings.ToLower(*flag.String("method", "single", "Method (single or producerConsumer)"))
	path := *flag.String("path", ".", "File path")
	presentOnlyDupes := *flag.Bool("onlyDupes", true, "Only present dupes")

	// Parse the command-line arguments
	flag.Parse()

	// Check if the method is one of the allowed values
	if method != "single" && method != "producerConsumer" {
		fmt.Println("Invalid method. Allowed values are 'single' and 'producerConsumer'.")
		os.Exit(1)
	}

	// At this point, you have valid values for method and path
	fmt.Printf("Method: %s\n", method)
	fmt.Printf("Path: %s\n", path)
	fmt.Printf("PresentOnlyDupes: %t\n", presentOnlyDupes)

	Run(path, method, presentOnlyDupes)
}

func Run(path, method string, presentOnlyDupes bool) {
	switch {
	case method == "single":
		singleThread.Run(path, presentOnlyDupes)
	case method == "producerconsumer":
		producerConsumer.Run(path, presentOnlyDupes)
	}
}
