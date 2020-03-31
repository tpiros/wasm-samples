package main

import (
	"strconv"
	"syscall/js"
)

func cube(this js.Value, inputs []js.Value) interface{} {
	number, _ := strconv.Atoi(inputs[0].String())
	result := number * number * number
	js.Global().Get("document").Call("getElementById", "result").Set("innerHTML", result)
	return nil
}

func main() {
	c := make(chan bool)
	js.Global().Set("cube", js.FuncOf(cube))
	<-c
}

// to create wasm, run: GOOS=js GOARCH=wasm go build -o cube.wasm
