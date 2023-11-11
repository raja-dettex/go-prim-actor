package actor

import (
	"fmt"
	"testing"
	"time"
)

const (
	minSize = 30
	maxSize = 100
)

type NumericMessage struct {
	value1 int
	value2 int
}

func NewNumericMessage(value1, value2 int) NumericMessage {
	return NumericMessage{
		value1: value1,
		value2: value2,
	}
}

func (m NumericMessage) Execute() (any, error) {
	time.Sleep(time.Millisecond * 100)
	return m.value1 + m.value2, nil
}
func (m NumericMessage) String() string {
	return fmt.Sprintf("message{val1 : %d, val2: %d}", m.value1, m.value2)
}

func BenchmarkTestIoActor(b *testing.B) {
	am := NewActorMaster("numericProcessor", minSize, maxSize)
	go am.Start()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			go am.AddTask(NewNumericMessage(j, j+1))
		}
	}
}
