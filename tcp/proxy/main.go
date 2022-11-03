package main

func main() {

	lc := listenConfig{
		port: 3306,
	}

	cc := clientConfig{
		host: "172.100.1.100",
		port: 3306,
	}

	createProxy(lc, cc)
}
