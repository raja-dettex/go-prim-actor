package actor

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	ch := make(chan int)
	// put 15 integer to the channel
	for i := 0; i < 15; i++ {
		ch <- i
	}
}

func TestMain(t *testing.T) {
	fmt.Println("hello")
}
