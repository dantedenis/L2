package grep

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type grep struct {
	rule *flags
	text map[int]pair
}

type pair struct {
	line string
	mark bool
}

func NewGrep() *grep {
	var gr grep
	gr.rule = NewFlags()
	gr.text = make(map[int]pair)

	return &gr
}

func (g *grep) Run() {
	scan := bufio.NewScanner(os.Stdin)

	for i := 0; scan.Scan(); i++ {
		g.text[i] = g.makePair(scan.Text())
	}
}

func (g *grep) makePair(line string) pair {
	var res pair

	res.line = line
	res.mark = !strings.Contains(line, g.rule.pattern)
	if !g.rule.fix {
		res.mark = !res.mark
	}

	return res
}

// TODO: допилить реализацию, и выполнить промежутки множества в одно целое

func (g *grep) PrintMap() {
	fmt.Println(g.text)
}
