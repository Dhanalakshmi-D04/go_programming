package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5}
	fmt.Printf("Array elements : %v \n", arr)

	fmt.Printf("Element at index 2: %d \n", arr[2])

	arr[2] = 6
	fmt.Printf("After modifying element at index 2: %v", arr)

	index := 2
	value := 7

	arr = append(arr[:index], append([]int{value}, arr[index:]...)...)
	fmt.Printf("After inserting 7 at index 2 : %v", arr)

	index = 2
	arr = append(arr[:index], arr[index+1:]...)
	fmt.Printf("After deleting element at index 2: %v", arr)

}
