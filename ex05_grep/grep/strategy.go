package grep

import (
	"fmt"
)

type IStrategy interface {
	Processing(*grep)
}

// Необходим был для дебага
type PrintPair struct{}

func (pr PrintPair) Processing(g *grep) {
	m := g.text
	for i := 0; i < len(m); i++ {
		fmt.Printf("%5d: %6v - [%s]\n", i, m[i].mark, m[i].line)
	}
}

// дефолтный принт всех совпадений
type Print struct{}

func (p Print) Processing(g *grep) {
	for i := 0; i < len(g.text); i++ {
		if g.text[i].mark {
			if g.rule.num {
				fmt.Printf("%d:", i)
			}
			fmt.Println(g.text[i].line)
		}
	}
}

// просто печает количество совпадений
type PrintCount struct{}

func (pr PrintCount) Processing(g *grep) {
	count := 0
	for _, v := range g.text {
		if v.mark {
			count++
		}
	}
	fmt.Println(count)
}

// необходимо для выполнения ремарков, по причине возможных интервалов вывода строк
type Remarking struct{}

func (r Remarking) Processing(g *grep) {
	indexing := make([]int, 0)

	// собираем слайс нидексов строк совпадений
	for i := 0; i < len(g.text); i++ {
		if g.text[i].mark {
			indexing = append(indexing, i)
		}
	}

	before, after := g.rule.bf, g.rule.aft
	if g.rule.ctx > before {
		before = g.rule.ctx
	}
	if g.rule.ctx > after {
		after = g.rule.ctx
	}

	// для каждого индекса пересчитав стартовыый индекс строки и конечный маркируем дополнительные строки
	for _, ind := range indexing {
		start, end := ind-before, ind+after
		if start < 0 {
			start = 0
		}
		if end >= len(g.text) {
			end = len(g.text) - 1
		}
		for start <= end {
			g.text[start].mark = true
			start++
		}
	}

}

// реверс маркированных строк
type Reverse struct{}

func (r Reverse) Processing(g *grep) {
	for _, v := range g.text {
		v.mark = !v.mark
	}
}
