package main

import "fmt"

func main() {
	data := []int{-1, 0, 3, 5, 9, 12}

	target := 13
	var s int
	s = seArch(data, target)
	fmt.Println(s)

}
func seArch(nums []int, target int) int {
	var i int = len(nums) - 1 //防止内存溢出
	var j int = 0
	var mid int = len(nums) / 2
	for i >= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			break

		} else if nums[mid] < target {
			j = mid + 1
		} else {
			i = mid - 1
		}

	}
	fmt.Println(mid)
	if nums[mid] != target {
		mid = -1
	}
	return mid
}
