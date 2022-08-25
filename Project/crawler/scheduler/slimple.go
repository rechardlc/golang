package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workerChan chan<- engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(r chan<- engine.Request) {
	s.workerChan = r
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workerChan <- r
	}()
}
