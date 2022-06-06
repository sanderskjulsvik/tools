package dedupe_test

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/sander-skjulsvik/tools/dedupe"
	"github.com/sander-skjulsvik/tools/os_spec"
)

var ROOT string = "tmp_test_dir/"

func setup(root string) {
	os.MkdirAll(root, 0740)
	// Setup dirs
	dirs := []string{
		"test_dir1/",
		"test_dir2/",
		"test_dir2/test_dir21/",
		"test_dir3/",
		"test_dir3/test_dir31/",
		"test_dir3/test_dir32/",
		"test_dir3/test_dir31/test_dir321/",
		"test_dir3/test_dir31/test_dir322/",
		"test_dir3/test_dir32/test_dir321/",
		"test_dir3/test_dir32/test_dir322/",
	}
	for _, dir := range dirs {
		os.MkdirAll(root+dir, 0740)
	}
	// Build unique files
	for _, dir := range dirs {
		for i := 1; i <= 3; i++ {
			path := root + dir + "test_file" + fmt.Sprint(i)
			f, err := os.Create(path)
			if err != nil {
				log.Fatalf("Could not create file: %s", path)
			}
			w := bufio.NewWriter(f)
			io.CopyN(w, rand.Reader, 1024*1024+1)

		}
	}

}

func cleanup(path string) {
	os.RemoveAll(path)
}

func TestDedupeFindingDupes(t *testing.T) {
	root := ROOT
	setup(root)
	// Duplicates
	d := [...][2]string{
		{root + "test_dir1/test_file1", root + "test_dir1/test_file_copy1"},
		{root + "test_dir2/test_file1", root + "test_dir1/test_file_copy2"},
		{root + "test_dir3/test_dir32/test_dir321/test_file1", root + "test_dir1/test_file_copy3"},
		{root + "test_dir3/test_dir32/test_dir321/test_file1", root + "test_dir3/test_dir32/test_dir322/test_file_copy1"},
	}
	// Create copies
	for _, c := range d {
		dedupe.Copy(c[0], c[1])
	}

	dupes := dedupe.Dupes{}
	files := dedupe.GetFiles(root)
	dupes.BuildDuplicates(files)
	duplicates := dupes.GetDupes()

	// Correct number of dupes
	n := 0
	for _, dupes := range duplicates {
		n += len(dupes)
	}
	if n != len(d) {
		t.Errorf("Wrong number of duplicates. Expected %d, got %d.%s", len(d), n, os_spec.LineBreak)
	}

	cleanup(root)
}
