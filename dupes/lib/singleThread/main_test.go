package singleThread_test

import (
	"testing"

	"github.com/sander-skjulsvik/tools/dupes/lib/common"
)

var TEST_DIR string = "test_dir/"

func TestRun(t *testing.T) {
	defer common.CleanUp(TEST_DIR)
	common.Setup(TEST_DIR)

}
