package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
)

//
//  listenerConfig
//  @Description: 监听配置
//
type listenerConfig struct {
	port int
}

//
//  backendConfig
//  @Description: 请求配置
//
type backendConfig struct {
	host string
	port int
}

//
// createProxy
//  @Description: user -> listener -> client
//                user <- listener <- client
//  @param lc
//  @param cc
//
func createProxy(lc listenerConfig, cc backendConfig) {
	fmt.Println("start listening server")

	PORT := ":" + strconv.Itoa(lc.port)
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {
			fmt.Println("listener close occurred error")
		}
	}(l)

	fmt.Println("start listener at ", PORT)

	for {
		lConn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Every time a connection is established, establish a connection to the proxy
		ccStr := cc.host + ":" + strconv.Itoa(cc.port)
		fmt.Println("connect to client server", ccStr)

		backConn, err := net.Dial("tcp", ccStr)

		if err != nil {
			fmt.Println("occurred error when establish a connection to the proxy")
			return
		}

		// go listenerTransferData(lConn, cConn)
		// go clientTransferData(lConn, cConn)
		go listenerTransferDataByCopy(lConn, backConn)
		go backendTransferDataByCopy(lConn, backConn)

	}

}

//
// listenerTransferData
//  @Description: implement (user -> listener -> client)
//  @param lConn
//  @param backConn
//
func listenerTransferData(lConn, backConn net.Conn) {
	lReadBuffer := make([]byte, 4096)
	for {
		lLen, lReadErr := lConn.Read(lReadBuffer)
		if lReadErr != nil {
			fmt.Println("transfer date from listener's client occurred error")
			return
		} else {
			if lLen > 0 {
				_, lWriteErr := backConn.Write(lReadBuffer[:lLen])
				fmt.Printf("transfer data to client (bytes's length is %d)\n", lLen)
				if lWriteErr != nil {
					fmt.Println("transfer data to client occurred error !!!")
					return
				}
			}
		}
	}
}

//
// listenerTransferDataByCopy
//  @Description: implement (user -> listener -> client)
//  @param lConn
//  @param backConn
//
func listenerTransferDataByCopy(lConn, backConn net.Conn) {
	_, err := io.Copy(backConn, lConn)
	if err != nil {
		fmt.Println("transfer date from listener's client occurred error !")
		return
	}
}

//
// backendTransferData
//  @Description: implement (user <- listener <- client)
//  @param lConn
//  @param backConn
//
func backendTransferData(lConn, backConn net.Conn) {
	cReadBuffer := make([]byte, 4096)
	for {
		cLen, cReadErr := backConn.Read(cReadBuffer)
		if cReadErr != nil {
			fmt.Println("transfer date from proxy's server occurred error !")
			return
		} else {
			if cLen > 0 {
				n, cWriteErr := lConn.Write(cReadBuffer[:cLen])
				fmt.Printf("transfer data to listener (bytes's length is %d)\n", n)
				if cWriteErr != nil {
					fmt.Println("transfer data to listener occurred error !!!")
					return
				}
			}
		}
	}
}

//
// backendTransferDataByCopy
//  @Description: implement (user <- listener <- client)
//  @param lConn
//  @param backConn
//
func backendTransferDataByCopy(lConn, backConn net.Conn) {
	_, err := io.Copy(lConn, backConn)
	if err != nil {
		fmt.Println("transfer date from proxy's server occurred error !")
		return
	}
}
