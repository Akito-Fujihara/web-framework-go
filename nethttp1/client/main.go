package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	RequestAndResponse(conn, "request 1")
	RequestAndResponse(conn, "request 2")
	RequestAndResponse(conn, "request 3")
	RequestAndResponse(conn, "close")
	time.Sleep(time.Hour)
}

func RequestAndResponse(conn net.Conn, requestData string) {
	requestByteData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Write(requestByteData)

	responseData := make([]byte, 100)
	n, err := conn.Read(responseData)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("response:")
	fmt.Println(string(responseData[:n]))
}
