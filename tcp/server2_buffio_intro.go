package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	//io.WriteString(conn, "\nHi From server!\n")
	//fmt.Fprintln(conn, "How are things?")

	scanner := bufio.NewScanner(conn)
	//scanner.Split(bufio.ScanWords)
	//scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		line := scanner.Text()
		log.Println("got text: ", line)
	}

	fmt.Println("Connection about to close!")
}
