package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

var logger = log.New(os.Stdout, "", log.Llongfile|log.LstdFlags)

const HttpRequest = `GET /worker HTTP/1.1
Host: localhost:8000

`

func main() {
	CallWorker()
}

func CallWorker() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		logger.Fatalln(err)
	}
	defer conn.Close()

	conn.Write([]byte(HttpRequest))
	buf := make([]byte, 4096) // 假定回复内容不超过4K字节
	_, err = conn.Read(buf)
	if err != nil {
		logger.Println(err)
		return
	}
	fmt.Println(string(buf))

	conn.Write([]byte(HttpRequest))
	// buf := make([]byte, 4096) // 一个极致危险的优化：不清空buf
	_, err = conn.Read(buf)
	if err != nil {
		logger.Println(err)
		return
	}
	fmt.Println(string(buf))
}

func Manager() {
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
