package singlethread

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/sander-skjulsvik/tools/dupes/lib/common"
)

func Run(src string) {
	dupes := &common.Dupes{
		D: map[string]*common.Dupe{},
	}

	err := filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		isFile := common.IsFile(info)
		if !isFile {
			return nil
		}
		hash, err := common.HashFile(path)
		if err != nil {
			fmt.Printf("File at %s is not hashable", path)
			return err
		}
		dupes = dupes.Append(path, hash)

		return nil
	})
	if err != nil {
		log.Fatalf("Failed to walk src: %s, with err: %s", src, err.Error())
	}
	dupes.Print()
}
