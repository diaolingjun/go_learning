package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const html = "https://sctapi.ftqq.com/"

// 主要函数，调用接口发送信息
func Link(title, desp, key string) (errr string) {
	link := html + key + ".send?title=" + title + "&desp=" + desp
	//fmt.Println("访问的连接", link)
	result, err := http.Get(link) //调用api，发送消息
	if err != nil {
		fmt.Println("http.get err:", err)
		return
	}
	html := Getlink(result) //调用函数
	//punshid, readkey := jiexi(html)
	//fmt.Println(punshid, readkey)
	punshid, readkey, err1 := jiexi(html)
	checkLink(punshid, readkey)
	return err1

}

func checkLink(punsid string, readkey string) {
	url := "https://sctapi.ftqq.com/push?id=" + punsid + "&readkey=" + readkey
	fmt.Println(url)
	result, err := http.Get(url) //调用checkapi获取状态
	if err != nil {
		fmt.Println("http.get2 err", err)
		return
	}
	//html2 := Getlink(result)
	html2 := Getlink(result)
	//fmt.Println("返回数据", html2) // {"code":0,"message":"","data":{"id":91670454,"uid":"177051","title":"666","desp":"diaoling","encoded":"","readkey":"SCTUTXnhKFLbGMU","wxstatus":"","created_at":"2022-10-18T16:36:14+08:00","updated_at":"2022-10-18T16:36:14+08:00"}}
	str := jiexi2(html2)
	fmt.Println(str)
}

// 获取网页返回json状态码，根据生成链接爬取网页内容
func Getlink(result *http.Response) string {

	buf, err := io.ReadAll(result.Body)
	if err != nil {
		fmt.Println("io.ReadAll err:", err)
	}
	fmt.Println(string(buf))
	defer result.Body.Close() //关闭连接
	return string(buf)
}

// 解析json数据获得想要的部分
func jiexi(jshtml string) (string, string, string) {
	var jss Returnjson
	err := json.Unmarshal([]byte(jshtml), &jss)
	//// 空接口什么都可以所以就要处理切片类型，第二个参数要地址传递
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
	}
	pushid1 := jss.Data.Pushid
	readkey1 := jss.Data.Readkey
	err1 := jss.Data.Error
	//fmt.Println("测试输出", pushid1, readkey1)
	return pushid1, readkey1, err1
}
func jiexi2(jshtml string) string {

	var js2 fanhuijson
	err := json.Unmarshal([]byte(jshtml), &js2)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
	}
	str := js2.Data.Wxstatus
	return str
}

type Returnjson struct {
	Data struct {
		Readkey string `json:"readkey"`
		Pushid  string `json:"pushid"`
		Error   string `json:"error"`
	} `json:"data"` //首字母必须是大写的，小写的读不出来，如果需要给为小写，可以在后面添加反引号，
}
type fanhuijson struct {
	Data struct {
		Wxstatus string `json:"wxstatus"`
	} `json:"data"` //首字母必须是大写的，小写的读不出来，如果需要给为小写，可以在后面添加反引号，
}

// 测试函数
