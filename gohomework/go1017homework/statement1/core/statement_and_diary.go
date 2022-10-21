package core

import (
	"fmt"
	"io"
	"net/http"
)

func Statement() (string, error) {
	const x = "https://api.ixiaowai.cn/ylapi/index.php"
	resp, err1 := http.Get(x)
	if err1 != nil {
		fmt.Println("http.err:", err1)
	}
	if resp.StatusCode != 200 {
		return "非正常输出", fmt.Errorf(resp.Status)
	}
	b, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return string(b), nil
}

func Diary() (string, error) {
	const y = "https://api.ixiaowai.cn/tgrj/index.php"
	resp, err1 := http.Get(y)
	if err1 != nil {
		fmt.Println("http.err:", err1)
	}
	if resp.StatusCode != 200 {
		return "非正常输出", fmt.Errorf(resp.Status)
	}
	b, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return string(b), nil
}
