package producerconsumer

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"sync"

	"github.com/sander-skjulsvik/tools/dupes/lib/common"
)

// Works like a generator, yelding all regular files
func getFiles(root string, filePaths chan<- string) {
	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Errorf("Failed to walk file %s, %v", path, err)
			return nil
		}
		// If regular file, send it to the channel
		if info.Mode().IsRegular() {
			filePaths <- path
		}
		return nil
	})
	close(filePaths)
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

func ProcessFiles(filePaths <-chan string) *common.Dupes {
	dupes := common.NewDupes()
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
	return &dupes
}

func ProcessFilesNCunsumers(filePaths <-chan string, numberOfConsumers int, doneWg *sync.WaitGroup) *common.Dupes {
	dupes := common.NewDupes()
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
	return &dupes
}

func Run(path string) *common.Dupes {
	filePaths := make(chan string)
	go getFiles(path, filePaths)
	// sleep 10 seconds
	dupes := ProcessFiles(filePaths)
	// storer(files)
	return dupes
}
