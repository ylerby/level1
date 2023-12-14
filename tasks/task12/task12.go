package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	// инициализируем map`у с ключом типа string а значением типа struct{} (используется для экономии памяти)
	set := make(map[string]struct{})

	// итерируемся по слайсу и вносим значения в "множество"
	for _, value := range arr {
		set[value] = struct{}{}
	}

	fmt.Printf("Множество для значений слайса - %v", set)
}
