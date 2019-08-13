package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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
		go handle3(conn)
	}
}

func handle3(conn net.Conn) {
	defer conn.Close()
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("Connection timedout!")
	}
	scanner := bufio.NewScanner(conn)

	// check for data in the stream
	for scanner.Scan() {
		// get the data, line by line
		line := scanner.Text()
		fmt.Println("got text from client: ", line)
		fmt.Fprintln(conn, "I head you say: ", line)
	}

	fmt.Println("Connection closed!")
}
