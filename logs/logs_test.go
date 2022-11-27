package logs

import (
	"errors"
	"testing"
)

func Test_System(t *testing.T) {
	System("System")
}

func Test_Success(t *testing.T) {
	Success("Success")
}

func Test_Error(t *testing.T) {
	Error("Error", errors.New("Error"))
}
