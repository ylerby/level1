package main

import (
	"fmt"
	"strconv"
)

func setByte(strNumber string, position uint8, value uint8) int64 {
	// формируем результирующее значение
	result := strNumber[0:position] + fmt.Sprintf("%b", value) + strNumber[position+1:]

	// конвертируем значение string -> int
	resultNumber, _ := strconv.Atoi(result)
	return int64(resultNumber)
}

func main() {
	// объявление переменных: число, позицию на которую нужно вставить бит,
	// значение, которое будет установлено в i-ый бит (0 или 1)
	var (
		number   int64
		position uint8
		value    uint8
	)

	// считывание числа из буфера
	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Printf("ошибка при получении числа - %s\n", err)
		return
	}

	// проверка введенного числа на валидность
	if number < 0 {
		fmt.Println("некорректное значение number")
		return
	}

	// представление двоичной записи числа в виде строки
	strNumber := fmt.Sprintf("%b", number)

	fmt.Printf("Число в двоичной форме = %s\n", strNumber)

	// считывание позиции из буфера
	fmt.Print("Введите позицию на которую будет вставлено значение: ")
	_, err = fmt.Scan(&position)
	if err != nil {
		fmt.Printf("ошибка при получении числа - %s\n", err)
		return
	}

	// проверка позиции на валидность. Позиция (индексация с 0) не может быть отрицательной
	// и не может быть больше количества разрядов
	if position < 0 || len(strNumber)-1 < int(position) {
		fmt.Println("некорректное значение number")
		return
	}

	// считывания устанавливаемого значения из буфера (0 или 1)
	fmt.Print("Введите значение, которое будет установлено на i-ую позицию (0/1): ")
	_, err = fmt.Scan(&value)
	if err != nil {
		fmt.Printf("ошибка при получении числа - %s\n", err)
		return
	}

	// проверка значения на валидность
	if value > 1 || value < 0 {
		fmt.Println("некорректное значение number")
		return
	}

	// вызов функции
	result := setByte(strNumber, position, value)

	// вывод результата
	fmt.Printf("Результат = %d\n", result)
}
