package main

import (
	"fmt"
	"sync"
)

// CounterStruct - структура, содержащая инкрементируемое значение и Mutex
type CounterStruct struct {
	counter int
	mu      sync.Mutex
}

// inc - функция инкрементирования
func (c *CounterStruct) inc() {
	c.mu.Lock()
	c.counter++
	c.mu.Unlock()
}

// New - создание нового объекта типа CounterStruct
func New() *CounterStruct {
	return &CounterStruct{
		counter: 0,
		mu:      sync.Mutex{},
	}
}

func main() {
	var goroutinesNumber int

	fmt.Print("Введите число инкрементирующих горутин: ")
	_, err := fmt.Scan(&goroutinesNumber)
	if err != nil {
		fmt.Printf("ошибка при получении числа горутин - %s", err)
		return
	}

	if goroutinesNumber <= 0 {
		fmt.Println("некорректное число горутин")
		return
	}

	// создаем новый объект типа CounterStruct
	currentCounter := New()

	// инициализируем WaitGroup для синхронизации горутин
	wg := &sync.WaitGroup{}

	for i := 0; i < goroutinesNumber; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// внутри горутины инкрементируем на удвоенное количество инкрементирующих горутин
			for i := 0; i < goroutinesNumber*2; i++ {
				currentCounter.inc()
			}
		}(wg)
	}

	wg.Wait()

	fmt.Printf("Результат конкурентной инкрементации %d\n", currentCounter.counter)
}
