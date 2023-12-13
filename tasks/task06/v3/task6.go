package main

import (
	"fmt"
	"sync"
	"time"
)

func stop(wg *sync.WaitGroup, timeCh <-chan time.Time) {
	defer wg.Done()
	for {
		select {
		// по истечению таймера, из канала считывается значение
		case <-timeCh:
			fmt.Println("Горутина завершена")
			return
		default:
			fmt.Println("Горутина продолжает работу")
		}
	}
}

func main() {

	// инициализация канала типа <- chan time.Time, по истечению time.Second * 2 из него будет считано
	// значение и горутина будет завершена
	timeCh := time.After(time.Second * 2)

	wg := &sync.WaitGroup{}

	wg.Add(1)

	go stop(wg, timeCh)

	wg.Wait()
}
