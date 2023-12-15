package main

import (
	"fmt"
)

func grouping(temperatureSlice []float64) map[int][]float64 {
	// инициализируем map`у, ключ - группа чисел (10, 20 и т.д.), а значение слайс из чисел
	// входящих в эту группу
	result := make(map[int][]float64)

	// итерируемся по значениям
	for _, value := range temperatureSlice {
		// находим группу для числа
		group := (int(value) / 10) * 10
		//добавляем значение в map`у
		result[group] = append(result[group], value)
	}

	return result
}

func main() {
	temperatureSlice := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	result := grouping(temperatureSlice)

	fmt.Printf("Результат группирования чисел: %v\n", result)
}
