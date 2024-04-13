package producerConsumer

import (
	"io/fs"
	"path/filepath"
	"sync"

	"github.com/sander-skjulsvik/tools/dupes/lib/common"
)

func getFiles(root string, filePaths chan<- string) {
	defer close(filePaths)

	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		// If regular file, send it to the channel
		if info.Mode().IsRegular() {
			filePaths <- path
		}
		return nil
	})
}

func appendFileTreadSafe(dupes *common.Dupes, path string, lock *sync.Mutex) {
	hash, err := common.HashFile(path)
	if err != nil {
		panic(err)
	}
	lock.Lock()
	defer lock.Unlock()
	dupes.AppendHashedFile(path, hash)
}

func ProcessFiles(filePaths <-chan string) *common.Dupes {
	dupes := common.Dupes.New(common.Dupes{})
	wg := sync.WaitGroup{}
	dupesWl := sync.Mutex{}
	for filePath := range filePaths {
		wg.Add(1)
		go func() {
			appendFileTreadSafe(&dupes, filePath, &dupesWl)
			wg.Done()
		}()
	}
	wg.Wait()
	return &dupes
}

func ProcessFilesNCunsumers(filePaths <-chan string, numberOfConsumers int) common.Dupes {
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
	return dupes
}

func presenter(dupes common.Dupes) {
	dupes.Print()
}

func Run(path string) *common.Dupes {
	filePaths := make(chan string)
	go getFiles(path, filePaths)
	dupes := ProcessFiles(filePaths)
	presenter(*dupes)
	// storer(files)
	return dupes
}
