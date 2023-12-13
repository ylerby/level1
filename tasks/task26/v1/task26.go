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
