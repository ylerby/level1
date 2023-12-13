package main

import (
	"fmt"
)

func getType(value interface{}) {
	// type switching с выводом типа аргумента
	switch value.(type) {
	case int:
		fmt.Printf("переменная типа %T\n", value)
	case bool:
		fmt.Printf("переменная типа %T\n", value)
	case string:
		fmt.Printf("переменная типа %T\n", value)
	case chan interface{}:
		fmt.Printf("переменная типа %T\n", value)
	default:
		fmt.Printf("переменная типа %T - "+
			"не является переменной ни одного из предложенных типов\n", value)
	}
}
func main() {
	// инициализируем значения для проверки
	intValue := 10
	stringValue := "str"
	chanValue := make(chan interface{})
	defaultCaseValue := int64(intValue)

	// вызываем анонимную функцию с переменным числом аргументов типа interface{}
	// тип interface{} имплементирует любой другой тип
	// сделано во избежании нескольких вызовов функции подряд
	func(in ...interface{}) {
		for _, value := range in {
			getType(value)
		}
	}(intValue, stringValue, chanValue, defaultCaseValue)
}
