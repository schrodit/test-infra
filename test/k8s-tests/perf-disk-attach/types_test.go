package perf_disk_attach_test

import (
	"fmt"
	"time"
)

// Result represents a single json output of a test statefulset
type Result struct {
	Name       string `json:"name"`
	Duration   int64  `json:"duration"`
	VolumesNum int    `json:"volumes"`
	duration   time.Duration
}

// Name generates the name of a statefulset with a specific number
func Name(i int) string {
	return fmt.Sprintf("test-%d", i)
}
