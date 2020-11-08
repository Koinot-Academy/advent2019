package main

import "fmt"

func main() {
	module := Module{12, 0}

	module.RequiredFuel = module.GetRequiredFuel()
	fmt.Println(module)
}
