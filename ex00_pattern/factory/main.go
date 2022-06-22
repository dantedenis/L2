package main

import (
	"ex00_pattern/factory/pkg"
	"fmt"
)

func main() {
	person, err := pkg.GetPerson("leks")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%T", person)
	}
}
