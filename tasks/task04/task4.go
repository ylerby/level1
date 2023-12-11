package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/

func main() {
	// инициализация переменных-границ формирования псевдо-случайных чисел
	var (
		minValue = 1
		maxValue = 1000
	)

	// получение с помощью cli параметра (количество "воркеров")
	workersNumber := flag.Int("workers", 0, "workers number")
	flag.Parse()

	// проверка на валидность введенного числа горутин
	if *workersNumber <= 0 {
		fmt.Println("Ошибка! Введено некорректное число горутин(воркеров)")
		os.Exit(1)
	}

	// создание буферизированного канала, в который помещаются значения
	ch := make(chan int, *workersNumber)

	// отложенный вызов закрытия канала
	defer close(ch)

	// создание канала для получения системных сигналов
	mainChannel := make(chan os.Signal, 1)
	defer close(mainChannel)

	// определение того, какие сигналы будут обрабатываться и записываться в канал mainChannel
	signal.Notify(mainChannel, os.Interrupt, syscall.SIGTERM)

	for i := 0; i < *workersNumber; i++ {
		// создание канала для получения системных сигналов
		goroutineChannel := make(chan os.Signal, 1)
		signal.Notify(goroutineChannel, os.Interrupt, syscall.SIGTERM)

		// вызов анонимной горутины (воркера)
		go func(index int, ch chan int, systemChannel chan os.Signal) {
			// отложенный вызов закрытия канала
			defer close(systemChannel)

			for {
				select {
				case <-systemChannel:
					fmt.Printf("горутина №%d остановлена!\n", index)
					return
				// чтение значения из буферизированного канала
				case value := <-ch:
					fmt.Printf("горутина №%d, значение %d\n", index, value)
				}
			}
		}(i, ch, goroutineChannel)
	}

	for {
		select {
		// в случае чтения из mainChannel горутина main завершает работу
		case <-mainChannel:
			fmt.Println("горутина main остановлена")
			return
		// запись псевдо-случайного числа в канал
		case ch <- minValue + rand.Intn(maxValue-minValue):
			// time sleep для более приемлемого вывода
			time.Sleep(time.Second / 5)
		}
	}
}
