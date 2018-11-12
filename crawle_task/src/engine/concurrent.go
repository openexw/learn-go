package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
	ReadyNotifier
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
} 
// run 并发版
func (ce *ConcurrentEngine) Run1(seeds ...Request) {


	// create worker
	in := make(chan Request)
	out := make(chan ParseResult)
	ce.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i<ce.WorkerCount ; i++ {
		createWorker1(in, out)
	}

	for _, r := range seeds {
 		ce.Scheduler.Submit(r)
	}


	// 输入
	userCount := 0
	for  {
		result := <- out
		for _, item := range result.Items {
			userCount++
			fmt.Printf("[Got Item]: #%d %v\n", userCount, item)
		}

		for _, request := range result.Requests {
			ce.Scheduler.Submit(request)
		}
	}
}

func (ce *ConcurrentEngine) Run(seeds ...Request) {


	// create worker
	//in := make(chan Request)
	out := make(chan ParseResult)
	//ce.Scheduler.ConfigureMasterWorkerChan(in)
	ce.Scheduler.Run()

	for i := 0; i<ce.WorkerCount ; i++ {
		createWorker(out, ce.Scheduler)
	}

	for _, r := range seeds {
		ce.Scheduler.Submit(r)
	}


	// 输入
	userCount := 0
	for  {
		result := <- out
		for _, item := range result.Items {
			userCount++
			fmt.Printf("[Got Item]: #%d %v\n", userCount, item)
		}

		for _, request := range result.Requests {
			ce.Scheduler.Submit(request)
		}
	}
}


// create worker 并发版
func createWorker1(in chan Request, out chan ParseResult) {
	go func() {
		for {
			// tell scheduler i'm ready
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func createWorker(out chan ParseResult, ready ReadyNotifier) {
	in := make(chan Request)
	go func() {
		for {
			// tell scheduler i'm ready
			ready.WorkerReady(in)
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}