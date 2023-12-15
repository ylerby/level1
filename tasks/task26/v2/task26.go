package main

import "fmt"

/*
Разработать программу, которая проверяет,
что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.
Например:
abcd — true abCdefAaf — false aabcd — false
*/

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
