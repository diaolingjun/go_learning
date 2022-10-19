package go1017homework

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

// Statement 获取一言语录
func Statement() (string, error) {
	const api = "https://api.ixiaowai.cn/ylapi/index.php"
	resp, err := http.Get(api)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf(resp.Status)
	}
	b, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	//fmt.Println(string(b))
	return string(b), err
}

// DogDiary 获取舔狗日记
func TestStatement(t *testing.T) {
	for i := 0; i <= 3; i++ {
		str, _ := Statement()
		if str == " " {
			t.Fatalf("函数出错")
		}
	}

}
