package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// Генерирует слайс случайных положительных чисел.
func generateRandomElements(size int) []int {
	// Обработка крайнего случая
	if size <= 0 {
		return []int{}
	}

	elements := make([]int, size)
	for i := 0; i < size; i++ {
		// Генерируем числа от 1 до 100_000
		elements[i] = rand.Intn(100_000) + 1
	}
	return elements
}

// Находит максимальное число в слайсе в один поток.
func maximum(data []int) int {
	// Обработка пустого слайса
	if len(data) == 0 {
		return 0
	}

	maxVal := data[0]
	for _, val := range data[1:] {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

// Находит максимальное число в слайсе, используя 8 горутин.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) < CHUNKS {
		return maximum(data)
	}

	chunkSize := len(data) / CHUNKS
	results := make([]int, CHUNKS)

	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		// Вычисляем начальный и конечный индексы
		start := i * chunkSize
		end := start + chunkSize

		// Если последний кусок, то забираем все оставшиеся элементы
		if i == CHUNKS-1 {
			end = len(data)
		}

		// Запускаем горутину
		go func(chunk []int, idx int) {
			defer wg.Done()
			results[idx] = maximum(chunk)
		}(data[start:end], i)
	}

	wg.Wait() // Ждем завершения всех 8 горутин

	// Находим максимум среди 8
	return maximum(results)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	startTime := time.Now()
	maxSingle := maximum(data)
	elapsedSingle := time.Since(startTime)

	// Время переводим в микросекунды, согласно ТЗ
	fmt.Printf("Максимальное значение элемента: %d\tВремя поиска: %d ms\n", maxSingle, elapsedSingle.Microseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	startTime = time.Now()
	maxMulti := maxChunks(data)
	elapsedMulti := time.Since(startTime)

	fmt.Printf("Максимальное значение элемента: %d\tВремя поиска: %d ms\n", maxMulti, elapsedMulti.Microseconds())
}
