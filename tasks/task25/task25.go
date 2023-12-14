package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	// иниуиализируем канал типа <- time.Time, из которого будет считано значение через время равное duration
	timer := time.After(duration)
	<-timer
}

func main() {
	var duration int

	fmt.Print("Введите продолжительность (секунды): ")
	_, err := fmt.Scan(&duration)
	if err != nil {
		fmt.Printf("ошибка при получении продолжительности - %s", err)
		return
	}

	if duration < 0 {
		fmt.Println("некорректное значение продолжительности")
		return
	}

	fmt.Println("Начало работы sleep")
	sleep(time.Second * time.Duration(duration))
	fmt.Println("Завершение работы sleep")
}
