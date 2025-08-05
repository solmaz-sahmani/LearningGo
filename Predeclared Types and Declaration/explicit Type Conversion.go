package main

import "fmt"

func main() {
	var x int = 10
	var y float32 = 3.5

	result := float32(x) + y
//  result := x + y           // error
	fmt.Println(result)
}