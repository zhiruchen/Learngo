package goroutines

import (
	"testing"
)

func TestSyncRoutine(t *testing.T) {
	SyncRoutine()
}

func TestCloseRoutine(t *testing.T) {
	CloseRoutine()
}

func TestForSelect(t *testing.T) {
	ForSelect()
}
