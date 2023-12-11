package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

/*
WaitGroup используется для синхронизации горутин, если не использовать WaitGroup, может возникнуть ситуация
когда в map`у записаны не все значения, а горутина main пытается проитерироваться по ней и вывести пары
ключ-значения. Из-за этого возникает ошибка: fatal error: concurrent map iteration and map write
*/

// инициализация константы для количества значений в map`e
const valueNumber = 500

// MapStruct объявление структуры, содержащей саму map`у и mutex для синхронизации горутин
type MapStruct struct {
	mu      *sync.Mutex
	hashMap map[int]int
}

// New - функция создания нового объекта типа MapStruct
func New() *MapStruct {
	return &MapStruct{
		mu:      &sync.Mutex{},
		hashMap: make(map[int]int, valueNumber),
	}
}

// AddValue - функция добавления нового значения в map`у
func (m *MapStruct) AddValue(wg *sync.WaitGroup, key, value int) {
	defer wg.Done()
	m.mu.Lock()
	m.hashMap[key] = value
	m.mu.Unlock()
}

func main() {
	// инициализация переменных-границ формирования псевдо-случайных чисел
	var (
		minValue = 1
		maxValue = 1000
	)

	// создание нового объекта типа MapStruct
	currentMap := New()

	// создание WaitGroup для синхронизации горутин
	wg := &sync.WaitGroup{}

	for i := 0; i < valueNumber; i++ {
		wg.Add(1)
		// формирование псевдо-случайного числа
		value := minValue + rand.Intn(maxValue-minValue)

		// вызов current.AddValue в качестве горутины
		go currentMap.AddValue(wg, i, value)
	}

	// ожидание выполнения горутин
	wg.Wait()

	// вывод map`ы
	fmt.Println(currentMap.hashMap)
}
