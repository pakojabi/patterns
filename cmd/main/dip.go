package main

import "fmt"

// Dependency inversion principle


type Relationship int
const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from *Person
	to *Person
	relationship Relationship
}

type Relationships struct {
	relations []Info
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for _, relation := range r.relations {
		if relation.relationship == Child && relation.to.name == name {
			result = append(result, relation.from) 
		}
	}
	return result
}

func (r *Relationships) AddParentAndChild(parent, child *Person){
	r.relations = append(r.relations, Info{parent, child, Parent})
	r.relations = append(r.relations, Info{child, parent, Child})
}

type Research struct {
	browser RelationshipBrowser
}

func (r *Research) Investigate() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called ", p.name)
	}
}

func run_dip() {
	parent := &Person{"John"}
	child1 := &Person{"Chris"}
	child2 := &Person{"Matt"}

	relationships := &Relationships{}
	relationships.AddParentAndChild(parent, child1)
	relationships.AddParentAndChild(parent, child2)

	research := Research{relationships}

	research.Investigate()
}