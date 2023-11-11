package actor

type Actor interface {
	Start()
	AddTask(task TaskExecutor)
	SetCloseSig()
	Close()
}

type TaskExecutor interface {
	Execute() (any, error)
	String() string
}

type TaskSchedular interface {
	Run() error
	AddTask(task TaskExecutor)
	Close()
	Put(val any)
	Receive()
}

type Scaler interface {
	Upscale()
	Downscale()
}
