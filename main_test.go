package main

import (
	"testing"
)

// Тесты для generateRandomElements
func TestGenerateRandomElements(t *testing.T) {
	// Тест на размер 0 (крайний случай)
	t.Run("size 0", func(t *testing.T) {
		res := generateRandomElements(0)
		if len(res) != 0 {
			t.Errorf("Ожидалась длина 0, получена %d", len(res))
		}
	})

	// Тест на отрицательный размер (крайний случай)
	t.Run("negative size", func(t *testing.T) {
		res := generateRandomElements(-5)
		if len(res) != 0 {
			t.Errorf("Ожидалась длина 0, получена %d", len(res))
		}
	})

	// Тест на нормальный размер
	t.Run("normal size", func(t *testing.T) {
		size := 100
		res := generateRandomElements(size)
		if len(res) != size {
			t.Errorf("Ожидалась длина %d, получена %d", size, len(res))
		}
		// Проверяем, что числа положительные, согласно ТЗ
		for _, v := range res {
			if v <= 0 {
				t.Errorf("Ожидалось положительное число, получено %d", v)
			}
		}
	})
}

// Тесты для maximum
func TestMaximum(t *testing.T) {
	// Крайний случай: пустой слайс
	t.Run("empty slice", func(t *testing.T) {
		res := maximum([]int{})
		if res != 0 {
			t.Errorf("Ожидалось 0, получено %d", res)
		}
	})

	// Крайний случай: один элемент
	t.Run("single element", func(t *testing.T) {
		res := maximum([]int{42})
		if res != 42 {
			t.Errorf("Ожидалось 42, получено %d", res)
		}
	})

	// Нормальный слайс
	t.Run("multiple elements", func(t *testing.T) {
		res := maximum([]int{5, 10, 3, 99, 1})
		if res != 99 {
			t.Errorf("Ожидалось 99, получено %d", res)
		}
	})

	// Слайс с дубликатами максимумов
	t.Run("duplicate maximums", func(t *testing.T) {
		res := maximum([]int{5, 10, 10, 5})
		if res != 10 {
			t.Errorf("Ожидалось 10, получено %d", res)
		}
	})
}
