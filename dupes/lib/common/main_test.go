package common_test

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var TEST_DIR string = "test_dir/"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func extractDir(path string) string {
	pathParts := strings.Split(path, "/")
	rightParts := pathParts[:len(pathParts)-1]
	return strings.Join(rightParts, "/")
}

func cleanUp() {
	os.RemoveAll(TEST_DIR)
}

func setup() {
	files := map[string][]string{
		"a93a1ffa-2674-436d-ad10-94dde9c697ea": {"1", "a/1", "a/a/1"},
		"49c2b6ec-4832-494e-a1d0-83670523fe32": {"2", "a/2", "b/a/2"},
		"6a6ec3af-a3c0-4005-8dfb-6cf9ef0416be": {"b/b/1"},
		"a3ea3e1f-9eea-4b25-8ec7-bda1dda06731": {"b/a/b/a/1", "a/3"},
	}

	for content, paths := range files {
		for _, path := range paths {
			err := os.MkdirAll(extractDir(TEST_DIR+path), os.ModePerm)
			if err != nil {
				fmt.Print(err)
				panic(err)
			}
			f, err := os.Create(TEST_DIR + string(path))
			check(err)
			defer f.Close()
			f.WriteString(content)
			f.Sync()
		}
	}
}

func TestAppend(t *testing.T) {
	// dupes := common.Dupes.New(common.Dupes{})
	defer cleanUp()
	setup()

	// dupes.Append()

}
