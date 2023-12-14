package main

import (
	"fmt"
	"io"
	"math/rand"
	"strings"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.
*/

// Использование justString[:100] может приводить к утечке памяти,
// более безопасным вариантом будет использование copy(destination, source)
// также стоит отказаться от использования (var justString string) глобальных переменных (от тех глобальных переменных,
// которые можно задать и внутри функции без потери в функционале)
func createHugeString(length int) (string, error) {
	builder := &strings.Builder{}
	// формируем длинную строку из случайных букв
	for i := 0; i < length; i++ {
		runeValue := 'a' + rune(rand.Intn(26))
		stringValue := string(runeValue)
		_, err := io.WriteString(builder, stringValue)
		if err != nil {
			return "", err
		}
	}
	return builder.String(), nil
}

func someFunc() error {
	v, err := createHugeString(1 << 10)
	if err != nil {
		return fmt.Errorf("ошибка при создании строки - %s", err)
	}

	fmt.Printf("Полученная строка - %v\n", v)
	justString := make([]byte, 100)

	copy(justString, v)

	finalString := string(justString)
	fmt.Printf("Скопированное значение ([]byte) - %v, type - %T\n", justString, justString)
	fmt.Printf("Скопированное значение (string) - %v, type - %T\n", finalString, finalString)
	return nil
}

func main() {
	err := someFunc()
	if err != nil {
		fmt.Printf("ошибка в функции someFunc - %s", err)
		return
	}
}
