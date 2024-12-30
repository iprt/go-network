package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
)

// listenerConfig
// @Description: 监听配置
type listenerConfig struct {
	port int
}

// backendConfig
// @Description: 请求配置
type backendConfig struct {
	host string
	port int
}

// creatProxy
//
//	@Description: 创建tcp代理
//	@param lConfig
//	@param bConfig
func createProxy(lConfig listenerConfig, bConfig backendConfig) {
	address := ":" + strconv.Itoa(lConfig.port)

	listen, err := net.Listen("tcp4", address)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("start proxy at %s \n", address)

	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			fmt.Println("listener close occurred error")
		}
	}(listen)

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Receive a connection ", conn.RemoteAddr())
		go handleConnection(conn, bConfig)
	}

}

// handleConnection
//
//	@Description: 处理客户端的连接
//	@param conn
//	@param config
func handleConnection(conn net.Conn, bConfig backendConfig) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)

	proxyAddress := bConfig.host + ":" + strconv.Itoa(bConfig.port)

	proxyConn, err := net.Dial("tcp", proxyAddress)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func(proxyConn net.Conn) {
		err := proxyConn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(proxyConn)

	exit := make(chan bool, 1)

	go transferByIoCopy(conn, proxyConn, exit)
	go transferByIoCopy(proxyConn, conn, exit)
	<-exit

	fmt.Println("Proxy occurred error, and then exit")
}

// transfer
//
//	@Description: 传输数据
//	@param from
//	@param to
func transfer(from, to net.Conn, exit chan bool) {
	buffer := make([]byte, 4096)
	for {
		bLen, fromErr := from.Read(buffer)
		if fromErr != nil {
			fmt.Println(fromErr)
			exit <- true
			return
		}

		_, toErr := to.Write(buffer[:bLen])
		if toErr != nil {
			fmt.Println(toErr)
			exit <- true
			return
		}
	}
}

// transferByIoCopy
//
//	@Description: reference https://www.cnblogs.com/mignet/p/go_transfer.html
//	@param from
//	@param to
//	@param exit
func transferByIoCopy(from, to net.Conn, exit chan bool) {
	_, err := io.Copy(to, from)
	if err != nil {
		fmt.Println(err)
		exit <- true
		return
	}
}
