package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
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
		go handle2(conn)
	}
}

func handle2(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	// check for data in the stream
	for scanner.Scan() {
		// get the data, line by line
		line := scanner.Text()
		fmt.Println("got text from client: ", line)
		fmt.Fprintln(conn, "I head you say: ", line)
	}

	//fmt.Println("Connection about to close!")
}
