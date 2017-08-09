package goroutines

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataFileWrite(t *testing.T) {
	df, err := NewDataFile("x.txt", 10)
	assert.Equal(t, nil, err)
	ss := []string{"1", "2", "#", "hello"}
	w := make(chan bool)
	go func() {
		for _, s := range ss {
			go df.Write([]byte(s))
		}
		w <- true
	}()
	<-w
}
