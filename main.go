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
	conn.Write([]byte("Hello, world!\n"))
	buf := make([]byte, 4096)
	total := 0
	for {
		n, err := conn.Read(buf)
		if err != nil {
			logger.Println(err)
			return
		}
		total += n
		if n == 4 {
			continue
		}
		conn.Write([]byte(fmt.Sprintf("You said %d bytes to me.\n", total)))
		total = 0
	}
}
