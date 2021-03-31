package main

import (
	"context"
	"sync"

	"github.com/samkreter/go-core/log"
)

type jobType *ParseInfo

type WorkQueue struct {
	wg      sync.WaitGroup
	jobChan chan jobType
}

func NewWorkQueue(maxQueueSize int) *WorkQueue {
	return &WorkQueue{
		jobChan: make(chan jobType, maxQueueSize),
	}
}

func (w *WorkQueue) StartWorkers(num int) {
	for i := 0; i < num; i++ {
		w.wg.Add(1)
		go w.worker()
	}
}

func (w *WorkQueue) AddJob(job jobType) bool {
	select {
	case w.jobChan <- job:
		return true
	default:
		return false
	}
}

func (w *WorkQueue) NumJobs() int {
	return len(w.jobChan)
}

func (w *WorkQueue) WaitForJobs() {
	w.wg.Wait()
}

func (w *WorkQueue) Shutdown() {
	close(w.jobChan)
}

// TODO: correctly handle errors
func (w *WorkQueue) worker() {
	ctx := context.Background()
	logger := log.G(ctx)

	defer w.wg.Done()

	for {
		select {
		case job, ok := <-w.jobChan:
			if !ok {
				return
			}

			if err := GenerateStringFunc(job); err != nil {
				logger.Errorln("failed to generate string func with err: ", err)
				continue
			}

			if err := writeASTToFile(job); err != nil {
				logger.Errorln("failed to write ast to file with err: ", err)
				continue
			}

			logger.Debugln("Completed task: ", job.OutputFile)

		default:
			return
		}
	}
}
