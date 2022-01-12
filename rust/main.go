package main

import (
	"fmt"
	"net/http"
)

// cargo install wasm-pack
// wasm-pack build --release --target web

func main() {
	err := http.ListenAndServe(":9090", http.FileServer(http.Dir("../rust")))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
