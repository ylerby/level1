package main

import (
	"fmt"
	"math/rand"
)

// arrFill - функция создания слайса и заполнения его случайными значениями
func arrFill(length int) []int {
	var (
		minValue = -1000
		maxValue = 1000
	)

	arr := make([]int, length)

	for i := 0; i < length; i++ {
		arr[i] = minValue + rand.Intn(maxValue-minValue)
	}

	return arr
}

// quickSort - функция сортировки слайса
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	pivot := rand.Int() % len(arr)

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}

func main() {
	var length int

	fmt.Print("Введите длину сортируемого слайса: ")
	_, err := fmt.Scan(&length)
	if err != nil {
		fmt.Printf("ошибка при получении длины слайса - %s", err)
		return
	}

	if length < 0 {
		fmt.Println("Длина слайса не может быть отрицательной")
		return
	}
	nonSortedArr := arrFill(length)

	fmt.Printf("Неотсортированный слайс - %v\n", nonSortedArr)
	sortedArr := quickSort(nonSortedArr)
	fmt.Printf("Отсортированный слайс - %v\n", sortedArr)
}
