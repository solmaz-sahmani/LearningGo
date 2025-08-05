package main

import "fmt"

func main() {
	
	var s []int
	fmt.Println("Is this array empty?", s == nil)
	
	s1 := []int{1,2,3}
	fmt.Println("s1 :", s1)
	fmt.Println("s1 Length:", len(s1))
	fmt.Println("s1 Capacity:", cap(s1))

	s2 := []int{}
	s2 = append(s1,4)
	s2 = append(s2,5,6,7)
	fmt.Println("s2 :", s2)

	nums := make([]int, 3,6) //length=3, capacity=6
	fmt.Println("nums :", nums, len(nums), cap(nums))
	
	numbers := []int{1,2,3,4,5,6}
	numbers = nil
	fmt.Println("numbers array become empty by nil :", numbers)

	number := []int{1,2,3,4,5,6}
	number = number[:0]
	fmt.Println("number array become empty by using number[:0] :", number)

	arr := []int{0,1,2,3,4,5,6,7,8}
	sub := arr[1:4]
	fmt.Println("slicing arr from index 1 to 3 :", sub)

	src := []int {1,2,3}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println("now src copied into dst :", dst)

	array := [3]int{1,2,3}
	slice := array[:]
	fmt.Println("convert array to slice :", slice)
}