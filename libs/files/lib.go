package files

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/sander-skjulsvik/tools/dupes/lib/common"
)

func GetNumberOfFiles(path string) (int, error) {
	n := 0
	err := filepath.Walk(
		path,
		func(path string, info fs.FileInfo, err error) error {
			isFile := common.IsFile(info)
			if !isFile {
				return nil
			}
			n++
			return nil
		},
	)
	if err != nil {
		return 0, fmt.Errorf("Unable find number of files in dir: %w", err)
	}
	return n, nil
}
