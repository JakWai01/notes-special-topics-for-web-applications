package main

import (
	"fmt"
	"syscall/js"
)

// Declare a main function, this is the entrypoint into our go module
// That will be run. In our example, we won't need this
func main() {
	fmt.Println("Golang Add Example")
	js.Global().Set("add", addWrapper())
	select {}
}

// This exports an add function.
// It takes in two 32-bit integer values
// And returns a 32-bit integer value.
// To make this function callable from JavaScript,
// we need to add the: "export add" comment above the function
//export add
func add(x int, y int) int {
	return x + y
}

func addWrapper() js.Func {
	addFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			return "Invalid no of arguments passed"
		}
		x := args[0].Int()
		y := args[1].Int()
		result := add(x, y)
		return result
	})
	return addFunc
}
