package main

import (
	"bufio"
	"fmt"
	"go-network/tcp/delimiter"
	"net"
	"strings"
)

func main() {
	address := ":1235"
	listen, err := net.Listen("tcp4", address)

	fmt.Printf("start server at %s \n", address)

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go process(conn)
	}
}

// process
//
//	@Description: 边界符解决粘包问题
//	@param conn
func process(conn net.Conn) {
	reader := bufio.NewReader(conn)
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)
	for {
		// 读取的是字节
		buffer, err := reader.ReadSlice(delimiter.DELIMITER)
		if err != nil {
			fmt.Println(err)
			return
		}

		// 接受的字符是包含 分隔符 的
		response := string(buffer)
		fmt.Printf("Receive from client > %s\n", response)

		// 转换成大写返回给客户端
		reply := strings.ToUpper(response)

		_, wErr := conn.Write([]byte(reply))

		if wErr != nil {
			fmt.Println(wErr)
			return
		}
		fmt.Printf("Reply to client > %s\n", reply)
		fmt.Println()
	}
}
