package main

import (
	"flag"
	"fmt"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

func main() {
	// получение с помощью cli параметра (продолжительность работы горутины)
	duration := flag.Int("duration", 0, "program duration")
	flag.Parse()

	// проверка полученного значения на валидность
	if *duration <= 0 {
		fmt.Println("Ошибка! Введено некорректное время")
		return
	}

	// инициализация канала для работы со значениями
	ch := make(chan int)
	defer close(ch)

	// инициализация канала, отвечающего за завершение работы горутины
	done := make(chan bool, 1)

	// инициализация канала типа <- chan time.Time, в который будет записано значение
	// по истечению времени *duration
	timeCh := time.After(time.Second * time.Duration(*duration))

	go func(ch chan int) {
		for {
			select {
			case <-done:
				fmt.Println("горутина завершена")
				return
			// считываем значение из канала
			case val := <-ch:
				fmt.Printf("считано значение %d\n", val)
			}
		}
	}(ch)

	i := 0
	for {
		i++
		select {
		// по истечению времени равному duration считываем из канала и завершаем выполнение горутины main
		case <-timeCh:
			done <- true
			return
		// записываем в канал значение
		case ch <- i:
			fmt.Printf("записано значение %d\n", i)
			// делаем time.Sleep для более приемлемого вывода считанных значений
			time.Sleep(time.Second / 5)
		}
	}
}
