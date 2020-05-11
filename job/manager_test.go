package job

import (
	"testing"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/raphi011/scores/test"
)

func TestManager(t *testing.T) {
	manager := NewManager(logrus.New())

	err := manager.Start(
		Job{
			Name: "Test",

			Interval: 1 * time.Second,
			MaxRuns:  3,

			Do: func() error {
				return nil
			},
		},
	)

	test.Check(t, "manager.Start() failed", err)

	time.Sleep(5 * time.Second)

	j, ok := manager.Job("Test")

	test.Assert(t, "manager.Job() can't retrieve a job", ok)
	test.Assert(t, "expected job to execute 3 times, got %d", j.Execution.Runs == 3, j.Execution.Runs)
}
