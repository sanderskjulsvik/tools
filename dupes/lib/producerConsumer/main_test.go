package producerConsumer_test

import (
	"testing"

	"github.com/sander-skjulsvik/tools/dupes/lib/common"
	"github.com/sander-skjulsvik/tools/dupes/lib/producerConsumer"
)

func TestMain(m *testing.M) {
	common.TestRun(common.DEFAULT_TEST_DIR, producerConsumer.Run)
}
