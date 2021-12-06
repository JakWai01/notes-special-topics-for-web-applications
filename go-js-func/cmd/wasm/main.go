package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	num := js.Global().Call("add", 12, 24)
	fmt.Println(num.Int())

	s := js.Global().Call("hello")
	fmt.Println(s)

	env := js.Global().Get("env").String()
	fmt.Println(env)

	js.Global().Set("env", "CHANGED")

	envChanged := js.Global().Get("env").String()
	fmt.Println(envChanged)

	js.Global().Get("config").Set("key", "12345")

	objectChanged := js.Global().Get("config").Get("key").String()
	fmt.Println(objectChanged)
}
