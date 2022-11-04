package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
)

//
//  listenConfig
//  @Description: 监听配置
//
type listenConfig struct {
	port int
}

//
//  clientConfig
//  @Description: 请求配置
//
type clientConfig struct {
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
func createProxy(lc listenConfig, cc clientConfig) {
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

		cConn, err := net.Dial("tcp", ccStr)

		if err != nil {
			fmt.Println("occurred error when establish a connection to the proxy")
			continue
		}

		// go listenerTransferData(lConn, cConn)
		go listenerTransferDataByCopy(lConn, cConn)
		go clientTransferDataByCopy(lConn, cConn)
	}

}

//
// listenerTransferData
//  @Description: implement (user -> listener -> client)
//  @param lConn
//  @param cConn
//
func listenerTransferData(lConn, cConn net.Conn) {
	for {
		lReadBuffer := make([]byte, 1024)
		lLen, lReadErr := lConn.Read(lReadBuffer)
		if lReadErr != nil {
			fmt.Println("transfer date from listener's client occurred error")
			return
		} else {
			if lLen > 0 {
				_, lWriteErr := cConn.Write(lReadBuffer[:lLen])
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
//  @param cConn
//
func listenerTransferDataByCopy(lConn, cConn net.Conn) {
	for {
		_, err := io.Copy(cConn, lConn)
		if err != nil {
			fmt.Println("transfer date from listener's client occurred error !")
		}
	}
}

//
// clientTransferData
//  @Description: implement (user <- listener <- client)
//  @param lConn
//  @param cConn
//
func clientTransferData(lConn, cConn net.Conn) {
	for {
		cReadBuffer := make([]byte, 1024)
		cLen, cReadErr := cConn.Read(cReadBuffer)
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
// clientTransferDataByCopy
//  @Description: implement (user <- listener <- client)
//  @param lConn
//  @param cConn
//
func clientTransferDataByCopy(lConn, cConn net.Conn) {
	for {
		_, err := io.Copy(lConn, cConn)
		if err != nil {
			fmt.Println("transfer date from proxy's server occurred error !")
		}
	}
}
