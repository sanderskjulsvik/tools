package singleThread_test

import (
	"testing"

	common "github.com/sander-skjulsvik/tools/dupes/lib/common"
	"github.com/sander-skjulsvik/tools/dupes/lib/singleThread"
)

func TestMain(m *testing.T) {
	common.TestRun("test_main_single_thread", singleThread.Run)
}
