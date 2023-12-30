package main

import (
	"fmt"
	"syscall/js"
)

func count() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		var totalCount = 0
		for i := 0; i < 500; i++ {
			fmt.Println(i)
			totalCount++
		}
		fmt.Printf("toal count %d\n", totalCount)
		return totalCount
	})
}

func write(ch chan int, number *int) {

	for i := 0; i < 10; i++ {
		ch <- i
		*number++
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func process() js.Func {
	number := 0

	return js.FuncOf(func(this js.Value, args []js.Value) any {
		// creates capacity of 2
		ch := make(chan int, 1)
		go write(ch, &number)
		for v := range ch {
			fmt.Println("reading this", v, "to ch")
		}
		return number
	})
}

func main() {
	fmt.Println("Web Assembly")
	js.Global().Set("process", process())
	<-make(chan bool)
}
