package dupes

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/sander-skjulsvik/tools/dupes/lib/common"
)

func visit(path string, f os.FileInfo, err error) (string, error) {
	if !f.Mode().IsRegular() {
		return "nil", nil
	}
	return path, nil
}

func directoryWalker(root string, files chan<- common.File) {
	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		_, err = visit(path, info, err)
		if err != nil {
			return err
		}
		files <- path
		return nil
	})
}

func processer(dupes sync.Map, files <-chan string) {

}

func main() {
	flag.Parse()
	src := flag.Arg(0)
	var files chan common.File
	directoryWalker(src, files)
	processer(files)
	storer(files)

	if d, ok := dupes[hash]; !ok {
		// If file hash has not been found yet
		dupe := Dupe{
			hash:  hash,
			Paths: []*string{&path},
		}

		dupes[hash] = &dupe
	} else {
		d.Paths = append(d.Paths, &path)
	}
	if err != nil {
		log.Fatalf("Failed to walk src: %s, with err: %s", src, err.Error())
	}
	dupes := map[string]*Dupe{}

	for _, dupe := range dupes {
		fmt.Printf("sha256:%s \n", dupe.hash)
		for _, path := range dupe.Paths {
			fmt.Printf("    %s \n", *path)
		}
		fmt.Println("")
	}
}
