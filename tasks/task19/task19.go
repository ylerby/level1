package main

import "fmt"

func wordReverse(currentString string) string {
	runeSlice := []rune(currentString)

	reversedString := make([]rune, len(runeSlice))

	for i := len(runeSlice) - 1; i >= 0; i-- {
		reversedString = append(reversedString, runeSlice[i])
	}

	return string(reversedString)
}

func main() {
	// объявление строки, символы в которой нужно развернуть
	var currentString string

	// получение строки из буфера и обработка ошибки
	fmt.Print("Введите строку: ")
	_, err := fmt.Scan(&currentString)
	if err != nil {
		fmt.Printf("ошибка при получении строки - %s", err)
		return
	}

	fmt.Printf("Изначальная строка - %s\n", currentString)
	result := wordReverse(currentString)
	fmt.Printf("Перевернутая строка - %s\n", result)
}
