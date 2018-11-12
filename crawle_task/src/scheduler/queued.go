package scheduler

import "engine"

type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (qs *QueueScheduler) Submit(r engine.Request) {
	qs.requestChan <- r
}

func (qs *QueueScheduler) WorkerReady(
	w chan engine.Request) {
	qs.workerChan <- w
}
func (qs *QueueScheduler) ConfigureMasterWorkerChan(r chan engine.Request) {
	panic("implement me")
}

func (qs *QueueScheduler) Run() {
	// 生成 chan
	qs.workerChan = make(chan chan engine.Request)
	qs.requestChan = make(chan engine.Request)

	//go协程
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for  {
			var activeRequest engine.Request
			var activeWorker chan engine.Request

			if len(requestQ) > 0 &&
				len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-qs.requestChan:
				requestQ = append(requestQ, r)
			case w := <-qs.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker<-activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
			/*r := <-qs.requestChan
			w := <-qs.workerChan*/
		}
	}()
}