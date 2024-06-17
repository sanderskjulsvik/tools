package files

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

const BUFFERSIZE int = 65_536

// cp from: https://opensource.com/article/18/6/copying-files-go
func Copy(src, dst string) error {
	source, err := os.Open(src)
	destination, err := os.Open(dst)
	if err != nil {
		return err
	}

	buf := make([]byte, BUFFERSIZE)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return nil
}

func GetNumberOfFiles(path string) (int, error) {
	n := 0
	err := filepath.Walk(
		path,
		func(path string, info fs.FileInfo, err error) error {
			isFile := IsFile(info)
			if !isFile {
				return nil
			}
			n++
			return nil
		},
	)
	if err != nil {
		return 0, fmt.Errorf("unable find number of files in dir: %w", err)
	}
	return n, nil
}

func IsFile(f os.FileInfo) bool {
	if f == nil {
		panic(fmt.Errorf("file info is nil"))
	}
	return f.Mode().IsRegular()
}

func GetSizeOfDirMb(path string) (int, error) {
	var size int64 = 0
	err := filepath.Walk(
		path,
		func(path string, info fs.FileInfo, err error) error {
			if info == nil {
				log.Printf("File info is nil for %s\n", path)
				return nil
			}
			isFile := IsFile(info)
			if !isFile {
				return nil
			}
			size += info.Size()
			return nil
		},
	)
	if err != nil {
		return 0, fmt.Errorf("unable find size of dir: %w", err)
	}
	return int(size / 1e6), nil
}
