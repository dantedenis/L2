package grep

import (
	"bufio"
	_ "fmt"
	"os"
	"strings"
)

// компинируем флаги и строки
type grep struct {
	rule *flags
	text map[int]*pair
}

// каждую строку помечаем флагом
type pair struct {
	line string
	mark bool
}

func NewGrep() *grep {
	var gr grep
	gr.rule = NewFlags()
	gr.text = make(map[int]*pair)

	return &gr
}

// Основной движок выполнен с помощью стратеги
func (g *grep) Run() {
	scan := bufio.NewScanner(os.Stdin)

	//
	for i := 0; scan.Scan(); i++ {
		g.text[i] = g.makePair(scan.Text())
	}

	if g.rule.aft > 0 || g.rule.bf > 0 || g.rule.ctx > 0 {
		g.IterMap(Remarking{})
	}

	if g.rule.invert {
		g.IterMap(Reverse{})
	}

	switch {
	case g.rule.count:
		g.IterMap(PrintCount{})
	default:
		g.IterMap(Print{})
	}
}

// сохраняем и маркируем строки
func (g *grep) makePair(line string) *pair {
	var res pair

	// на основе поднятого флага о необходимости игнорировании регистра
	res.line = line
	if g.rule.ignore {
		res.mark = strings.Contains(strings.ToLower(line), strings.ToLower(g.rule.pattern))
	} else {
		res.mark = strings.Contains(line, g.rule.pattern)
	}

	return &res
}

// Метод вызывает необходимую стратегию для структуры
func (g *grep) IterMap(s IStrategy) {
	s.Processing(g)
}
