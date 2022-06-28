package pkg

import (
	"errors"
	"strings"
)

type IPerson interface {
	GetName() string
	GetPosition() string
	GetSalary() int
}

type person struct {
	name     string
	position string
	salary   int
}

// Проверка, что персон имплементирует интерфейс
var _ IPerson = &person{}

func (p *person) GetName() string {
	return p.name
}

func (p *person) GetPosition() string {
	return p.position
}

func (p *person) GetSalary() int {
	return p.salary
}

type Director struct {
	person
}

type TeamLead struct {
	person
}

type ProgrammerGo struct {
	person
}

// GetPerson порождающий объекты метод, центр самого паттерна
func GetPerson(worker string) (IPerson, error) {
	switch strings.ToLower(worker) {
	case "director":
		return &Director{}, nil
	case "teamlead":
		return &TeamLead{}, nil
	case "programmergo":
		return &ProgrammerGo{}, nil
	}
	return nil, errors.New("unknown type person")
}
