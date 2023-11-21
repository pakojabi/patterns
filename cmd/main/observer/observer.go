package observer

import (
	"container/list"
	"fmt"
)

type Observable struct {
	subs *list.List
}

type Observer interface {
	Notify(d interface {})
}

type Person struct {
	Observable
	Name string
}

func NewPerson(name string) *Person {
	return &Person{Name: name, Observable: Observable{subs: list.New()}}
}

type DoctorService struct {
	DoctorName string
}

func NewDoctorService(doctorName string) *DoctorService {
	return &DoctorService{DoctorName: doctorName}
}

func (o *Observable) Subscribe(obs Observer) {
	o.subs.PushBack(obs)
}

func (o *Observable) Unsubscribe(obs Observer) {
	for i := o.subs.Front(); i != nil; i = i.Next() {
		if i.Value == obs {
			o.subs.Remove(i)
		}
	}
}

func (o *Observable) Fire(d interface {}) {
	for i := o.subs.Front(); i != nil; i = i.Next() {
		i.Value.(Observer).Notify(d)
	}
}

func (p *Person) GotSick() {
	p.Fire(p.Name)
}

func (ds *DoctorService) Notify(d interface{}) {
	fmt.Println(ds.DoctorName, "was called by", d.(string))
}

func RunObserver() {
	person1 := NewPerson("manolo")
	person2 := NewPerson("pepe")

	doctorService := NewDoctorService("matasanos")

	person1.Subscribe(doctorService)
	person2.Subscribe(doctorService)

	person1.GotSick()
	person2.GotSick()
}





