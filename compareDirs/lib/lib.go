package comparedirs

import (
	"fmt"

	"github.com/sander-skjulsvik/tools/dupes/lib/common"
	producerconsumer "github.com/sander-skjulsvik/tools/dupes/lib/producerConsumer"
)

// CompareDirs compares two directories
func CompareDirs(dir1, dir2 string, mode string) (*common.Dupes, error) {
	var newD *common.Dupes
	switch mode {
	// Show dupes that is present in both directories
	case "OnlyInboth":
		newD = OnlyInboth(dir1, dir2)
	// Show dupes that is only present in first
	case "onlyInFirst":
		newD = OnlyInFirst(dir1, dir2)
	case "all":
		newD = All(dir1, dir2)
	default:
		return nil, fmt.Errorf("unknown mode: %s", mode)
	}
	return newD, nil
}

// OnlyInboth returns dupes that is present in both directories
func OnlyInboth(path1, path2 string) *common.Dupes {
	d1 := producerconsumer.Run(path1)
	d2 := producerconsumer.Run(path2)
	return d1.OnlyInBoth(d2)
}

// OnlyInFirst returns dupes that is only present in first directory
func OnlyInFirst(path1, path2 string) *common.Dupes {
	d1 := producerconsumer.Run(path1)
	d2 := producerconsumer.Run(path2)
	return d1.OnlyInSelf(d2)
}

// All returns all dupes in both directories
func All(path1, path2 string) *common.Dupes {
	d1 := producerconsumer.Run(path1)
	d2 := producerconsumer.Run(path2)
	d1.AppendDupes(d2)
	return d1
}
