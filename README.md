# go network

## Create a TCP Client and Server

[serverMain.go](./tcp/simple/server/serverMain.go)

[clientMain.go](./tcp/simple/client/clientMain.go)

[https://www.linode.com/docs/guides/developing-udp-and-tcp-clients-and-servers-in-go/](https://www.linode.com/docs/guides/developing-udp-and-tcp-clients-and-servers-in-go/)

start sever

```shell
go run serverMain.go 8888
```

start client

```shell
go run clientMain.go 127.0.0.1:8888
```

**goland Edit Configuration**

- `Program arguments`

## Create a UDP Client and Server

[serverMain.go](./udp/simple/server/serverMain.go)

[clientMain.go](./udp/simple/client/clientMain.go)

[https://www.linode.com/docs/guides/developing-udp-and-tcp-clients-and-servers-in-go/](https://www.linode.com/docs/guides/developing-udp-and-tcp-clients-and-servers-in-go/)

## Create a Concurrent TCP Server

[concServer.go](./tcp/conc/server/concServer.go)

## Creat a Reconnectable TCP Client

[recoClient.go](./tcp/reconnect/client/recoClient.go)

## Create a TCP Proxy

[proxy.go](./tcp/proxy/proxy.go)