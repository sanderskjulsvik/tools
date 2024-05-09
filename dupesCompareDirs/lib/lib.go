package dupescomparedirs

import (
	"log"
	"sync"

	"github.com/sander-skjulsvik/tools/dupes/lib/common"
	"github.com/sander-skjulsvik/tools/dupes/lib/singleThread"
	"github.com/sander-skjulsvik/tools/libs/files"
)

// OnlyInboth returns dupes that is present in both directories
func OnlyInboth(path1, path2 string, parallel bool) *common.Dupes {
	ds := runDupes(parallel, []string{path1, path2})
	d1 := ds[0]
	d2 := ds[1]

	return d1.OnlyInBoth(d2)
}

// OnlyInFirst returns dupes that is only present in first directory
func OnlyInFirst(path1, path2 string, parallel bool) *common.Dupes {
	ds := runDupes(parallel, []string{path1, path2})
	d1 := ds[0]
	d2 := ds[1]

	log.Printf("Number of dupes in first directory: %d\n", len(d1.D))
	log.Printf("Number of dupes in second directory: %d\n", len(d1.D))

	return d1.OnlyInSelf(d2)
}

// All returns all dupes in both directories
func All(parallel bool, paths []string) *common.Dupes {
	dupes := common.NewDupes()
	for _, dupe := range runDupes(parallel, paths) {
		dupes.AppendDupes(dupe)
	}
	return &dupes
}

func runDupes(parralel bool, paths []string) []*common.Dupes {
	// var runFunc common.RunWithProgressBar
	// switch parralel {
	// // case true:
	// // 	runFunc = producerconsumer.Run
	// default:
	// 	runFunc = singleThread.RunWithProgressBar
	// }
	wg := sync.WaitGroup{}
	wg.Add(len(paths))
	dupesCollection := make([]*common.Dupes, len(paths))

	progressBars := common.NewUiProgressBars()
	progressBars.Start()

	for ind, path := range paths {
		go func() {
			log.Printf("Running dupes on: %s", path)
			n, _ := files.GetNumberOfFiles(path)
			bar := progressBars.AddBar(n)
			dupesCollection[ind] = singleThread.RunWithProgressBar(path, bar)
			wg.Done()
		}()
	}
	wg.Wait()

	return dupesCollection
}
