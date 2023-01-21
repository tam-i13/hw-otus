package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	// нет задач, ничего не делаем
	if len(tasks) == 0 {
		return nil
	}

	// счетчик, номер таска который будет запускать
	tmpTask := 0
	mu := sync.Mutex{}

	wg := sync.WaitGroup{}
	wg.Add(n)

	// канал в который передеём ошибку, говорим всем горутинам что пора заканчивать
	errorsFlagCh := make(chan struct{}, n)
	// канал в который забиваем пустыми структурами, если есть ошибка вычитываем из него, если канал закрыт
	// (нет возможности прочитатать то значит количество ошибок закончилось) пора завершать
	errorCountCh := make(chan struct{}, m)
	for i := 0; i < m; i++ {
		errorCountCh <- struct{}{}
	}
	close(errorCountCh)

	// запускам горутины  n штук
	for i := 0; i < n; i++ {
		go func() {
			for {
				// проверяем канал с ошибками
				select {
				case <-errorsFlagCh:
					wg.Done()
					return
				default:
				}

				// забираем из канала номер задачки

				mu.Lock()
				taskForRun := tmpTask
				tmpTask++
				mu.Unlock()

				// проверяем канал true есть задачи можем забирать
				if taskForRun < len(tasks) {
					// запускаем задачу, проверяем успешность выполнения
					if ok := tasks[taskForRun](); ok != nil {
						// ошибка, читаем из канала ошибок, если он закрыт то завершем выполнение
						_, closeCh := <-errorCountCh
						if !closeCh {
							errorsFlagCh <- struct{}{}
							wg.Done()
							return
						}
					}
				} else {
					// задач нет завершаем
					wg.Done()
					return
				}
			}
		}()
	}

	wg.Wait()
	// проверка канала с ошибками
	select {
	case <-errorsFlagCh:
		return ErrErrorsLimitExceeded
	default:
		return nil
	}
}
