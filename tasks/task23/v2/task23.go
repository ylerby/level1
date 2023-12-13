package main

import "fmt"

func main() {
	// инициализация переменной - индекса удаляемого элемента
	var position int

	// считывание числа из буфера
	fmt.Print("Введите индекс удаляемого элемента: ")
	_, err := fmt.Scan(&position)
	if err != nil {
		fmt.Printf("ошибка при получении индекса - %s", err)
		return
	}

	// инициализация слайса
	arr := []int{2, 5, 12, 6, 8, 19, 3, 7, 33, -1, 0, 18, -3}

	// проверка введенного значения на валидность
	if position < 0 || position > len(arr)-1 {
		fmt.Println("Некорректное значение индекса")
		return
	}

	fmt.Printf("Слайс до удаления элемента: %v\n", arr)

	// ставим последний элемент на позицию удаляемого
	arr[position] = arr[len(arr)-1]

	// с помощью среза ограничиваем слайс
	arr = arr[:len(arr)-1]

	fmt.Printf("Слайс после удаления элемента: %v\n", arr)
}
