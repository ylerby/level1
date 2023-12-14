package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func stringReverse(currentString string) string {
	// добавляем слова в слайс разделив их по сепаратору " " (пробел)
	stringSlice := strings.Split(currentString, " ")

	// инициализируем слайс, в который добавляем значения
	reversedStringSlice := make([]string, len(stringSlice))

	// итерируемся по слайсу в обратном порядке и добавляем значения
	for i := len(stringSlice) - 1; i >= 0; i-- {
		reversedStringSlice = append(reversedStringSlice, stringSlice[i])
	}

	// обратно соединениям слайс в строку
	reversedString := strings.Join(reversedStringSlice, " ")

	return reversedString
}

func main() {
	// считываем строку и обрабатываем ошибку
	fmt.Printf("Введите строку: ")
	currentString, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Printf("ошибка при считывании строки из буфера - %s", err)
		return
	}

	// проверяем строку на то что строка пустая
	if currentString == "" {
		fmt.Println("пустая строка")
		return
	}

	fmt.Printf("Изначальная строка - %s\n", currentString)
	result := stringReverse(currentString)
	fmt.Printf("Перевернутая строка - %s\n", result)
}
