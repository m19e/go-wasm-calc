package main

import (
	"strconv"
	"syscall/js"
)

var (
	currentNum  = ""
	accumulator = ""
	operator = None
)

type Operator int

const (
	Plus Operator = iota
	Sub
	Mul
	Div
	None
)

func inputNum(this js.Value, i []js.Value) interface{} {
	currentNum += i[0].String()
	js.Global().Get("document").Call("getElementById", "result").Set("textContent", currentNum)
	return nil
}

func clearNum(this js.Value, i []js.Value) interface{} {
	currentNum = ""
	accumulator = ""
	js.Global().Get("document").Call("getElementById", "result").Set("textContent", "0")
	return nil
}

func doPlus(this js.Value, i []js.Value) interface{} {
	accumulator += currentNum
	currentNum = ""
	js.Global().Get("document").Call("getElementById", "result").Set("textContent", "0")
	return nil
}

func doOperate(this js.Value, i []js.Value) interface{} {
	accumulator += currentNum
	currentNum = ""
	js.Global().Get("document").Call("getElementById", "result").Set("textContent", "0")

	switch i[0].String() {
	case "+":
		operator = Plus
	case "-":
		operator = Sub
	case "*":
		operator = Mul
	case "/":
		operator = Div
	default:
		operator = None
	}

	return nil
}

func doEqual(this js.Value, i []js.Value) interface{} {
	int1, _ := strconv.Atoi(currentNum)
	int2, _ := strconv.Atoi(accumulator)

	accumulator = ""

	var result int
	switch operator {
	case Plus:
		result = int1 + int2
	case Sub:
		result = int2 - int1
	case Mul:
		result = int1 * int2
	case Div:
		result = int2 / int1
	}

	js.Global().Get("document").Call("getElementById", "result").Set("textContent", result)
	currentNum = strconv.Itoa(result)
	return nil
}

func registerCallbacks() {
	js.Global().Set("inputNum", js.FuncOf(inputNum))
	js.Global().Set("clearNum", js.FuncOf(clearNum))
	js.Global().Set("doOperate", js.FuncOf(doOperate))
	js.Global().Set("doEqual", js.FuncOf(doEqual))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
