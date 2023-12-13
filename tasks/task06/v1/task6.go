package main

import (
	"fmt"
	"sync"
)

func stop(wg *sync.WaitGroup, stopValue, maxValue int) {
	defer wg.Done()

	// цикл для нахождения искомого значения
	for i := 0; i < maxValue; i++ {

		// если искомое значение найдено, горутина завершена
		if i == stopValue {
			fmt.Printf("Горутина завершена, достигнув значение %d\n", i)
			return
		}
		fmt.Printf("Значение %d\n", i)
	}
}

func main() {
	// инициализация переменных, отвечающих за завершение горутины
	var (
		stopValue int
		maxValue  int
	)

	// ввод переменных из буфера
	fmt.Print("Введите искомое значение: ")
	_, err := fmt.Scan(&stopValue)
	if err != nil {
		fmt.Println("ошибка при получении числа")
		return
	}

	if stopValue <= 0 {
		fmt.Println("некорректное значение stopValue")
		return
	}

	fmt.Print("Введите максимальное значение: ")
	_, err = fmt.Scan(&maxValue)
	if err != nil {
		fmt.Println("ошибка при получении числа")
		return
	}

	if maxValue < stopValue {
		fmt.Println("некорректное значение maxValue")
		return
	}

	// инициализация WaitGroup для ожидания выполнения горутины
	wg := &sync.WaitGroup{}

	wg.Add(1)

	// вызов горутины
	go stop(wg, stopValue, maxValue)
	wg.Wait()
}
