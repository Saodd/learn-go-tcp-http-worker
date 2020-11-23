package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

var logger = log.New(os.Stdout, "", log.Llongfile|log.LstdFlags)

func main() {
	server, err := net.Listen("tcp", "0.0.0.0:1080")
	if err != nil {
		logger.Fatalln(err)
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			logger.Println(err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	remote := conn.RemoteAddr().String()
	defer func() {
		conn.Close()
		logger.Println(remote, "closed!")
	}()
	logger.Println(remote)
	buf := make([]byte, 4096) // 假定请求内容不超过4K字节
	_, err := conn.Read(buf)
	if err != nil {
		logger.Println(err)
		return
	}
	fmt.Println(string(buf))
}
