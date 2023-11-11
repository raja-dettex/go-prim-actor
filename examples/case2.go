package main

import (
	"fmt"
	"time"

	"github.com/raja-dettex/go-prim-actor/actor"
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
	return m.value1, nil
}
func (m NumericMessage) String() string {
	return fmt.Sprintf("message{val1 : %d, val2: %d}", m.value1, m.value2)
}

func main() {
	actorMaster := actor.NewActorMaster("numeric task processor", minSize, maxSize)
	go actorMaster.Start()
	start := time.Now()
	for i := 0; i < 50; i++ {
		go actorMaster.AddTask(NewNumericMessage(i, i+1))
	}

	//go actorMaster.ReceiveValues()
	time.Sleep(time.Second * 20)
	fmt.Println(time.Since(start))
	actorMaster.SetCloseSig()
	select {}
}

// func main() {
// 	ch := make(chan int, 10)
// 	// put 15 integer to the channel
// 	// go func(ch chan int) {
// 	// 	for val := range ch {
// 	// 		fmt.Println("value ", val)
// 	// 	}
// 	// }(ch)
// 	go func(ch chan int) {
// 		for i := 0; i < 15; i++ {
// 			if len(ch) == 10 {
// 				panic("hello")
// 			}
// 			ch <- i
// 			fmt.Println(len(ch))
// 		}
// 	}(ch)
// 	time.Sleep(time.Second * 3)
// }
