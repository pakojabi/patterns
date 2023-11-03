package composite

import "fmt"

type Neuron struct {
	In, Out []*Neuron
}

func (left *Neuron) ConnectTo(right *Neuron) {
	left.Out = append(left.Out, right)
	right.In = append(right.In, left)
}

type NeuronLayer struct {
	Neurons []Neuron
}

func NewNeuronLayer(count int) *NeuronLayer {
	return &NeuronLayer{make([]Neuron, count)}
}

type NeuronInterface interface {
	List() []*Neuron
}

func (n *Neuron) List() []*Neuron {
	return []*Neuron{n}
}

func (l *NeuronLayer) List() []*Neuron {
	result := make([]*Neuron, 0)
	for i := range l.Neurons {
		result = append(result, &l.Neurons[i])
	}
	return result
}

func Connect(left, right NeuronInterface) {
	for _, i := range left.List() {
		for _, j := range right.List() {
			i.ConnectTo(j)
		}
	}
}

func RunNeurons() {
	n1, n2 := &Neuron{}, &Neuron{}
	l1 := NewNeuronLayer(2)
	Connect(n1, n2)
	Connect(n2, l1)

	fmt.Println("n1: ", n1)
	fmt.Println("n2: ", n2)
	fmt.Println("l1: ", l1)
}
