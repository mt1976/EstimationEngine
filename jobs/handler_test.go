package jobs

import (
	"testing"

	"github.com/mt1976/ebEstimates/core"
)

// TestRunJobHeartBeat is a test
func TestStart(t *testing.T) {
	core.Initialise()
	Start()
}
