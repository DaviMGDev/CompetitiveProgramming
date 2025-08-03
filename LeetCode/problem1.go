package main

import "fmt"

func twoSum(nums []int, target int) []int {
	size := len(nums)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	result := twoSum([]int{0, 1, 2, 3, 4}, 3) 
	if result != nil {
		fmt.Println(result)
	}
}
