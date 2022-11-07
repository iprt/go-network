package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

var count = 0

func main() {
	arguments := os.Args

	var PORT string
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		// return
		PORT = ":1234"
	} else {
		PORT = ":" + arguments[1]
	}

	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	fmt.Println("start tcp server at", PORT)

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
		count++
	}
}

func handleConnection(c net.Conn) {
	fmt.Println(".")
	defer c.Close()
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}
		fmt.Println(temp)
		counter := strconv.Itoa(count) + "\n"
		c.Write([]byte(counter))
	}
}
