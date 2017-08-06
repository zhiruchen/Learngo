package goroutines

import (
	"testing"
)

func TestSendMessageToChannel(t *testing.T) {
	SendMesageToChannel()
}

func TestMoveBoxes(t *testing.T) {
	MoveBoxes(10)
}

func TestSyncChannel(t *testing.T) {
	SyncChannel()
}

func TestSendJobds(t *testing.T) {
	SendJobs()
}

func TestMakeTicker(t *testing.T) {
	MakeTicker()
}

func TestWorkerPool(t *testing.T) {
	WorkerPool()
}
