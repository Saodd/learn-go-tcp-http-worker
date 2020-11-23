package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/url"
	"strconv"
	"sync"
)

const HttpRequestStr = `GET /worker HTTP/1.1
Host: any

`

type ConnWorker struct {
	conn net.Conn
	buf  []byte
}

var HttpRequestBytes = []byte(HttpRequestStr)
var ConnPool = sync.Pool{New: func() interface{} {
	return &ConnWorker{buf: make([]byte, 1024)}
}}

func OneCall(address string) int {
	w := ConnPool.Get().(*ConnWorker)
	defer ConnPool.Put(w)
	if w.conn == nil {
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Println(err)
			return 0
		}
		w.conn = conn
	}

	conn := w.conn
	buf := w.buf
	conn.Write(HttpRequestBytes)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	// 识别出json的部分
	index := bytes.Index(buf[:n], []byte("\r\n\r\n"))
	// 这里暂时不处理找不到的情况
	js := buf[index+4 : n]
	var body WorkerBody
	err = json.Unmarshal(js, &body)
	if err != nil {
		fmt.Println(err)
	}
	return body.Data
}

type WorkerBody struct {
	Data int `json:"data"`
}

func work(cb chan<- int, address string) {
	cb <- OneCall(address)
}

func manage(n int, u string) int {
	uu, err := url.Parse(u)
	if err != nil {
		fmt.Println(err)
	}
	cb := make(chan int, 10)
	for i := 0; i < n; i++ {
		go work(cb, uu.Host)
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += <-cb
	}
	return sum
}

func main() {
	g := gin.New()
	g.GET("/manager", func(c *gin.Context) {
		u := c.Query("url")
		numStr := c.Query("n")
		num, _ := strconv.Atoi(numStr)
		data := manage(num, u)
		c.JSON(200, gin.H{"data": data})
	})
	g.Run("0.0.0.0:5004")
}
