package singleThread_test

import (
	"testing"

	"github.com/sander-skjulsvik/tools/dupes/lib/singleThread"
	"github.com/sander-skjulsvik/tools/dupes/lib/test"
)

func TestMain(t *testing.T) {
	test.TestRun("test_main_single_thread", singleThread.Run, t)
}
