package main

import (
	"ex00_pattern/visitor/pkg"
)

/*
	посетитель - поведенческий паттерн проектирования, который позволяет добавлять в программу новые операции, не изменяю класссы объектов

	+:
	упрощает добавление операций
	объединяет родственные операции в одном классе
	посетитель может накапливать состояние при обходе структуры элементов

	-:
	паттерн не оправдан если иерархия часто меняется
	может привести к нарушению инкапсуляции
*/

func main() {
	circle := &pkg.Circle{}
	react := &pkg.Rectangle{}
	circle.Accept(&pkg.AreaCalculate{})
	circle.Accept(&pkg.MiddleCoordinates{})
	react.Accept(&pkg.AreaCalculate{})
	react.Accept(&pkg.MiddleCoordinates{})
}
