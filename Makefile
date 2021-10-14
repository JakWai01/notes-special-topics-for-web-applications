go-example: 
	GOOS=js GOARCH=wasm go build -o go/assets/app.wasm go/cmd/wasm/main.go
	go run go/cmd/server/main.go