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

	redisMap := make(map[string]string)

	// check for data in the stream
	for scanner.Scan() {
		// get the data, line by line
		line := scanner.Text()
		fs := strings.Fields(line) // splits to words
		fmt.Println("got text from client: ", line)

		switch fs[0] {
		case "SET":
			if len(fs) < 3 {
				fmt.Fprintln(conn, "Error: Expected value is missing!")
				continue
			}
			redisMap[fs[1]] = fs[2]
		case "GET":
			fmt.Fprintln(conn, redisMap[fs[1]])
		case "DEL":
			delete(redisMap, fs[1])
		case "DUMP":
			fmt.Fprintf(conn, "%v\n", redisMap)
		default:
			fmt.Fprintln(conn, "INVALID COMMAND!")
		}

		//fmt.Fprintf(conn, "%s\n\n", )
	}

	//fmt.Println("Connection about to close!")
}
