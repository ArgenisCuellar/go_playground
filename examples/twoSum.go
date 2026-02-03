package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, ok := numMap[complement]; ok { // Custom way to do an If Ok in this case is true is the value is in the map
			return []int{j, i}
		}
		numMap[num] = i
	}
	return nil
}

func main() {
	nums := []int{2, 11, 15, 7}
	target := 9
	fmt.Println(twoSum(nums, target))
}
