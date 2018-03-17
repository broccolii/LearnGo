package main

import (
	"fmt"
	"os"
	"net"
	"io"
)

// 发送文件
func SendFile(path string, conn net.Conn) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Print("os.Open err = ", err)
		return
	}

	defer file.Close()

	buf := make([]byte, 1024 * 4)

	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("f.Read err = ", err)
			}
			return
		}

		conn.Write(buf[:n])
	}
}

func main() {
	fmt.Println("请输入文件路径:")

	var path string
	fmt.Scan(&path)

	// 获取文件
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat err = ", err)
		return
	}

	// 连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial err = ", err);
		return
	}

	defer conn.Close()

	// 发送文件名
	n, err := conn.Write([]byte(info.Name()))
	if err != nil {
		fmt.Println("net.Write err = ", err)
		return
	}

	// 等待服务器返回 ok 信息
	buf := make([]byte, 1024)
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("net.Read err = ", err)
	}

	if "ok" == string(buf[:n]) {
		SendFile(path, conn)
	}
}