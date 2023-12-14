package main

import (
	"fmt"
	"sync"
)

/*
 Написать программу, которая конкурентно рассчитает значение квадратов чисел,
 взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

// инициализация константы для длины массива
const arraySize = 5

func main() {
	// инициализация массива
	arr := [arraySize]int{2, 4, 6, 8, 10}

	// создание WaitGroup для синхронизации горутин
	wg := &sync.WaitGroup{}

	for index, value := range arr {
		// инкрементируем счетчик горутин
		wg.Add(1)

		// анонимная горутина для вывода квадрата элементов
		go func(wg *sync.WaitGroup, elementIndex, element int) {
			// в конце выполнения горутины декрементируем счетчик горутин
			defer wg.Done()

			// вывод элемента, его квадрата, индекс
			fmt.Printf("Элемент = %d, Квадрат элемента = %d, Индекс = %d\n",
				element, element*element, elementIndex)
		}(wg, index, value)
	}

	// ожидание выполнения горутин
	wg.Wait()
}
