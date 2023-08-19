package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

const (
	PORT = "8080"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", PORT))
	if err != nil {
		panic(nil)
	}

	fmt.Printf("listening on :%s", PORT)
	for {
		conn, err := lis.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			fmt.Println("From: ", conn.RemoteAddr())
			req, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}
			fmt.Printf("Requst: %v", req)

			res := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body: io.NopCloser(
					strings.NewReader("Hello, Client!")),
			}
			res.Write(conn)
			conn.Close()
		}()
	}
}
