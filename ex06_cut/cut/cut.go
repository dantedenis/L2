package cut

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/slices"
	"io"
	"os"
	"strings"
)

// Используем композицию
type cut struct {
	*flags
}

func NewCut() *cut {
	return &cut{
		// в структуру встроены только флаги парсим их..
		ParseFlags(),
	}
}

func (c *cut) Run() (err error) {
	var input io.Reader

	input = os.Stdin
	// если нужно читать с файла перезаписываем input
	if c.input != "" {
		file, err := os.Open(c.input)
		if err != nil {
			return err
		}
		// добавляем отложенную функцию, тк нужно закрыть дескриптор и проверяем на ошибку
		defer func(file *os.File) {
			err = file.Close()
		}(file)
		// перезаписываем интерфейс input (Reader)
		input = file
	}

	// инициализируем скинер с помощью которого будет считывать байты из io.Reader
	scanner := bufio.NewScanner(input)
	// сканируем до EOF
	for scanner.Scan() {
		str := scanner.Text()
		// сплитим по разделителю инициализируем слайс для вывода на STDOUT
		words, prepare := strings.Split(str, c.d), make([]string, 0)
		for i, w := range words {
			// выполняем фильтр колонок, если нужно
			if c.targetRow != nil && slices.Contains(c.targetRow, i+1) {
				prepare = append(prepare, w)
				// иначе каждое слово добавляем в слайс
			} else if c.targetRow == nil {
				prepare = append(prepare, w)
			}
		}
		// вывод полученной строки
		fmt.Println(strings.Join(prepare, " "))
	}
	return nil
}
