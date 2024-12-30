package main

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"go-network/tcp/delimiter"
	"net"
	"strconv"
	"time"
)

func main() {
	address := "127.0.0.1:1235"
	conn, err := net.Dial("tcp", address)

	fmt.Printf("connect to %s\n", address)
	if err != nil {
		fmt.Println(err)
		return
	}

	process(conn)
}

func process(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)
	exit := make(chan bool)
	go handleWrite(conn, exit)
	go handleRead(conn, exit)
	<-exit

	fmt.Printf("close connection %v\n", conn)
}

func handleWrite(conn net.Conn, exit chan bool) {
	// 每隔3秒发送一个数据
	// writer := bufio.NewWriter(conn)
	for {

		time.Sleep(2 * time.Second)

		// 批量发送数据
		btSize := 3
		func(batchSize int, exit chan bool) {
			for i := 1; i <= batchSize; i++ {
				// uuid - 111$
				request := getUUID() + "-" + repeat(strconv.Itoa(i), batchSize) + string(delimiter.DELIMITER)
				fmt.Printf("Send to server > %s\n", request)

				_, err := conn.Write([]byte(request))

				if err != nil {
					fmt.Println(err)
					exit <- true
					return
				}
			}
		}(btSize, exit)

	}
}

func handleRead(conn net.Conn, exit chan bool) {
	reader := bufio.NewReader(conn)
	for {
		bytes, err := reader.ReadSlice(delimiter.DELIMITER)
		if err != nil {
			fmt.Println(err)
			exit <- true
			return
		}
		// 虽然用 \n 作为分割符号，接受的字节数组是包含 $ 的
		response := string(bytes)
		fmt.Printf("Receive from server > %s \n", response)
	}
}

func getUUID() string {
	return uuid.New().String()
}

func repeat(str string, count int) string {
	s := ""
	for i := 0; i < count; i++ {
		s = s + str
	}
	return s
}
