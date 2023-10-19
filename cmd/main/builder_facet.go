package main

import "fmt"

type BPerson struct {
	Address, City, Postcode string
	Company, Position string
	AnnualIncome int
}

type PersonBuilder struct {
	person *BPerson
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&BPerson{}}
}

func (pb *PersonBuilder) Build() *BPerson {
	return pb.person
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (pb *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*pb}
}

func (pb *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*pb}
}

func (it *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	it.person.Address = streetAddress
	return it
}
  
func (it *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	it.person.City = city
	return it
}
  
func (it *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	it.person.Postcode = postcode
	return it
}

func (pjb *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
  pjb.person.Company = companyName
  return pjb
}

func (pjb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
  pjb.person.Position = position
  return pjb
}

func (pjb *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
  pjb.person.AnnualIncome = annualIncome
  return pjb
}

func RunBuilderFacet() {
	person := NewPersonBuilder().
		Lives().
			At("address").
			In("City").
			WithPostcode("123").
		Works().
			At("company").
			AsA("position").
			Earning(20000).Build()
	fmt.Println(*person)
}




