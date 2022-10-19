package go1010homework

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	data := []struct { //定义一个具有输入值和答案的结构体
		val    string
		answer string
	}{
		{val: "abcd", answer: "dcba"},
		{val: "5123", answer: "3215"},
		{val: "你好世界", answer: "界世好你"}, //初始化
	}

	for _, v := range data {
		tmp := []byte(v.val)  //定义一个数组将数据输入进去
		reverseString(tmp[:]) //执行函数获得返回值
		if string(tmp) != v.answer {
			t.Fatalf("expect:[%s] != result[%s]", string(tmp), v.answer)
		}
	}
}
func reverseString(s []byte) {
	str := []rune(string(s)) //先将byte转为字符串，再将其转为字符数组
	i, j := 0, len(str)-1
	for i < j {
		str[i], str[j] = str[j], str[i] //交换两位置处的数据
		i++
		j--
	}

	strbyte := []byte(string(str)) //定义一个byte类型的数组
	for i = 0; i < len(s); i++ {
		s[i] = strbyte[i] //将之前的数据赋值过去
	}
}
