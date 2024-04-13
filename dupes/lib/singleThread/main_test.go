package singleThread_test

import (
	"testing"

	common "github.com/sander-skjulsvik/tools/dupes/lib/common"
	"github.com/sander-skjulsvik/tools/dupes/lib/singleThread"
)

func TestMain(m *testing.M) {
	common.TestRun(common.DEFAULT_TEST_DIR, singleThread.Run)
}
