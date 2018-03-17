package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 100)

	for {
		n, readErr := conn.Read(buf)
		if readErr != nil {
			fmt.Println("readErr: ", readErr)
			return
		}

		fmt.Println("buf = ", string(buf[:n]))
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func main() {
	// 监听本机 8000 端口
	listener, err := net.Listen("tcp", "127.0.0.1:8000")

	if err != nil {
		fmt.Println("Listen err:", err)
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("accept error: ", err)
			return
		}

		go HandleConn(conn)
	}
}
