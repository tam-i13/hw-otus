package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	if len(tasks) == 0 {
		return nil
	}

	tmpTask := 0
	mu := sync.Mutex{}

	errorCount := 0
	mu1 := sync.Mutex{}

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
						mu1.Lock()
						errorCount++
						tmpErrorCount := errorCount
						mu1.Unlock()

						if tmpErrorCount >= m {
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
