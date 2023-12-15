package main

import (
	"context"
	"fmt"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

func stop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина завершена")
			return
		default:
			fmt.Println("Горутина продолжает работу")
		}
	}
}

func main() {

	// инициализация контекста
	ctx, cancel := context.WithCancel(context.Background())

	// вызов горутины
	go stop(ctx)
	time.Sleep(2 * time.Second)

	// завершение контекста
	cancel()
}
