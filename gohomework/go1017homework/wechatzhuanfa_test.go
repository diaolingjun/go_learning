package go1017homework

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

// const SendKey = "SCT177051TE8Or5OV0emXDX0bqKhZLRav3"
const html = "https://sctapi.ftqq.com/SCT177051TE8Or5OV0emXDX0bqKhZLRav3.send"

// 主要函数，调用接口发送信息
func Link(title, desp string) (errr string) {
	link := html + "?title=" + title + "&desp=" + desp
	//fmt.Println("访问的连接", link)
	result, err := http.Get(link) //调用api，发送消息
	if err != nil {
		fmt.Println("http.get err:", err)
		return
	}
	html := Getlink(result) //调用函数
	//punshid, readkey := jiexi(html)
	//fmt.Println(punshid, readkey)
	_, _, err1 := jiexi(html)
	//checkLink(punshid, readkey)
	return err1

}

//	func checkLink(punsid string, readkey string) {
//		url := "https://sctapi.ftqq.com/push?id=" + punsid + "&readkey=" + readkey
//		fmt.Println(url)
//		result, err := http.Get(url) //调用checkapi获取状态
//		if err != nil {
//			fmt.Println("http.get2 err", err)
//			return
//		}
//		html2 := Getlink(result)
//		fmt.Println("返回数据", html2) // {"code":0,"message":"","data":{"id":91670454,"uid":"177051","title":"666","desp":"diaoling","encoded":"","readkey":"SCTUTXnhKFLbGMU","wxstatus":"","created_at":"2022-10-18T16:36:14+08:00","updated_at":"2022-10-18T16:36:14+08:00"}}
//
// }

// 获取网页返回json状态码，根据生成链接爬取网页内容
func Getlink(result *http.Response) (htm string) {
	var html string //存储返回的网页值
	buf := make([]byte, 1024*8)
	for {
		n, _ := result.Body.Read(buf)
		if n == 0 { //读取结束，或者出问题
			//fmt.Println("result.body.read err:", err)
			break
		}
		html += string(buf[:n]) //获取网页返回json状态码{"code":0,"message":"","data":{"pushid":"91663840","readkey":"SCTbohz6w9CEkIO","error":"SUCCESS","errno":0}}
		//fmt.Println("html:", html)

	}
	defer result.Body.Close() //关闭连接
	return html
}

// 解析json数据获得想要的部分
func jiexi(jshtml string) (punshid string, readkey string, err2 string) {
	var jss Returnjson
	err := json.Unmarshal([]byte(jshtml), &jss)
	//// 空接口什么都可以所以就要处理切片类型，第二个参数要地址传递
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}
	//if jss.Data["pushid"].(string)
	punshid1, _ := jss.Data["pushid"].(string)
	readkey1, _ := jss.Data["readkey"].(string)
	err1, _ := jss.Data["error"].(string)
	return punshid1, readkey1, err1
}

type Returnjson struct {
	Data map[string]interface{} `json:"data"` //首字母必须是大写的，小写的读不出来，如果需要给为小写，可以在后面添加反引号，
}

// 测试函数
func TestLink(t *testing.T) {
	data := []struct {
		title  string
		word   string
		answer string
	}{
		{"test", "123", "success"},//测试输入消息后和获取返回的错误标识进行对比
		{"你好123", "qwe", "success"},
		{"凋零", "测试", "success"}}
	for _, val := range data {
		result := Link(val.title, val.word)
		if result == val.answer {
			t.Fatalf("expect:[%v] != result[%v]", val.answer, result)
		}
	}
}
