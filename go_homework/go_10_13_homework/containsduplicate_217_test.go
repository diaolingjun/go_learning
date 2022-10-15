package go_10_13_homework

import "testing"

func ContainsDuplicate(nums []int) bool {
	newmap := make(map[int]int) //利用map的检索来查询重复
	for _, v := range nums {
		_, ok := newmap[v]
		if ok {
			return true //查找成功即存在重复
		} else {
			newmap[v] = 1 //如果没有重复就将v的值存储到map中
		}

	}
	return false
}

func TestContainsDuplicate(t *testing.T) {
	data := []struct {
		nums   []int
		answer bool
	}{
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 1, 1, 1}, true},
	}
	for _, val := range data {
		result := ContainsDuplicate(val.nums)
		if result != val.answer {
			t.Fatalf("expect:[%v] != result[%v]", result, val.answer)
		}
	}
}
