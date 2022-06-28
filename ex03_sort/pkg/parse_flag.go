package pkg

import (
	"errors"
	"flag"
	"fmt"
)

type flags struct {
	n, r, u       bool
	k             int
	input, output string
}

func (f *flags) validate() error {
	if f.k <= 0 {
		return errors.New(fmt.Sprintf("invalid count at start of '%d'", f.k))
	}
	if f.input == "" || f.output == "" {
		return errors.New(fmt.Sprintf("invalid parameters for I/O files"))
	}
	return nil
}

func ParserFlags() (*flags, error) {
	f := &flags{}
	flag.BoolVar(&f.n, "n", false, "числовая сортировка, т.е. сравнение ведется по числовому значению")
	flag.BoolVar(&f.r, "r", false, "сортировка выполняется в обратном порядке (по убыванию)")
	flag.BoolVar(&f.u, "u", false, "не выводить повторяющие строки")
	flag.IntVar(&f.k, "k", 1, "указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умлчанию разделитель - пробел)")
	flag.Parse()
	f.input = flag.Arg(0)
	f.output = flag.Arg(1)
	err := f.validate()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("invalid number at field start: %s", err.Error()))
	}
	f.k--
	return f, nil
}

func (f *flags) GetN() bool {
	return f.n
}

func (f *flags) GetR() bool {
	return f.r
}

func (f *flags) GetU() bool {
	return f.u
}

func (f *flags) GetK() int {
	return f.k
}
