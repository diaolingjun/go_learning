package go_10_13_homework

import "testing"

// 算法函数
func MiddleNode(head *ListNode) *ListNode {
	var point *ListNode
	point = head
	num := 0
	for point != nil { //循环找到最后的链表位置
		num = num + 1 //统计链表长度
		point = point.Next
	}
	num = num / 2
	point = head
	for num > 0 {
		point = point.Next //遍历到中间位置的链表位置
		num = num - 1
	}
	return point

}

// 结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

// 生成数组函数
func newlist(args ...int) *ListNode {
	var head ListNode
	var p *ListNode = &head //定义指针用来初始化赋值
	for _, v := range args {
		node2 := new(ListNode)
		p.Next = node2
		node2.Val = v
		p = p.Next

	}
	return head.Next //不包含头结点，所以得排除在外
}

// 测试函数
func TestMiddleNode(t *testing.T) {

	data := []struct {
		list   *ListNode
		answer int
	}{
		{newlist(1, 2, 3, 4, 5), 3},
		{newlist(1, 2, 3, 4, 5, 6), 4},
		{newlist(1), 1},
	}
	for _, val := range data {
		result := MiddleNode(val.list)
		if val.answer != result.Val {
			t.Fatalf("expect:[%v] != result[%v]", val.answer, result.Val)
		}
	}
}
