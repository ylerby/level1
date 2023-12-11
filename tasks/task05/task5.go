package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	duration := flag.Int("duration", 0, "program duration")
	flag.Parse()

	if *duration <= 0 {
		fmt.Println("Ошибка! Введено некорректное время")
		os.Exit(1)
	}

	ch := make(chan int)
	defer close(ch)

	done := make(chan bool, 1)

	timeCh := time.After(time.Second * time.Duration(*duration))

	go func(ch chan int) {
		for {
			select {
			case <-done:
				fmt.Println("горутина завершена")
				return
			case val := <-ch:
				fmt.Printf("считано значение %d\n", val)
			}
		}
	}(ch)

	i := 0
	for {
		i++
		select {
		case <-timeCh:
			done <- true
			return
		case ch <- i:
			fmt.Printf("записано значение %d\n", i)
			time.Sleep(time.Second / 5)
		}
	}
}
