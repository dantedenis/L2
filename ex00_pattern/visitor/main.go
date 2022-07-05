package main

import (
	"ex00_pattern/visitor/pkg"
)

func main() {
	circle := &pkg.Circle{}
	react := &pkg.Rectangle{}
	circle.Accept(&pkg.AreaCalculate{})
	circle.Accept(&pkg.MiddleCoordinates{})
	react.Accept(&pkg.AreaCalculate{})
	react.Accept(&pkg.MiddleCoordinates{})
}
