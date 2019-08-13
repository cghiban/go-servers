package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	// get a listener
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()

	for {
		// acceppt connections
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// handle the connection
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	// check for data in the stream
	for scanner.Scan() {
		// get the data, line by line
		line := strings.ToLower(scanner.Text())
		fmt.Println("got text from client: ", line)
		bsr13 := rot13([]byte(line))
		fmt.Fprintf(conn, "%s\n\n", bsr13)
	}

	//fmt.Println("Connection about to close!")
}

func rot13(bs []byte) []byte {
	r13 := make([]byte, len(bs))

	for i, v := range bs {
		// ascii 97-122
		if v <= 109 {
			r13[i] = bs[i] + 13
		} else {
			r13[i] = bs[i] - 13
		}
	}

	return r13
}
