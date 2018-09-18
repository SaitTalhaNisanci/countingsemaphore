package countingsemaphore

type sem struct {
	ch chan struct{}
}

// New creates and returns a new counting semaphore with the given size.
func New(size int) Sem {
	return &sem{
		ch : make(chan struct{}, size),
	}
}

type Sem interface {
	// Lock will block until the semaphore has a free resource.
	Lock()
	// Unlock releases a resource.
	Unlock()
}



func (s *sem) p(n int){
	for i := 0 ; i < n ; i++ {
		s.ch <- struct{}{}
	}
}

func (s *sem) v(n int){
	for i := 0 ; i < n ; i++ {
		<-s.ch
	}
}

func (s *sem) Lock() {
	s.p(1)
}

func (s *sem) Unlock() {
	s.v(1)
}