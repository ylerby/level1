package main

import "fmt"

func binarySearch(target int, arr []int) (int, bool) {
	var mid, val int

	// устанавливаем изначальные границы поиска
	low := 0
	high := len(arr) - 1

	// цикл работает до тех пор, пока верхняя граница больше или равна нижней, если границы сдвигаются
	// до того момента, когда нижняя граница больше верхней, то значение не найдено
	for low <= high {
		mid = (low + high) / 2
		val = mid
		// найдено искомое значение
		if val == target {
			return val, true
		}
		// если текущее значение больше искомого, то убираем все числа, которые больше
		// "отсекаем" правую часть значений
		if val > target {
			high = mid - 1
			// если текущее значение меньше искомого, то убираем все числа, которые меньше
			// "отсекаем" левую часть значений
		} else {
			low = mid + 1
		}
	}
	return -1, false
}

func main() {
	// объявление переменных
	var (
		length int
		value  int
		arr    []int
	)

	// считывание длины слайса из буфера
	fmt.Print("Введите длину слайса: ")
	_, err := fmt.Scan(&length)
	if err != nil {
		fmt.Printf("ошибка при получении длины слайса - %s\n", err)
		return
	}

	// считывание искомого значения из буфера
	fmt.Print("Введите искомое значение: ")
	_, err = fmt.Scan(&value)
	if err != nil {
		fmt.Printf("ошибка при получении значения - %s\n", err)
		return
	}

	// инициализация слайса длиной length
	arr = make([]int, length)

	// заполняем слайс значениями
	for i := 0; i < length; i++ {
		arr[i] = i
	}

	fmt.Printf("Искомое значение: %d\nИсходный слайс: %v\n", value, arr)

	// вызываем функцию бинарного поиска
	result, found := binarySearch(value, arr)
	if !found {
		fmt.Printf("Значение %d не найдено\n", value)
		return
	} else {
		fmt.Printf("Значение %d найдено\n", result)
	}
}
