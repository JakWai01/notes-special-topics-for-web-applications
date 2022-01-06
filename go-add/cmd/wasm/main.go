package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Golang Add Example")
	// Set property and value in the global context
	js.Global().Set("add", addWrapper())
	select {}
}

// Add function
func add(x int, y int) int {
	return x + y
}

// Wrapper
func addWrapper() js.Func {
	// Define function, which is callable from JS
	addFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			return "Invalid no of arguments passed"
		}
		x := args[0].Int()
		y := args[1].Int()
		// Call the add function
		result := add(x, y)
		return result
	})
	// addWrapper returns the function, which is callable from JS
	return addFunc
}
