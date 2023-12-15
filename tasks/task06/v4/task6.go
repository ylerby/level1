package main

import (
	"fmt"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

func stop(done chan bool) {
	for {
		select {
		// запись значения в канал
		case done <- true:
			fmt.Println("Горутина завершена")
			return
		default:
			fmt.Println("Горутина продолжает работу")
		}
	}
}

func main() {
	// инициализация канала, отвечающего за завершение горутины
	done := make(chan bool)
	go stop(done)
	time.Sleep(time.Second * 3)

	// считываем значение из канала спустя 3 секунды
	<-done
}
