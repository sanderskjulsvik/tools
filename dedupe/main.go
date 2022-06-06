package main

import (
	"os"

	"github.com/sander-skjulsvik/tools/dedupe/dedupe"
)

func main() {
	dedupe.Run(os.Args[1])
}
