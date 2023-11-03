package main

import (
	"github.com/pakojabi/patterns/cmd/main/bridge"
	"github.com/pakojabi/patterns/cmd/main/composite"
	"github.com/pakojabi/patterns/cmd/main/decorator"
	"github.com/pakojabi/patterns/cmd/main/factory"
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
}
