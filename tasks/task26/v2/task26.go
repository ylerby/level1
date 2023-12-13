package main

import "fmt"

func main() {
	// объявление переменной строки
	var currentString string

	// считывание строки из буфера
	fmt.Print("Введите строку: ")
	_, err := fmt.Scan(&currentString)
	if err != nil {
		fmt.Printf("ошибка при получении строки %s", err)
		return
	}

	result := func(str string) bool {
		// инициализируем map`у со значением типа bool (по умолчанию false)
		uniqueMap := make(map[rune]bool)
		for _, value := range currentString {
			// если значение по ключу value является true, значит значение не является уникальным
			if uniqueMap[value] {
				return false
			}
			uniqueMap[value] = true
		}
		return true
	}(currentString)

	fmt.Printf("Результат проверки строки на уникальность: %v\n", result)
}
