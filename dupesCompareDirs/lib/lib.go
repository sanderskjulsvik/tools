package dupescomparedirs

import (
	"log"
	"sync"

	"github.com/sander-skjulsvik/tools/dupes/lib/common"
	"github.com/sander-skjulsvik/tools/dupes/lib/singleThread"
	"github.com/sander-skjulsvik/tools/libs/files"
	"github.com/sander-skjulsvik/tools/libs/progressbar"
)

// OnlyInboth returns dupes that is present in both directories
func OnlyInAll(paths ...string) *common.Dupes {
	ds := runDupes(paths...)
	first := ds[0]

	for _, d := range ds {
		first = first.OnlyInBoth(d)
	}

	return first
}

// OnlyInFirst returns dupes that is only present in first directory
func OnlyInFirst(paths ...string) *common.Dupes {
	ds := runDupes(paths...)
	first := ds[0]
	for _, d := range ds {
		first = first.OnlyInSelf(d)
	}
	return first
}

// All returns all dupes in both directories
func All(paths ...string) *common.Dupes {
	dupes := common.NewDupes()
	for _, dupe := range runDupes(paths...) {
		dupes.AppendDupes(dupe)
	}
	return &dupes
}

func runDupes(paths ...string) []*common.Dupes {
	wg := sync.WaitGroup{}
	wg.Add(len(paths))
	dupesCollection := make([]*common.Dupes, len(paths))

	progressBars := progressbar.NewUiProgressBars()
	progressBars.Start()

	for ind, path := range paths {
		go func() {
			log.Printf("Running dupes on: %s", path)
			n, _ := files.GetNumbeSizeOfDirMb(path)
			bar := progressBars.AddBar(path, n)
			dupesCollection[ind] = singleThread.RunWithProgressBar(path, bar)
			wg.Done()
		}()
	}
	wg.Wait()
	// progressBars.Stop()
	// time.Sleep(10 * time.Second)

	return dupesCollection
}
