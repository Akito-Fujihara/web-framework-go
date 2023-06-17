package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	for {
		// time.Sleep(time.Second * 3)

		buf := make([]byte, 100)

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		RequestInfo := string(buf[:n])

		fmt.Println("request:")
		fmt.Println(RequestInfo)

		if RequestInfo == `"close"` {
			fmt.Println("closing from client....")
			conn.Close()
			return
		}

		responseData := "response"

		responseByteData, err := json.Marshal(responseData)

		_, err = conn.Write(responseByteData)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
