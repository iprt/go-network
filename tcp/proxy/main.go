package main

func main() {

	lc := listenerConfig{
		port: 3306,
	}

	cc := backendConfig{
		host: "172.100.1.100",
		port: 3306,
	}

	createProxy(lc, cc)
}
