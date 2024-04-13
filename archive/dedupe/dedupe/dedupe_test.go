package dedupe_test

import (
	"bufio"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/sander-skjulsvik/tools/dedupe/dedupe"
	"github.com/sander-skjulsvik/tools/dedupe/util"
	"github.com/sander-skjulsvik/tools/os_spec"
)

var ROOT string = "tmp_test_dir/"
var void bool = true

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
		util.Copy(c[0], c[1])
	}

	dupes := dedupe.Dupes{}
	files := dedupe.GetFiles(root)
	dupes.BuildDuplicates(files)
	duplicates := dupes.GetDuplicates()

	// Correct number of dupes
	n := 0
	for _, dupes := range duplicates {
		n += len(dupes) - 1
	}
	if n != len(d) {
		t.Errorf("Wrong number of duplicates. Expected %d, got %d.%s", len(d), n, os_spec.LineBreak)
	}

	// Correctdupe: Calc in exp

	// calculated maps
	calculated := [len(d)]map[string]bool{}
	i := 0
	for _, d := range duplicates {
		calculated[i] = make(map[string]bool)
		for _, f := range d {
			calculated[i][f.Path] = true
		}
		i++
	}

	exp := [len(d)]map[string]bool{
		{d[0][0]: true, d[0][1]: true},
		{d[1][0]: true, d[1][1]: true},
		{d[2][0]: true, d[2][1]: true, d[3][1]: true},
	}
	for _, c := range calculated {
		x := false
		for _, e := range exp {
			if reflect.DeepEqual(c, e) {
				x = true
				break
			}
		}
		if x == false {
			cJson, err := json.Marshal(c)
			if err != nil {
				fmt.Println("Unable to marshal calculated to json")
			}
			eJson, err := json.Marshal(exp)
			if err != nil {
				fmt.Println("Unable to marshal expected to json")
			}
			t.Errorf("Could not find calculated %s, in expected %s%s", cJson, eJson, os_spec.LineBreak)
		}
	}

	// Correctdupe: Exp in calc

	for _, e := range exp {
		x := false
		for _, c := range calculated {
			if reflect.DeepEqual(c, e) {
				x = true
				break
			}
		}
		if x == false {
			cJson, err := json.Marshal(calculated)
			if err != nil {
				fmt.Println("Unable to marshal calculated to json")
			}
			eJson, err := json.Marshal(e)
			if err != nil {
				fmt.Println("Unable to marshal expected to json")
			}
			t.Errorf("Could not find expected %s, in calculated %s%s", cJson, eJson, os_spec.LineBreak)
		}
	}

	cleanup(root)
}
