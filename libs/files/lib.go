package files

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

const BUFFERSIZE int = 65_536
const WRITE_BUFFER_SIZE = int64(1e6)

// cp from: https://opensource.com/article/18/6/copying-files-go
func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return err
	}
	return out.Sync()
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

func CreateLargeFile(path string, size int64) error {
	// Setup File
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return fmt.Errorf("unable to create dir: %w", err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to create file: %w", err)
	}
	defer file.Close()

	// Prepp for writing
	b := make([]byte, WRITE_BUFFER_SIZE)
	reader := bytes.NewReader(b)
	func(r io.Reader, w *os.File) {
		iterations := size / WRITE_BUFFER_SIZE
		log.Printf("createLargeFile: iterations: %d", iterations)
		for range iterations {
			// Make content to write
			for i := int64(0); i < WRITE_BUFFER_SIZE; i++ {
				b[i] = byte(i % 256)
			}
			// Write content
			if err != nil {
				if err == io.EOF {
					return
				}
				continue
			}
			w.Write(b)
		}
	}(reader, file)

	return nil
}
