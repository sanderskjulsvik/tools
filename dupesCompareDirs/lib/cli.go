package dupescomparedirs

import (
	"flag"
	"fmt"
	"os"

	"github.com/sander-skjulsvik/tools/libs/progressbar"
)

func RunComparison(comparisonFunc ComparisonFunc) {
	outputJson := flag.Bool("json", false, "If set to true Output as json")
	dir1 := flag.String("dir1", "", "Path to 1st dir")
	dir2 := flag.String("dir2", "", "Path to 2nd dir")
	flag.Parse()

	// Check if directory paths are provided
	if *dir1 == "" || *dir2 == "" {
		fmt.Println("Please provide directory paths to compare")
		os.Exit(1)
	}

	// Progress bar
	pbs := progressbar.NewUiPCollection()
	dupes := comparisonFunc(pbs, *dir1, *dir2)

	if *outputJson {
		fmt.Println(string(dupes.GetJSON()))
	} else {
		dupes.Present(false)
	}
}
