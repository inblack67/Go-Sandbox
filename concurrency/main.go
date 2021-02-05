package main

import (
	"fmt"
	"time"
)

func loopOver(value int){
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main(){
	fmt.Println("go routines")
	// loopOver(10)

	// async
	go loopOver(10)

	// halt main
	fmt.Scanln()
}