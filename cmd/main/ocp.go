package main

import "fmt"

type Product struct {
	name string
	size Size
	color Color
}

type Color int
type Size int

const (	
	small Size = iota
	medium
	large
)

const (
	red Color = iota
	green
	blue
)

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (spec ColorSpecification) IsSatisfied (p *Product) bool {
	return spec.color == p.color
}

type BetterFilter struct {}

func (*BetterFilter) Filter (products []Product, spec Specification) []*Product {
	result := []*Product {}
	for _, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &v)
		}
	}
	return result
}

func run_ocp() {
	product1 := Product{"product1", large, blue}
	product2 := Product{"product2", small, green}

	filter := BetterFilter{}

	result_products := filter.Filter([]Product{product1, product2}, ColorSpecification{green})

	fmt.Printf("Filtering products by color\n")
	for _, product := range result_products{
		fmt.Printf(" - %s\n", product.name)
	}
}


