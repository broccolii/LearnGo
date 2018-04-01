package main

import (
	"net"
	"fmt"
	"os"
	"io"
)

func RecvFile(fileName string, conn net.Conn) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err = ", err)
		return
	}
	
	buf := make([]byte, 1024 * 4)
	
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("conn.Read err = ", err)
			}
			return
		}
		f.Write(buf[:n])
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}
	defer listener.Close()

	// 阻塞等待用户连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err = ", err)
		return
	}
	defer conn.Close()

	// 读取对方发送的文件名
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err = ", err)
		return
	}

	fileName := string(buf[:n])

	// 回复 "ok"
	conn.Write([]byte("ok"))

	// 接收文件
	RecvFile(fileName, conn)
}
