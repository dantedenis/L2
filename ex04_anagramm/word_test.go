package ex04_anagramm

import (
	"fmt"
	"testing"
)

func TestWord(t *testing.T) {
	set := NewSet()

	words := []string{"спал", "Упал", "ПрОпал", "пятак", "пятак", "пятак",
		"пЯтка", "яткап", "тяпка", "листок", "слиток", "столик",
		"кочка", "кочан", "чакоч", "ночак", "ночак"}

	set.Add(&words)
	m := *set.ToMap()
	for k, v := range m {
		fmt.Println(k, *v)
	}
}
