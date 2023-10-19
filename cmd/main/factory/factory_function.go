package factory

import "fmt"

// constructor

type Person struct {
	Name string
	NumEyes int
}

func NewPerson(name string) *Person {
	return &Person{Name: name, NumEyes: 2}
}

/// interface factory

type Person2 interface {
	SayHello()
}

type personPrivate struct {
	name string
	numEyes int
}

func (p *personPrivate) SayHello() {
	fmt.Println("Hello ", p.name)
}

func NewPerson2(name string) Person2 {
	return &personPrivate{name, 2}
}