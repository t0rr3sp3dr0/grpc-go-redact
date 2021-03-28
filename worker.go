package main

import (
	"fmt"
	"log"
	"sync"
)

type WorkQueue struct {
	wg      sync.WaitGroup
	jobChan chan *ParseInfo
}

func NewWorkQueue() *WorkQueue {
	return &WorkQueue{
		jobChan: make(chan *ParseInfo),
	}
}

func (w *WorkQueue) StartWorkers(num int) {
	for i := 0; i < num; i++ {
		w.wg.Add(1)
		go w.worker()
	}
}

func (w *WorkQueue) AddJob(job *ParseInfo) bool {
	select {
	case w.jobChan <- job:
		return true
	default:
		return false
	}
}

func (w *WorkQueue) WaitForJobs() {
	w.wg.Wait()
}

func (w *WorkQueue) Shutdown() {
	close(w.jobChan)
}

// TODO: correctly handle errors
func (w *WorkQueue) worker() {
	defer w.wg.Done()

	for {
		select {
		case job, ok := <-w.jobChan:
			if !ok {
				return
			}

			if err := GenerateStringFunc(job); err != nil {
				log.Println(err)
				continue
			}

			if err := writeASTToFile(job); err != nil {
				log.Println(err)
				continue
			}

			fmt.Println("completed generated file: ", job.OutputFile)
		default:
			return
		}
	}
}
