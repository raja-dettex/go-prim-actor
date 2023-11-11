package actor

import "fmt"

type TaskAssigner struct {
	MaxSize     int
	CloseSig    chan bool
	ResultQueue chan any
	isSurplus   chan bool
	Scaler      *AutoScaler
}

func NewTaskAssigner(size int, scaler *AutoScaler) *TaskAssigner {
	return &TaskAssigner{
		MaxSize:     size,
		CloseSig:    make(chan bool),
		ResultQueue: make(chan any),
		isSurplus:   make(chan bool),
		Scaler:      scaler,
	}
}

func (ta *TaskAssigner) Put(val any) {
	ta.ResultQueue <- val
}

func (ta *TaskAssigner) Receive() {
	for result := range ta.ResultQueue {
		fmt.Println(result)
	}
}

func (ta *TaskAssigner) Close() {
	ta.CloseSig <- true
}

func (ta *TaskAssigner) AddTask(task TaskExecutor) {
	if len(ta.Scaler.ch) == ta.MaxSize {
		ta.Scaler.Upscale()
	}
	fmt.Println("task message d", task)
	ta.Scaler.ch <- task
}

func (ta *TaskAssigner) Run() error {
	for {
		select {
		case task := <-ta.Scaler.ch:
			val, err := task.Execute()
			if err != nil {
				fmt.Println(err)
				continue
			}
			go ta.Put(val)
		case val := <-ta.ResultQueue:
			fmt.Println("val ", val)
		case sig := <-ta.CloseSig:
			if sig {
				return fmt.Errorf("assigner received an close signal")
			}
		default:
			continue
		}
	}
}
