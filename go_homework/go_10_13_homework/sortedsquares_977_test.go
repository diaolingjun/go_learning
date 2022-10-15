package go_10_13_homework

import (
	"reflect"
	"sort"
	"testing"
)

// 算法函数
func sortedSquares(nums []int) []int {
	var numsnew = make([]int, len(nums)) //创建一个动态数组存储平方数组
	for key, val := range nums {
		numsnew[key] = val * val
	}
	sort.Ints(numsnew) //调用sort排序函数
	return numsnew
}

// 测试函数
func TestSortedSquares(t *testing.T) {
	data := []struct {
		nums   []int
		answer []int
	}{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-4, -1, 0, 1, 10}, []int{0, 1, 1, 16, 100}},
		{[]int{11, -1, 0, 3, 10}, []int{0, 1, 9, 100, 121}},
	}
	for _, val := range data {
		result := sortedSquares(val.nums)
		if !reflect.DeepEqual(val.answer, result) { //判断切片是否相等
			t.Fatalf("expect:[%v] != result[%v]", result, val.answer)
		}
	}
}
