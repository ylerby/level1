package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу,
которая перемножает, делит, складывает, вычитает две числовых переменных a,b,
значение которых > 2^20.
*/

func Add(a *big.Int, b *big.Int) *big.Int {
	result := new(big.Int)
	return result.Add(a, b)
}

func Sub(a *big.Int, b *big.Int) *big.Int {
	result := new(big.Int)
	return result.Sub(a, b)
}

func Mul(a *big.Int, b *big.Int) *big.Int {
	result := new(big.Int)
	return result.Mul(a, b)
}

func Div(a *big.Int, b *big.Int) *big.Int {
	result := new(big.Int)
	return result.Div(a, b)
}

func main() {
	a := big.NewInt(0)
	b := big.NewInt(0)

	// указываем значение из строки, так как задать значение напрямую не получится, возникнет ошибка
	// переполнения
	a.SetString("600000000000000000000", 10)
	b.SetString("300000000000000000000", 10)

	fmt.Printf("Результат сложения: %v\n", Add(a, b))
	fmt.Printf("Результат вычитания: %v\n", Sub(a, b))
	fmt.Printf("Результат умножения: %v\n", Mul(a, b))
	fmt.Printf("Результат деления: %v\n", Div(a, b))
}
