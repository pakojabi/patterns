package main

import (
	"github.com/pakojabi/patterns/cmd/main/bridge"
	chainofresposibility "github.com/pakojabi/patterns/cmd/main/chain_of_resposibility"
	"github.com/pakojabi/patterns/cmd/main/composite"
	"github.com/pakojabi/patterns/cmd/main/decorator"
	"github.com/pakojabi/patterns/cmd/main/factory"
	"github.com/pakojabi/patterns/cmd/main/observer"
	"github.com/pakojabi/patterns/cmd/main/state"
	"github.com/pakojabi/patterns/cmd/main/visitor"
)

func main() {
	println("OCP")
	RunOcp()
	println()
	println("DIP")
	RunDip()
	println()
	println("Builder")
	RunBuilder()
	println()
	println("BuilderFacet")
	RunBuilderFacet()
	println()
	println()
	println("Factory Generator")
	factory.RunFactoryGenerator()
	println()
	println("Prototype Factory")
	factory.RunPrototypeFactory()
	println()
	println("Bridge")
	bridge.RunBridge()
	println()
	println("Composite - Neural network")
	composite.RunNeurons()
	println()
	println("Decorator")
	decorator.RunDecorator()
	println()
	println("Chain of Responsibility - Method chain")
	chainofresposibility.RunMethodChain()
	println()
	println("Chain of Responsibility - Broker chain")
	chainofresposibility.RunBrokerChain()
	println()
	println("Observer")
	observer.RunObserver()
	println()
	println("State (Classic)")
	state.RunClassicState()
	println()
	// println("State (Handmade)")
	// state.RunHandmadeState()
	// println()
	println("Visitor (Classic)")
	visitor.RunClassicVisitor()
	println()
}
