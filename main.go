package main

import "server"

func main() {
	port := server.GetPort()
	server := server.New(server.Config{
		Host: "172.20.10.6",
		Port: port,
	})
	server.Run()
}
