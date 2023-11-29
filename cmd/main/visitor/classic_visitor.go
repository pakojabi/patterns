package visitor

import (
	"fmt"
	"strings"
)

type ExpressionVisitor interface {
	VisitDoubleExpresion(de *DoubleExpression)
	VisitAdditionExpression(ae *AdditionExpression)
}

type Expression interface {
	Accept(ev ExpressionVisitor)
}

type DoubleExpression struct {
	value float64
}

func (de *DoubleExpression) Accept(ev ExpressionVisitor) {
	ev.VisitDoubleExpresion(de)
}

type AdditionExpression struct {
	left, right Expression
}
func (ae *AdditionExpression) Accept(ev ExpressionVisitor) {
	ev.VisitAdditionExpression(ae)
}

type Printer struct {
	sb strings.Builder
}
func (pr *Printer) VisitDoubleExpresion(de *DoubleExpression) {
	pr.sb.WriteString(fmt.Sprintf("%g", de.value))
}

func (pr *Printer) VisitAdditionExpression(ae *AdditionExpression) {
	pr.sb.WriteRune('(')
	ae.left.Accept(pr)
	pr.sb.WriteRune('+')
	ae.right.Accept(pr)
	pr.sb.WriteRune(')')
}

func (pr *Printer) String() string {
	return pr.sb.String()
}

type Calculator struct {
	result float64
}

func (c *Calculator) Result() float64 {
	return c.result
}

func (c *Calculator) VisitDoubleExpresion(de *DoubleExpression) {
	c.result = de.value
}

func (c *Calculator) VisitAdditionExpression(ae *AdditionExpression) {
	ae.left.Accept(c)
	result := c.result
	ae.right.Accept(c)
	result += c.result
	c.result = result
}

func (c *Calculator) String() string {
	return fmt.Sprintf("%g", c.result)
}



func RunClassicVisitor() {
	// 1 + (2+3)
	expression := &AdditionExpression{
		&DoubleExpression{1},
		&AdditionExpression{
			&DoubleExpression{2},
			&DoubleExpression{3},
		},
	}

	pr := &Printer{}
	expression.Accept(pr)
	c := &Calculator{}
	expression.Accept(c)

	fmt.Printf("%s = %s\n", pr, c)
}








