package main

import (
	"ex00_pattern/builder/pkg"
	"fmt"
)

func main() {
	a := pkg.NewPersonBuilder().Lives().
		Name("Make Take").
		Address("NY").
		Works().
		Company("Google").
		Position("Programmer").
		Salary(1111111).
		Build()
	fmt.Println(a)
}
