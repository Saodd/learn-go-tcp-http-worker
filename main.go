package main

import (
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
	logger.Println(remote)
	conn.Write([]byte("Hello, world!\n"))
	conn.Close()
}
