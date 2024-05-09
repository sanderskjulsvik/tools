package singleThread

import (
	"io/fs"
	"log"
	"path/filepath"

	"github.com/sander-skjulsvik/tools/dupes/lib/common"
)

func Run(src string) *common.Dupes {
	dupes := &common.Dupes{
		D: map[string]*common.Dupe{},
		// ProgressBar: common.NewSchollzProgressbar(),
	}

	err := filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		isFile := common.IsFile(info)
		if !isFile {
			return nil
		}

		dupes, err = dupes.Append(path)
		if err != nil {
			return nil
		}

		return nil
	})
	if err != nil {
		log.Fatalf("Failed to walk src: %s, with err: %s", src, err.Error())
	}
	return dupes
}

func RunWithProgressBar(src string, bar *common.ProgressBar) *common.Dupes {
	dupes := &common.Dupes{
		D: map[string]*common.Dupe{},
		// ProgressBar: common.NewSchollzProgressbar(),
	}

	err := filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		isFile := common.IsFile(info)
		if !isFile {
			return nil
		}

		dupes, err = dupes.Append(path)
		bar.
		if err != nil {
			return nil
		}

		return nil
	})
	if err != nil {
		log.Fatalf("Failed to walk src: %s, with err: %s", src, err.Error())
	}
	return dupes
}
