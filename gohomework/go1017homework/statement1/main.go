package main

import (
	"fmt"
	"goProject1/gohomework/go1017homework/statement1/core"
)

func main() {
	fmt.Println("输入你想要获得的文字，1：一言语录，2：舔狗日记，3：退出")

	for {

		var i int
		fmt.Scanf("%d", &i)
		var (
			str string
			err error
		)
		switch i {

		case 1: //调用一眼语录
			str, err = core.Statement()

		case 2: //调用舔狗日记
			str, err = core.Diary()
		case 3: //退出系统
			fmt.Println("退出系统")
			return
		default:
			fmt.Println("重新输入指令")

		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(str)

	}

}
