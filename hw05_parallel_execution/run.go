package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	if len(tasks) == 0 {
		return nil
	}
	var errorCount int64
	errorCount = 0

	tmpTask := 0
	mu := sync.Mutex{}

	wg := sync.WaitGroup{}
	wg.Add(n)

	errorsFlagCh := make(chan struct{}, n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-errorsFlagCh:
					return
				default:
				}

				mu.Lock()
				taskForRun := tmpTask
				tmpTask++
				mu.Unlock()

				if taskForRun < len(tasks) {
					if ok := tasks[taskForRun](); ok != nil {

						atomic.AddInt64(&errorCount, 1)

						if errorCount >= int64(m) {
							errorsFlagCh <- struct{}{}
							return
						}
					}
				} else {
					return
				}
			}
		}()
	}

	wg.Wait()
	select {
	case <-errorsFlagCh:
		return ErrErrorsLimitExceeded
	default:
		return nil
	}
}
