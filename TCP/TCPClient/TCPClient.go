package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 创建连接
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Dial err = ", err)
	}

	defer conn.Close()

	go func() {
		str := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("os.Stdin.Read error = ", err)
				return
			}

			conn.Write(str[:n])
		}
	}()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read error = ", err)
			return
		}
		fmt.Println("conn.write buf = ", string(buf[:n]))
	}
}
