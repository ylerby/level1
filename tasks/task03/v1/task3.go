package main

import (
	"fmt"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

// объявление константы для длины массива
const arrSize = 5

func main() {
	// объявление массива заданных чисел
	arr := [arrSize]int{2, 4, 6, 8, 10}

	// создание WaitGroup для синхронизации горутин
	wg := &sync.WaitGroup{}

	// объявление переменной для суммы
	sum := 0

	for _, value := range arr {
		// инкрементируем счетчик горутин
		wg.Add(1)

		// анонимная горутина для вычисления суммы квадратов элементов
		go func(arrValue int) {
			defer wg.Done()

			// вычисление суммы
			sum += arrValue * arrValue
		}(value)
	}

	// ожидание выполнения горутин
	wg.Wait()

	// вывод суммы
	fmt.Printf("Сумма равна = %d\n", sum)
}