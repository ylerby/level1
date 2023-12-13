package main

import "fmt"

func main() {
	// инициализация переменных
	var (
		a,
		b int
	)

	// ввод числа a и обработка ошибки
	fmt.Print("Введите число a: ")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Printf("ошибка при получении числа a - %s", err)
		return
	}

	// ввод числа b и обработка ошибки
	fmt.Print("Введите число b: ")
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Printf("ошибка при получении числа b - %s", err)
		return
	}

	fmt.Printf("Значение a - %d, значение b - %d\n", a, b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("Значение a - %d, значение b - %d\n", a, b)
}
