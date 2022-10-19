package go1010homework

import (
	"testing"
)

// go test -run ^Testsearch$ ./go1010homework
func TestSearch(t *testing.T) {
	data := []struct {
		nums   []int
		target int
		answer int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 4, -1},
		{[]int{-1, 0, 3, 5, 9, 12}, 3, 2},
		{[]int{-1, 0, 3, 5, 9, 12}, 12, 5},
	}
	for _, v := range data {
		x := Search(v.nums, v.target) //获取返回值
		if x != v.answer {
			t.Fatalf("expect:[%v] != result[%v]", x, v.answer)
		}

	}
}
func Search(nums []int, target int) int {
	var i int = len(nums) - 1
	var j int = 0
	var mid int = len(nums) / 2
	for i >= j {
		mid = (i + j) / 2
		if nums[mid] == target {
			return mid //符合条件，查找成功跳出循环，减少时间消耗

		}
		if nums[mid] < target {
			j = mid + 1
		}
		if nums[mid] > target {
			i = mid - 1
		}

	}
	//fmt.Println(mid)
	if nums[mid] != target { //查找到最后找不到就把结果赋值为-1返回出去
		mid = -1
	}
	return mid
}
