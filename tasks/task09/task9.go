package main

import (
	"fmt"
	"math/rand"
)

/*
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

var (
	minValue = -20
	maxValue = 20
)

func generateSlice(valueNumber int) []int {
	generatedSlice := make([]int, valueNumber)
	for i := 0; i < valueNumber; i++ {
		generatedSlice[i] = minValue + rand.Intn(maxValue-minValue)
	}
	return generatedSlice
}

func writeValue(valueSlice []int, out chan<- int) {
	defer close(out)
	for i := range valueSlice {
		out <- valueSlice[i]
	}
}

func multiplyValue(in <-chan int, out chan<- int) {
	for num := range in {
		out <- num * 2
	}
	close(out)
}

func main() {
	var valueNumber int

	fmt.Print("Введите количество элементов: ")
	_, err := fmt.Scan(&valueNumber)
	if err != nil {
		fmt.Printf("ошибка при получении числа - %s", err)
		return
	}

	if valueNumber <= 0 {
		fmt.Println("некорректное число")
		return
	}

	// формирование слайса из псевдо-случайных значений
	valueSlice := generateSlice(valueNumber)

	// инициализация каналов
	firstChan := make(chan int)
	secondChan := make(chan int)

	// записываем значения из слайса в канал
	go writeValue(valueSlice, firstChan)

	// умножение значения из первого канала на 2 и запись во второй канал
	go multiplyValue(firstChan, secondChan)

	// вывод значений
	for value := range secondChan {
		fmt.Printf("Значение: %d\n", value)
	}
}
