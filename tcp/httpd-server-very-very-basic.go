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
	i := 0
	content := ""
	for scanner.Scan() {
		// get the data, line by line
		line := scanner.Text()
		if line == "" {
			break
		}
		fmt.Println("got text from client: ", line)
		if i == 0 {
			fs := strings.Fields(line) // splits to words
			content += fmt.Sprintf("<div>method: %s</div>\n", fs[0])
			content += fmt.Sprintf("<div>uri: %s</div>\n", fs[1])
		} else {
			//fmt.Fprintln(conn, line)
		}
		i++
		//fmt.Fprintf(conn, "%s\n\n", )
	}
	respose(conn, content)

}

func respose(conn net.Conn, content string) {
	tmpl := `<!DOCTYPE html><html lang="en">
	<head><title>hi</title></head>
	<body>%s</body>
	</html>
	`
	body := fmt.Sprintf(tmpl, content)
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(conn, "Content-type: text/html\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n\r\n", len(body))
	fmt.Fprintln(conn, body)
}
