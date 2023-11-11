package actor

type AutoScaler struct {
	Max int
	Min int
	ch  chan TaskExecutor
}

func NewAutoScaler(max, min int) *AutoScaler {
	return &AutoScaler{
		Max: max,
		Min: min,
		ch:  make(chan TaskExecutor, min),
	}
}

func (s *AutoScaler) Upscale() {
	newCh := make(chan TaskExecutor, s.Max)
	for val := range s.ch {
		newCh <- val
	}
	s.ch = newCh
}
func (s *AutoScaler) Downscale() {
	newCh := make(chan TaskExecutor, s.Min)
	for val := range s.ch {
		newCh <- val
	}
	s.ch = newCh
}
