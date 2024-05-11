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
func OnlyInAll(progressBars progressbar.ProgressBarCollection, paths ...string) *common.Dupes {
	ds := runDupes(progressBars, paths...)
	first := ds[0]

	for _, d := range ds {
		first = first.OnlyInBoth(d)
	}

	return first
}

// OnlyInFirst returns dupes that is only present in first directory
func OnlyInFirst(progressBarCollection progressbar.ProgressBarCollection, paths ...string) *common.Dupes {
	ds := runDupes(progressBarCollection, paths...)
	first := ds[0]
	for _, d := range ds {
		first = first.OnlyInSelf(d)
	}
	return first
}

// All returns all dupes in both directories
func All(progressBarCollection progressbar.ProgressBarCollection, paths ...string) *common.Dupes {
	dupes := common.NewDupes()
	for _, dupe := range runDupes(progressBarCollection, paths...) {
		dupes.AppendDupes(dupe)
	}
	return &dupes
}

func runDupes(progressBarCollection progressbar.ProgressBarCollection, paths ...string) []*common.Dupes {
	wg := sync.WaitGroup{}
	wg.Add(len(paths))
	dupesCollection := make([]*common.Dupes, len(paths))

	progressBarCollection.Start()

	for ind, path := range paths {
		go func() {
			log.Printf("Running dupes on: %s", path)
			n, _ := files.GetNumbeSizeOfDirMb(path)
			bar := progressBarCollection.AddBar(path, n)
			dupesCollection[ind] = singleThread.RunWithProgressBar(path, bar)
			wg.Done()
		}()
	}
	wg.Wait()
	progressBarCollection.Stop()
	// time.Sleep(10 * time.Second)

	return dupesCollection
}
