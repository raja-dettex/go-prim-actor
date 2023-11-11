package actor

import "fmt"

type ActorConfig struct {
	name     string
	Assigner TaskSchedular
}

type ActorMaster struct {
	Config *ActorConfig
}

func NewActorMaster(name string, minSize, maxSize int) *ActorMaster {
	scaler := NewAutoScaler(minSize, maxSize)
	conf := &ActorConfig{
		name:     name,
		Assigner: NewTaskAssigner(minSize, scaler),
	}

	return &ActorMaster{
		Config: conf,
	}
}

func (am *ActorMaster) AddTask(task TaskExecutor) {
	am.Config.Assigner.AddTask(task)
}

func (am *ActorMaster) Start() {
	if err := am.Config.Assigner.Run(); err != nil {
		am.Close(err)
	}
}

func (am *ActorMaster) SetCloseSig() {
	am.Config.Assigner.Close()
}

func (am *ActorMaster) Close(err error) {
	panic(fmt.Sprintf("Closed %v", err))
}

func (am *ActorMaster) PutToAssginerResult(val any) {
	am.Config.Assigner.Put(val)
}

func (am *ActorMaster) ReceiveValues() {
	am.Config.Assigner.Receive()
}
