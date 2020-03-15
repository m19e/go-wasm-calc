package main

import (
	"fmt"
	"syscall/js"
)

func print(this js.Value, i []js.Value) interface{} {
	fmt.Println(i[0])
	return nil
}

func registerCallbacks()  {
	js.Global().Set("print", js.FuncOf(print))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}