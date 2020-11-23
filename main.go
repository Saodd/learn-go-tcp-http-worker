package main

import (
	"bytes"
	"encoding/json"
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
	data := CallWorker()
	fmt.Println(data)
}

func CallWorker() int {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		logger.Fatalln(err)
	}
	defer conn.Close()

	conn.Write([]byte(HttpRequest))
	buf := make([]byte, 4096) // 假定回复内容不超过4K字节
	n, err := conn.Read(buf)
	if err != nil {
		logger.Println(err)
		return 0
	}

	// 识别出json的部分
	index := bytes.Index(buf[:n], []byte("\r\n\r\n"))
	// 这里暂时不处理找不到的情况
	js := buf[index+4 : n]
	var body WorkerBody
	err = json.Unmarshal(js, &body)
	if err != nil {
		logger.Println(err)
	}
	return body.Data
}

type WorkerBody struct {
	Data int `json:"data"`
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
