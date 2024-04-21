package producerConsumer

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"sync"

	"github.com/sander-skjulsvik/tools/dupes/lib/common"
)

// Works like a generator, yelding all regular files
func getFiles(root string, filePaths chan<- string, doneWg *sync.WaitGroup) {
	defer func() {
		fmt.Println("I am in defer get files")
		doneWg.Done()
		doneWg.Wait()
		close(filePaths)
	}()

	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Errorf("Failed to walk file %s, %v", path, err)
			return nil
		}
		// If regular file, send it to the channel
		if info.Mode().IsRegular() {
			fmt.Printf("Walked to %s\n", path)
			filePaths <- path
		}
		return nil
	})

}

func appendFileTreadSafe(dupes *common.Dupes, path string, lock *sync.Mutex) {
	hash, err := common.HashFile(path)
	if err != nil {
		fmt.Printf("Could not hash: %s, err: %s\n", path, err.Error())
		return
	}
	lock.Lock()
	defer lock.Unlock()
	dupes.AppendHashedFile(path, hash)
	// dupes.ProgressBar.Add1()
}

func ProcessFiles(filePaths <-chan string, doneWg *sync.WaitGroup) *common.Dupes {
	dupes := common.Dupes.New(common.Dupes{})
	wg := sync.WaitGroup{}
	dupesWl := sync.Mutex{}
	// if chans.IsClosed(filePaths) {
	// 	log.Fatalln("Chan closed before managed to access it 1")
	// }
	for filePath := range filePaths {
		wg.Add(1)
		go func(fp string) {
			appendFileTreadSafe(&dupes, fp, &dupesWl)
			wg.Done()
		}(filePath)
	}
	wg.Wait()
	doneWg.Done()
	return &dupes
}

func ProcessFilesNCunsumers(filePaths <-chan string, numberOfConsumers int, doneWg *sync.WaitGroup) *common.Dupes {
	dupes := common.Dupes.New(common.Dupes{})
	wg := sync.WaitGroup{}
	dupesWl := sync.Mutex{}
	wg.Add(numberOfConsumers)
	for i := 0; i < numberOfConsumers; i++ {
		go func() {
			for filePath := range filePaths {
				appendFileTreadSafe(&dupes, filePath, &dupesWl)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	doneWg.Done()
	doneWg.Wait()
	return &dupes
}

func Run(path string, presentOnlyDupes bool) *common.Dupes {
	filePaths := make(chan string, 10)
	doneGroup := sync.WaitGroup{}
	doneGroup.Add(2)
	go getFiles(path, filePaths, &doneGroup)
	dupes := ProcessFiles(filePaths, &doneGroup)
	dupes.Present(presentOnlyDupes)
	// storer(files)
	return dupes
}
