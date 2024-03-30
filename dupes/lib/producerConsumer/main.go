package producerConsumer

import (
	"flag"
	"io/fs"
	"path/filepath"

	"github.com/sander-skjulsvik/tools/dupes/lib/common"
)

func getFiles(root string, filePaths chan<- string) {
	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		// If regular file, send it to the channel
		if info.Mode().IsRegular() {
			filePaths <- path
		}
		return nil
	})
}

// Consumer
func ProcessFiles(filePaths <-chan string) common.Dupes {
	dupes := common.Dupes.New(common.Dupes{})

	for filePath := range filePaths {
		dupes.Append(filePath)
	}

	return dupes
}

func ProcessFilesNCunsumers(filePaths <-chan string, numberOfConsumers int) common.Dupes {
	dupes := common.Dupes.New(common.Dupes{})
	for i := 0; i < numberOfConsumers; i++ {
		go func() {
			for filePath := range filePaths {
				dupes.Append(filePath)
			}
		}()
	}

	return dupes
}

func presenter(dupes common.Dupes) {
	dupes.Print()
}
func Run() {
	flag.Parse()
	src := flag.Arg(0)
	filePaths := make(chan string)
	go getFiles(src, filePaths)
	dupes := ProcessFiles(filePaths)
	presenter(dupes)
	// storer(files)
}

// func Run(src string) {

// 	// Find the files
// 	var filePaths chan string
// 	directoryWalker(src, filePaths)

// 	// Process files
// 	dupes, err := processor(filePaths)
// 	if err != nil {
// 		fmt.Printf("Failed to process directory: %s\n", err)
// 		panic("Stopping because off processing dir error")
// 	}

// 	// Present the result (?)
// 	presenter(dupes)

// }
