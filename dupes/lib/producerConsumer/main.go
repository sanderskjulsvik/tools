package producerConsumer

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/sander-skjulsvik/tools/dupes/lib/common"
)

func visit(path string, f os.FileInfo, err error) (string, error) {
	if !f.Mode().IsRegular() {
		return "nil", nil
	}
	return path, nil
}

func directoryWalker(root string, filePaths chan<- string) {
	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		_, err = visit(path, info, err)
		if err != nil {
			return err
		}
		filePaths <- path
		return nil
	})
}

func processor(filePaths <-chan string) (common.Dupes, error) {
	dupes := common.Dupes.New(common.Dupes{})

	for filePath := range filePaths {
		dupes.Append(filePath)
	}

	return dupes, nil
}

func presenter(dupes common.Dupes) {
	dupes.Print()
}
func Run() {
	flag.Parse()
	src := flag.Arg(0)
	var files chan common.File
	directoryWalker(src, files)
	processer(files)
	storer(files)

func Run(src string) {

	// Find the files
	var filePaths chan string
	directoryWalker(src, filePaths)

	// Process files
	dupes, err := processor(filePaths)
	if err != nil {
		fmt.Printf("Failed to process directory: %s\n", err)
		panic("Stopping because off processing dir error")
	}

	// Present the result (?)
	presenter(dupes)

}
