package main

import (
	"fmt"

	comparedirs "github.com/sander-skjulsvik/tools/dupesCompareDirs/lib"
	"github.com/sander-skjulsvik/tools/libs/progressbar"
)

func main() {
	input := comparedirs.HandleCliInput()
	// Progress bar
	pbs := progressbar.NewUiPCollection()
	dupes := comparedirs.OnlyInFirst(pbs, input.Dir1, input.Dir2)

	if input.OutputJson {
		fmt.Println(string(dupes.GetJSON()))
	} else {
		dupes.Present(false)
	}
}
