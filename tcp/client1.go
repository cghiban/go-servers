package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	log.Println("trying to connect to localhost:8080")
	log.Println("(server1.go would be a good start for the server")
	// get a listener
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(bs))
}
