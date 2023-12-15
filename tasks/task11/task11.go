package main

import "fmt"

func merge[T comparable](firstSet, secondSet []T) []T {
	var (
		overallLen int
	)

	// находим большую длину слайса
	if len(firstSet) >= len(secondSet) {
		overallLen = len(firstSet)
	} else {
		overallLen = len(secondSet)
	}

	mergeMap := map[T]bool{}

	for _, value := range firstSet {
		mergeMap[value] = true
	}

	mergedSet := make([]T, 0, overallLen)

	// проверяем, если значение из слайса уже есть в map`e, то добавляем его в результирующий слайс
	for _, value := range secondSet {
		if mergeMap[value] {
			mergedSet = append(mergedSet, value)
		}
	}

	return mergedSet
}

func main() {
	firstIntSet := []int{1, 5, 2, 3, 9, 7, 12, -6, 4}
	secondIntSet := []int{-16, 4, 5, 31, 0, 1}
	intResult := merge(firstIntSet, secondIntSet)
	fmt.Printf("Результат пересечения множеств - %v\n", intResult)

	firstRuneSet := []rune{'a', 'b', 'c', 'd', 'e', 'f'}
	secondRuneSet := []rune{'v', 'w', 'u', 'f', 'a', 'c'}
	runeResult := merge(firstRuneSet, secondRuneSet)

	fmt.Print("Результат пересечения множеств - ")
	for _, value := range runeResult {
		fmt.Printf("%c ", value)
	}

	fmt.Print("\n")

	firstStringSet := []string{"ab", "cd", "ef", "gh"}
	secondStringSet := []string{"sx", "gr", "cb", "ab", "r3"}
	stringResult := merge(firstStringSet, secondStringSet)

	fmt.Print("Результат пересечения множеств - ")
	for _, value := range stringResult {
		fmt.Printf("%s ", value)
	}
	fmt.Print("\n")
}
