package main

import "fmt"

var ch chan int = make(chan int)
var complete chan int = make(chan int)
func foo(){
	ch <- 0
}

func loop(){
	for i := 0; i < 10; i++{
		fmt.Println("%d ", i)
	}

	complete <- 0
}
func main(){
	go loop()
	<- complete
}
