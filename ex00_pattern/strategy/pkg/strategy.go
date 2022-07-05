package pkg

import (
	"fmt"
	"strings"
	"sort"
)

// Общий интерфейс алгоритмнов
type IStrategy interface {
	Process(string) string
}

// Метод, который обощает и вызывает необходимую стратегию
func RunStrategy(str string, op IStrategy) {
	fmt.Println(op.Process(str))
}

// Класс, который переводит в верхний регистр
type Upper struct{}
// Процесс-метод перевода в верхний регистр
func (Upper) Process(str string) string {
	return strings.ToUpper(str)
}
// Класс, который переводит в нижний регистр
type Lower struct{}
// Процесс-метод перевода в нижний регистр
func (Lower) Process(str string) string {
	return strings.ToLower(str)
}

// Класс, который сортирует элементы
type Sort struct{}

// Процесс-метод, который выполняет сортировку
 func (Sort) Process(str string) string {
	temp := []rune(str)
	sort.Slice(temp, func(i,j int) bool {
		return temp[i] < temp[j]
	})
	return string(temp)
 }
