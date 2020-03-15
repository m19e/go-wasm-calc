package main

import (
	"fmt"
	"syscall/js"
)

func print(this js.Value, i []js.Value) interface{} {
	fmt.Println(i[0])
	return nil
}

func manupilateDom(this js.Value, i []js.Value) interface{}  {
	js.Global().Get("document").Call("getElementById", "result").Set("textContent", "WebAssembly!")
	return nil
}

func registerCallbacks()  {
	js.Global().Set("print", js.FuncOf(print))
	js.Global().Set("manupilateDom", js.FuncOf(manupilateDom))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}