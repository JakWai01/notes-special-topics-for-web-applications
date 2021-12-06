package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Golang Add Example")
	js.Global().Set("add", addWrapper())
	select {}
}

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
