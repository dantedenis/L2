package ex04_anagramm

// переопределяем слайс рун, что бы имплементировать интерфейс сорт
type sliceRune []rune

func (b sliceRune) Len() int {
	return len(b)
}

func (b sliceRune) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b sliceRune) Less(i, j int) bool {
	return b[i] < b[j]
}

func StringToRune(str string) []rune {
	var r []rune
	for _, runeVal := range str {
		r = append(r, runeVal)
	}
	return r
}
