package core

import (
	"io"
	"log"
	"os"
	"testing"
)

func TestLink(t *testing.T) {
	path := "./wechatzhuanfa" //读取文件获取密钥
	fi, _ := os.Open(path)
	//fmt.Println(fi)
	n, err := io.ReadAll(fi)
	if err != nil {
		log.Fatal("err:", err)
	}
	//fmt.Println(n)
	data := []struct {
		title  string
		word   string
		answer string
	}{
		{"test", "123", "SUCCESS"}, //测试输入消息后和获取返回的错误标识进行对比
		//{"你好123", "qwe", "SUCCESS"},
		//{"凋零1", "测试", "SUCCESS"}
	}
	for _, val := range data {
		result := Link(val.title, val.word, string(n))
		if result != val.answer {
			t.Fatalf("expect:[%v] != result[%v]", val.answer, result)
		}
	}
}
