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
		for outerIndex, outerValue := range currentString {
			for innerIndex, innerValue := range currentString {
				// итерируемся по строке и сверяем равны ли значения
				// но при этом индексы должны быть различными (чтобы не учитывалось одно и то же значение)
				if outerValue == innerValue && outerIndex != innerIndex {
					// если найдены равные значения и индексы не равны, значит строка содержит неуникальные
					// значения
					return false
				}
			}
		}
		return true
	}(currentString)

	fmt.Printf("Результат проверки строки на уникальность: %v\n", result)
}
