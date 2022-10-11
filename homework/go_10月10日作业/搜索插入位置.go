package main

import "fmt"

func main() {

	data := []int{1, 3, 5, 6}

	target := 7
	var s int
	s = sEarchInsert(data, target)
	fmt.Println(s)

}

func sEarchInsert(nums []int, target int) int { //和内部函数重复，改名后提交通过
	var i int = 1
	for key, value := range nums {
		//fmt.Println(key, value, len(nums))//测试代码

		if value >= target { //判断输入key和value值关系
			return key
		} else if value < target && key == len(nums)-1 { //超出范围会返回最后的位置
			i = len(nums)
		}
	}
	return i
}
