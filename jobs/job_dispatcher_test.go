package jobs

import "testing"

// TestRunJobHeartBeat is a test
func TestRunJobDispatch(t *testing.T) {
	DataDispatcher_Run("MARKET")
}
