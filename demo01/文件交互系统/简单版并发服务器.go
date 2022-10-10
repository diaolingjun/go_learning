package main

import (
	"fmt"
	"net"
)

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8011")
	if err != nil {
		fmt.Println("err", err)
		return
	}

	//执行完毕关闭端口连接
	defer listener.Close()
	//阻塞等待用户连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.accept err:", err)
		return
	}
	buf := make([]byte, 1024) //留下1024大小的缓冲区
n:
	n, err1 := conn.Read(buf) //读取缓冲区
	if err1 != nil {

	}
}
