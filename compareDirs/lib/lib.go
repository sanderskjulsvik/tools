package comparedirs

import (
	"fmt"

	producerconsumer "github.com/sander-skjulsvik/tools/dupes/lib/producerConsumer"
)

// CompareDirs compares two directories
func CompareDirs(dir1, dir2 string, mode) {
	fmt.Printf("Comparing directories: %s and %s\n", dir1, dir2)
	dir1dupes := producerconsumer.Run(dir1)
	dir2dupes := producerconsumer.Run(dir2)
	switch mode {
	case "both":
		both := onlyInBoth(dir1dupes, dir2dupes)
		both.Present(true)

	// Show dupes that is present in both directories

	// Show dupes that is only present in dir1

	// Show dupes that is only present in dir2
}


func onlyInBoth(dupes1 *common.Dupes, dupes2 *common.Dupes) *common.Dupes {

}

func onlyInFirst(dupes1 *common.Dupes, dupes2 *common.Dupes) *common.Dupes {

}

func onlyDestinct(dupes1 *common.Dupes, dupes2 *common.Dupes) *common.Dupes {

}
