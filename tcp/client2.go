package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	//log.Println("trying to connect to localhost:8080")
	//log.Println("(server3_rw.go would be a good start for the server")
	// get a listener
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		fmt.Fprintf(conn, "Hello, Server. I'm a client #%d\n", i)
		/*bs, err := ioutil.ReadAll(conn)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Server says: ", string(bs))*/
	}

}
