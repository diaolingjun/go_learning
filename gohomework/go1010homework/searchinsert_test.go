package go1010homework

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {

	data := []struct {
		nums   []int
		target int
		answer int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 20, 4},
		{[]int{1, 3, 5, 6}, 1, 0},
	}
	for _, v := range data {
		result := SearchInsert(v.nums, v.target)
		if result != v.answer {
			t.Fatalf("expect:[%v] != result[%v]", result, v.answer)
		}
	}

}

func SearchInsert(nums []int, target int) int { //和内部函数重复，改名后提交通过
	var i int = 1
	for key, value := range nums {
		//fmt.Println(key, value, len(nums))//测试代码

		if value >= target { //判断输入key和value值关系
			return key
		}
		if value < target && key == len(nums)-1 { //超出范围会返回最后的位置
			i = len(nums)
		}
	}
	return i
}
