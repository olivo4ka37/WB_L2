package main

import (
	"log"
	"sort"
)

//Найти три максимума (112, 243, 567)

func main() {
	ints := []int{3, 4, 64, 12, 45, 1, 0, 35, 243, 2, -4, 112, -12, 567}
	mxs := detectMxs(ints)
	log.Println(mxs)
}

func detectMxs(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	result := nums[:3]

	return result
}
