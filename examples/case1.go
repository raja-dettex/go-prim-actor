package main

type Message struct {
	msg string
}

func (m Message) Execute() (any, error) {
	return m.msg, nil
}

func NewMessage(msg string) Message {
	return Message{
		msg: msg,
	}
}

// func main() {
// 	taskCh := make(chan actor.TaskExecutor)
// 	actorMaster := actor.NewActorMaster("master actor", 10, taskCh)
// 	//wg.Add(1)
// 	//defer wg.Done()
// 	go actorMaster.Start()
// 	for i := 0; i < 30; i++ {
// 		task := NewMessage(fmt.Sprintf("hello %d", i))
// 		actorMaster.AddTask(task)
// 	}
// 	go actorMaster.ReceiveValues()
// 	//wg.Wait()
// 	time.Sleep(time.Second * 3)
// 	actorMaster.SetCloseSig()
// 	select {}

// }
