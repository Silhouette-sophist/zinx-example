package main

import "zinx-example/znet"

func main() {
	server := znet.Server{
		Name:      "test",
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8081,
	}
	server.Serve()
}
