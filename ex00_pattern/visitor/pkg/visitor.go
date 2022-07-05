package pkg

import (
	"fmt"
)

type Shape interface {
	GetType() string
	Accept(Visitor)
}

type Circle struct {
	radius int
}

func (c *Circle) GetType() string {
	return "Circle"
}

func (c *Circle) Accept(v Visitor) {
	v.VisitForCircle(c)
}

type Rectangle struct {
	height int
	length int
}

func (r *Rectangle) GetType() string {
	return "Rectangle"
}

func (r *Rectangle) Accept(v Visitor) {
	v.VisitForReactangle(r)
}

type Visitor interface {
	VisitForReactangle(*Rectangle)
	VisitForCircle(*Circle)
}

type AreaCalculate struct{}

func (a *AreaCalculate) VisitForCircle(c *Circle) {
	fmt.Println("Calculate area for", c.GetType())
}

func (a *AreaCalculate) VisitForReactangle(r *Rectangle) {
	fmt.Println("Calculate area for", r.GetType())
}

type MiddleCoordinates struct {}

func (m *MiddleCoordinates) VisitForCircle(c *Circle) {
	fmt.Println("Calculate middle cordinates for", c.GetType())
}

func (m *MiddleCoordinates) VisitForReactangle(c *Rectangle) {
	fmt.Println("Calculate middle cordinates for", c.GetType())
}