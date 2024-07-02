package main

import (
	"fmt"
	"os"
	"server"
)

func main() {
	port := server.GetPort()

	localIp, errIP := server.GetLocalIP()
	if errIP != nil {
		fmt.Fprintln(os.Stderr, "An unexpected error has occured retrieving local IP:", errIP.Error())
		os.Exit(1)
	}

	server := server.New(server.Config{
		Host: localIp,
		Port: port,
	})
	server.Run()
}
