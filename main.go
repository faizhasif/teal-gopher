package main

import (
	"fmt"
	"teal-gopher/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	fmt.Println("Hello!")
}
