package main
import "github.com/pakojabi/patterns/cmd/main/factory"

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
}
