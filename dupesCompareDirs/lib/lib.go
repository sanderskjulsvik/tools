package dupescomparedirs

import (
	"github.com/sander-skjulsvik/tools/dupes/lib/common"
	producerconsumer "github.com/sander-skjulsvik/tools/dupes/lib/producerConsumer"
)

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
