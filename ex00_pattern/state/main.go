package main

import (
	"ex00_pattern/state/pkg"
)

/*
	Состояние - поведенческий паттерн, позволяет менять поведение в зависимости от внутреннего состояние

	+:
	избавляет от множества условных операторов
	упрощает код контекста

	-:
	может усложжнить код если состояний много и редко меняются
*/

func main() {
	sc := pkg.NewStateContext()
	sc.Heat()
	sc.Heat()
	sc.Heat()
	sc.Freeze()
	sc.Freeze()
	sc.Freeze()
}
