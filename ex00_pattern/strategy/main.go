package main

import (
	"ex00_pattern/strategy/pkg"
)

/*
	Стратегия - поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов
	и помещает каждый из них в собственный класс, после чего их можно зваимоизменять в рантайме
	
	
*/
func main() {
	str := "HellO wORld"
	pkg.RunStrategy(str, pkg.Upper{})
	pkg.RunStrategy(str, pkg.Sort{})
	pkg.RunStrategy(str, pkg.Lower{})
}