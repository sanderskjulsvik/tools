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

func CreateLargeFile(path string, size int64, mod int64) error {
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

	writeBufferSzie := WRITE_BUFFER_SIZE
	if WRITE_BUFFER_SIZE > size {
		writeBufferSzie = size
	}

	// Prepp for writing
	b := make([]byte, writeBufferSzie)
	iterations := size / writeBufferSzie
	log.Printf("createLargeFile: iterations: %d", iterations)
	for range iterations {
		// Make content to write
		for i := int64(0); i < writeBufferSzie; i++ {
			b[i] = byte(i % mod)
		}
		// Write content
		if err != nil {
			if err == io.EOF {
				return nil
			}
			continue
		}
		file.Write(b)
	}

	return nil
}

func FilesEqual(src, dst string) (bool, error) {

	srcStat, err := os.Stat(src)
	if os.IsNotExist(err) {
		return false, fmt.Errorf("Could not find src file")
	}
	if err != nil {
		return false, fmt.Errorf("error with src: %s", err)
	}
	srcSize := srcStat.Size()

	dstStat, err := os.Stat(dst)
	if os.IsNotExist(err) {
		return false, fmt.Errorf("Could not find dst file")
	}
	if err != nil {
		return false, fmt.Errorf("error with dst: %s", err)
	}
	dstSize := dstStat.Size()

	if dstSize != srcSize {
		return false, nil
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return false, fmt.Errorf("Failed to open src file: %s, err: %v", src, err)
	}
	dstFile, err := os.Open(dst)
	if err != nil {
		return false, fmt.Errorf("Failed to open dst file: %s, err: %v", dst, err)
	}
	srcConent := make([]byte, srcSize)
	srcFile.Read(srcConent)
	dstContent := make([]byte, dstSize)
	dstFile.Read(dstContent)

	return bytes.Equal(srcConent, dstContent), nil
}
