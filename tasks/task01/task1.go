package main

import (
	"fmt"
)

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

// Human определение структуры
type Human struct {
	FirstName,
	LastName string
	Age uint8
}

// Action определение структуры, встраиваение структуры Human
type Action struct {
	*Human
}

// NewHuman - создание нового объекта типа Human
func NewHuman(firstName, lastName string, age uint8) (*Human, error) {
	if firstName == "" {
		return nil, fmt.Errorf(`поле "имя" не может быть пустой строкой`)
	}

	if lastName == "" {
		return nil, fmt.Errorf(`поле "фамилия" не может быть пустой строкой`)
	}

	return &Human{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}, nil
}

// getFirstName - метод получения поля "имя"
func (h Human) getFirstName() string {
	return h.FirstName
}

// getLastName - метод получения поля "фамилия"
func (h Human) getLastName() string {
	return h.LastName
}

// getAge - метод получения поля "возраст"
func (h Human) getAge() uint8 {
	return h.Age
}

func main() {
	var (
		newFirstName string
		newLastName  string
		newAge       uint8
	)

	// получение поля "имя" из буфера
	_, err := fmt.Scan(&newFirstName)
	if err != nil {
		fmt.Printf(`Ошибка при получении поля "имя" - %s`, err)
		return
	}

	// получение поля "фамилия" из буфера
	_, err = fmt.Scan(&newLastName)
	if err != nil {
		fmt.Printf(`Ошибка при получении поля "фамилия" - %s`, err)
		return
	}

	// получение поля "возраст" из буфера
	_, err = fmt.Scan(&newAge)
	if err != nil {
		fmt.Printf(`Ошибка при получении поля "возраст" - %s`, err)
		return
	}

	// создание нового объекта типа Human
	human, err := NewHuman(newFirstName, newLastName, newAge)
	if err != nil {
		fmt.Printf("Ошибка при создании - %s", err)
		return
	}

	// создание нового объекта типа Action
	action := &Action{human}

	// вывод полей структуры
	fmt.Printf("Имя: %s\nФамилия: %s\nВозраст: %d\n\n",
		action.FirstName, action.LastName, action.Age)

	// получение полей структуры, с помощью методов
	fmt.Printf("Имя: %s\nФамилия: %s\nВозраст: %d\n",
		action.getFirstName(), action.getLastName(), action.getAge())
}
