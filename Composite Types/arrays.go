package main

import "fmt"

func main() {
	var arr [3]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30

	num := [6]int{1,2,3,4,5,6}

	numbers := [...]int{1,2,3,4,5,6,7,8,9}

	for i, v := range numbers {
		fmt.Println("index:",i, "value:", v)
	}

	fmt.Println("arr:",arr)
	fmt.Println("Element 1:", arr[0])
	fmt.Println("num:", num)
	fmt.Println("numbers:", numbers)
	
}