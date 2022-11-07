package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	arguments := os.Args
	var CONNECT string
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		// return
		CONNECT = "127.0.0.1:1234"
	} else {
		CONNECT = arguments[1]
	}

	for {
		conn, err := net.Dial("tcp", CONNECT)
		fmt.Println("start connect", CONNECT)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("connect", CONNECT, "success")
			defer conn.Close()
			handleConnection(conn)
		}

		time.Sleep(3 * time.Second)
	}

}

func handleConnection(conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		_, writeErr := fmt.Fprintf(conn, "%s\n", text)
		if writeErr != nil {
			break
		}
		// conn.Write([]byte(text + "\n"))
		message, readErr := bufio.NewReader(conn).ReadString('\n')
		if readErr != nil {
			break
		}
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
