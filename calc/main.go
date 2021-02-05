package main

import (
	"fmt"
)

func add(num1 int, num2 int)(res int){
	res = num1 + num2
	return res
}

func main(){
	result := add(2,2)
	fmt.Println("summed up to ", result)
}