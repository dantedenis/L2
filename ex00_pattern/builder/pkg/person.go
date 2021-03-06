package pkg

import (
	"fmt"
)

type Person struct {
	//personal data
	name    string
	address string

	//job data
	company  string
	position string
	salary   int
}

func (p *Person) String() string {
	return fmt.Sprintf("Personal\n\tName: %s\n\tAddress: %s\nJob\n\tCompany: %s\n\tPosition: %s\n\tSalary: %d\n",
		p.name, p.address, p.company, p.position, p.salary)
}

type PersonBuilder struct {
	person *Person
}

// PersonalDetails Делим наш объект на 2 половины, для более простой инициализации (разделяй и властвуй)
type PersonalDetails struct {
	PersonBuilder
}

type JobDetails struct {
	PersonBuilder
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{person: &Person{}}
}

// И для каждой из половин реализуем сеттеры полей, тем самым отправляю в объект необходимые нам значения
func (p *PersonBuilder) Lives() *PersonalDetails {
	return &PersonalDetails{*p}
}

func (p *PersonBuilder) Works() *JobDetails {
	return &JobDetails{*p}
}

func (p *PersonalDetails) Name(name string) *PersonalDetails {
	p.person.name = name
	return p
}

func (p *PersonalDetails) Address(address string) *PersonalDetails {
	p.person.address = address
	return p
}

func (j *JobDetails) Company(company string) *JobDetails {
	j.person.company = company
	return j
}

func (j *JobDetails) Position(pos string) *JobDetails {
	j.person.position = pos
	return j
}

func (j *JobDetails) Salary(salary int) *JobDetails {
	j.person.salary = salary
	return j
}

// и в конечном итоге возвращаем наш заполненый объект
func (p *PersonBuilder) Build() *Person {
	return p.person
}
