package main

import "fmt"

func twosum(nums []int, target int) []int {
	l := len(nums)

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twosum(nums, target)
	fmt.Println(result)
}
