package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	// Call JS add function with parameters 12 and 24 and print the result
	sum := js.Global().Call("add", 12, 24)
	fmt.Println(sum.Int())

	// Call JS hello function and print the result
	message := js.Global().Call("hello")
	fmt.Println(message.String())

	// Get the value of the variable name
	name := js.Global().Get("name")
	fmt.Println(name.String())

	// Change the value of the variable
	js.Global().Set("name", "Lord Voldemort")

	// Get the value of the variable name again
	name = js.Global().Get("name")
	fmt.Println(name.String())

	// Get the object and set a key and value
	js.Global().Get("obj").Set("Horkrux", "Nagini")

	// Get the value of the Horkrux key
	object := js.Global().Get("obj").Get("Horkrux")
	fmt.Println(object.String())
}
