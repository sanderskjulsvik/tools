package dupescomparedirs

import (
	"log"

	"github.com/sander-skjulsvik/tools/dupes/lib/common"
	producerconsumer "github.com/sander-skjulsvik/tools/dupes/lib/producerConsumer"
	"github.com/sander-skjulsvik/tools/dupes/lib/singleThread"
)

// OnlyInboth returns dupes that is present in both directories
func OnlyInboth(path1, path2 string, parallel bool) *common.Dupes {
	var d1, d2 *common.Dupes
	if parallel {
		d1 = producerconsumer.Run(path1)
		d2 = producerconsumer.Run(path2)
	} else {
		d1 = singleThread.Run(path1)
		d2 = singleThread.Run(path2)
	}
	return d1.OnlyInBoth(d2)
}

// OnlyInFirst returns dupes that is only present in first directory
func OnlyInFirst(path1, path2 string, parallel bool) *common.Dupes {
	var d1, d2 *common.Dupes
	if parallel {
		d1 = producerconsumer.Run(path1)
		d2 = producerconsumer.Run(path2)
	} else {
		d1 = singleThread.Run(path1)
		d2 = singleThread.Run(path2)
	}
	log.Printf("Number of dupes in first directory: %d\n", len(d1.D))
	log.Printf("Number of dupes in second directory: %d\n", len(d1.D))

	return d1.OnlyInSelf(d2)
}

// All returns all dupes in both directories
func All(parallel bool, paths ...string) *common.Dupes {
	var dupes common.Dupes
	if parallel {
		for _, path := range paths {
			dupes.AppendDupes(producerconsumer.Run(path))
		}
	} else {
		for _, path := range paths {
			dupes.AppendDupes(singleThread.Run(path))
		}
	}
	return &dupes
}
