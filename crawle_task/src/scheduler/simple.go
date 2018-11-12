package scheduler

import "engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (ss *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	ss.workerChan = c
}

func (ss *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		ss.workerChan <- r
	}()
}